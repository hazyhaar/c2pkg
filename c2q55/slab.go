package c2q55

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"syscall"
)

var (
	ErrSlabCorrupted = errors.New("c2q55/slab: body CRC32-C checksum mismatch")
	ErrSlabFull      = errors.New("c2q55/slab: slab arena capacity exceeded (live data protected by low watermark)")
)

const (
	FlagExternalPayload uint32 = 0x1000 // Indique que le corps est stocké dans l'arène de slabs
	DefaultSlabSize     int64  = 64 * 1024 * 1024 // 64 Mo préalloués par arène
	MaxPayloadSize      uint32 = 16 * 1024 * 1024 // 16 Mo max par message
	SlabHeaderMagic     uint32 = 0x4332534C // "C2SL"
	SlabHeaderSize      int64  = 64         // 64 octets d'en-tête persistant
)

// SlabDescriptor représente le pointeur de 16 octets inliné dans le Slot pour les charges lourdes.
type SlabDescriptor struct {
	Offset  uint64 // 0..7  - Offset dans le fichier mmapé
	Length  uint32 // 8..11 - Longueur du corps en octets
	BodyCRC uint32 // 12..15 - Checksum matériel Castagnoli sur le corps
}

// Encode encode le descripteur dans un tampon de 16 octets.
func (d *SlabDescriptor) Encode(dst *[16]byte) {
	binary.LittleEndian.PutUint64(dst[0:8], d.Offset)
	binary.LittleEndian.PutUint32(dst[8:12], d.Length)
	binary.LittleEndian.PutUint32(dst[12:16], d.BodyCRC)
}

// DecodeSlabDescriptor décode le descripteur depuis les 16 octets du Payload d'un Slot.
func DecodeSlabDescriptor(src *[16]byte) SlabDescriptor {
	return SlabDescriptor{
		Offset:  binary.LittleEndian.Uint64(src[0:8]),
		Length:  binary.LittleEndian.Uint32(src[8:12]),
		BodyCRC: binary.LittleEndian.Uint32(src[12:16]),
	}
}

// SlabArena gère un segment de mémoire mmapé contigu protégé par ligne d'eau basse et en-tête persistant.
type SlabArena struct {
	file         *os.File
	mmapData     []byte
	sizeBytes    int64
	writePos     atomic.Uint64
	lowWatermark atomic.Uint64
	mu           sync.Mutex
}

// OpenSlabArena ouvre ou crée une arène préallouée mmapée avec reprise de position.
func OpenSlabArena(filePath string, sizeBytes int64) (*SlabArena, error) {
	if sizeBytes <= 0 {
		sizeBytes = DefaultSlabSize
	}

	flags := os.O_CREATE | os.O_RDWR | syscall.O_NOATIME
	f, err := os.OpenFile(filePath, flags, 0600)
	if err != nil {
		f, err = os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0600)
		if err != nil {
			return nil, fmt.Errorf("c2q55/slab: open failed: %w", err)
		}
	}

	fi, err := f.Stat()
	if err != nil {
		_ = f.Close()
		return nil, err
	}

	totalSize := SlabHeaderSize + sizeBytes
	if fi.Size() < totalSize {
		if err := f.Truncate(totalSize); err != nil {
			_ = f.Close()
			return nil, fmt.Errorf("c2q55/slab: preallocate failed: %w", err)
		}
	}

	mmapData, err := syscall.Mmap(int(f.Fd()), 0, int(totalSize), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		_ = f.Close()
		return nil, fmt.Errorf("c2q55/slab: mmap failed: %w", err)
	}

	a := &SlabArena{
		file:      f,
		mmapData:  mmapData,
		sizeBytes: sizeBytes,
	}

	// Lecture de l'en-tête persistant pour reprise après redémarrage
	magic := binary.LittleEndian.Uint32(mmapData[0:4])
	if magic == SlabHeaderMagic {
		savedWrite := binary.LittleEndian.Uint64(mmapData[8:16])
		savedWatermark := binary.LittleEndian.Uint64(mmapData[16:24])
		a.writePos.Store(savedWrite)
		a.lowWatermark.Store(savedWatermark)
	} else {
		binary.LittleEndian.PutUint32(mmapData[0:4], SlabHeaderMagic)
		binary.LittleEndian.PutUint32(mmapData[4:8], 1) // Version 1
		binary.LittleEndian.PutUint64(mmapData[8:16], 0)
		binary.LittleEndian.PutUint64(mmapData[16:24], 0)
		_ = syscall.Fdatasync(int(f.Fd()))
	}

	return a, nil
}

// Write stocke une charge utile lourde dans l'arène sans écraser les données sous la ligne d'eau basse.
func (a *SlabArena) Write(body []byte) (SlabDescriptor, error) {
	bodyLen := uint32(len(body))
	if bodyLen > MaxPayloadSize {
		return SlabDescriptor{}, ErrPayloadTooLarge
	}

	alignedLen := uint64((bodyLen + 63) &^ 63)
	if alignedLen > uint64(a.sizeBytes) {
		return SlabDescriptor{}, ErrPayloadTooLarge
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	curPos := a.writePos.Load()
	lowWatermark := a.lowWatermark.Load()

	rawOffset := curPos % uint64(a.sizeBytes)
	if rawOffset+alignedLen > uint64(a.sizeBytes) {
		// Pas assez d'espace contigu en fin de fichier : saut au début du cycle suivant
		nextCycle := ((curPos / uint64(a.sizeBytes)) + 1) * uint64(a.sizeBytes)
		curPos = nextCycle
		rawOffset = 0
	}

	// Protection contre le dépassement de la ligne d'eau basse (zéro écrasement de corps vivants)
	if int64(curPos+alignedLen-lowWatermark) > a.sizeBytes {
		return SlabDescriptor{}, ErrSlabFull
	}

	fileOffset := SlabHeaderSize + int64(rawOffset)

	// Copie en mémoire virtuelle projetée
	copy(a.mmapData[fileOffset:fileOffset+int64(bodyLen)], body)

	// Calcul du CRC32-C matériel Castagnoli sur le corps
	crc := HardwareCRC32C(body)

	newPos := curPos + alignedLen
	a.writePos.Store(newPos)

	// Mise à jour synchrone de l'en-tête persistant
	binary.LittleEndian.PutUint64(a.mmapData[8:16], newPos)

	return SlabDescriptor{
		Offset:  uint64(fileOffset),
		Length:  bodyLen,
		BodyCRC: crc,
	}, nil
}

// AdvanceWatermark avance la ligne d'eau basse après consommation/acquittement pour recycler l'espace.
func (a *SlabArena) AdvanceWatermark(newWatermark uint64) {
	a.mu.Lock()
	defer a.mu.Unlock()

	cur := a.lowWatermark.Load()
	if newWatermark > cur {
		a.lowWatermark.Store(newWatermark)
		binary.LittleEndian.PutUint64(a.mmapData[16:24], newWatermark)
	}
}

// Read lit une charge utile lourde avec copie défensive hors mmap et vérification CRC.
func (a *SlabArena) Read(desc SlabDescriptor) ([]byte, error) {
	if int64(desc.Offset+uint64(desc.Length)) > int64(len(a.mmapData)) {
		return nil, errors.New("c2q55/slab: offset out of bounds")
	}

	// Copie défensive obligatoire pour isoler le lecteur des écritures concurrentes
	res := make([]byte, desc.Length)
	copy(res, a.mmapData[desc.Offset:desc.Offset+uint64(desc.Length)])

	calcCRC := HardwareCRC32C(res)
	if calcCRC != desc.BodyCRC {
		return nil, ErrSlabCorrupted
	}

	return res, nil
}

// Sync force la synchronisation matérielle fdatasync du segment de slabs.
func (a *SlabArena) Sync() error {
	return syscall.Fdatasync(int(a.file.Fd()))
}

// Close libère la projection mémoire et ferme l'arène.
func (a *SlabArena) Close() error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.mmapData != nil {
		_ = syscall.Fdatasync(int(a.file.Fd()))
		_ = syscall.Munmap(a.mmapData)
		a.mmapData = nil
	}
	return a.file.Close()
}
