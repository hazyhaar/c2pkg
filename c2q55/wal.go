// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2q55

import (
	"fmt"
	"hash/crc32"
	"os"
	"sort"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"
)

const (
	WALBlockSize     = 4096
	WALMagic         = uint32(0x43325157) // "C2QW"
	MaxSlotsPerBlock = 63                 // 63 * 64B = 4032B + 64B Header = 4096B
)

// Table Castagnoli matérielle (utilise SSE4.2 / ARMv8 PMULL à > 36 Go/s)
var castagnoliTable = crc32.MakeTable(crc32.Castagnoli)

// HardwareCRC32C calcule la somme de contrôle Castagnoli accélérée matériellement.
func HardwareCRC32C(data []byte) uint32 {
	return crc32.Checksum(data, castagnoliTable)
}

// WALHeader représente l'en-tête structuré de 64 octets en tête de chaque bloc 4K (ARCHTIME POD Layout).
type WALHeader struct {
	Magic       uint32   // 0..3   - 0x43325157
	EntryCount  uint32   // 4..7   - Nombre de slots valides dans le bloc (1..63)
	LSN         uint64   // 8..15  - Numéro de séquence de log monotone
	TimestampNs uint64   // 16..23 - Horodatage physique du bloc
	CRC32C      uint32   // 24..27 - Somme de contrôle Castagnoli sur les données
	_           [36]byte // 28..63 - Padding aligné 64 octets exacts
}

// Invariants matériels ARCHTIME bidirectionnels : taille strictement égale à 64 octets.
var (
	_ [64 - unsafe.Sizeof(WALHeader{})]byte
	_ [unsafe.Sizeof(WALHeader{}) - 64]byte
)

// WALFile gère la persistance circulaire mmapée sur disque sans aucun syscall sur le chemin chaud.
type WALFile struct {
	file       *os.File
	filePath   string
	sizeBytes  int64
	mmapData   []byte
	currentLSN atomic.Uint64
	writePos   int64
	mu         sync.Mutex
	blockBuf   [WALBlockSize]byte
}

type blockEntry struct {
	lsn        uint64
	offset     int64
	entryCount uint32
}

// OpenWAL ouvre, préalloue et mappe en mémoire virtuelle (mmap) un fichier WAL circulaire.
func OpenWAL(filePath string, sizeBytes int64) (*WALFile, error) {
	if sizeBytes < WALBlockSize*16 {
		sizeBytes = WALBlockSize * 16
	}

	flags := os.O_CREATE | os.O_RDWR | syscall.O_NOATIME
	f, err := os.OpenFile(filePath, flags, 0600)
	if err != nil {
		f, err = os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0600)
		if err != nil {
			return nil, fmt.Errorf("c2q55/wal: open failed: %w", err)
		}
	}

	fi, err := f.Stat()
	if err != nil {
		_ = f.Close()
		return nil, err
	}

	if fi.Size() < sizeBytes {
		if err := f.Truncate(sizeBytes); err != nil {
			_ = f.Close()
			return nil, fmt.Errorf("c2q55/wal: preallocate failed: %w", err)
		}
	}

	mmapData, err := syscall.Mmap(int(f.Fd()), 0, int(sizeBytes), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		_ = f.Close()
		return nil, fmt.Errorf("c2q55/wal: mmap failed: %w", err)
	}

	w := &WALFile{
		file:      f,
		filePath:  filePath,
		sizeBytes: sizeBytes,
		mmapData:  mmapData,
	}

	var maxLSN uint64 = 0
	var headPos int64 = 0

	for offset := int64(0); offset+WALBlockSize <= sizeBytes; offset += WALBlockSize {
		hdr := (*WALHeader)(unsafe.Pointer(&mmapData[offset]))
		if hdr.Magic != WALMagic {
			continue
		}

		if hdr.EntryCount == 0 || hdr.EntryCount > MaxSlotsPerBlock {
			continue
		}

		dataLen := int64(hdr.EntryCount * 64)
		calcCRC := HardwareCRC32C(mmapData[offset+64 : offset+64+dataLen])
		if hdr.CRC32C != calcCRC {
			continue
		}

		if hdr.LSN > maxLSN {
			maxLSN = hdr.LSN
			headPos = offset + WALBlockSize
			if headPos >= sizeBytes {
				headPos = 0
			}
		}
	}

	w.currentLSN.Store(maxLSN)
	w.writePos = headPos

	return w, nil
}

// WriteBatch sérialise directement en mémoire virtuelle mmapée (Zéro Syscall Write).
func (w *WALFile) WriteBatch(slots []Slot) (uint64, error) {
	if len(slots) == 0 {
		return w.currentLSN.Load(), nil
	}

	w.mu.Lock()
	defer w.mu.Unlock()

	var lastLSN uint64
	total := len(slots)
	offset := 0

	for offset < total {
		chunkLen := total - offset
		if chunkLen > MaxSlotsPerBlock {
			chunkLen = MaxSlotsPerBlock
		}

		lsn := w.currentLSN.Add(1)
		lastLSN = lsn

		for i := range w.blockBuf {
			w.blockBuf[i] = 0
		}

		for i := 0; i < chunkLen; i++ {
			slotPtr := (*[64]byte)(unsafe.Pointer(&slots[offset+i]))
			copy(w.blockBuf[64+i*64:64+(i+1)*64], slotPtr[:])
		}

		dataLen := chunkLen * 64
		crc := HardwareCRC32C(w.blockBuf[64 : 64+dataLen])

		hdr := (*WALHeader)(unsafe.Pointer(&w.blockBuf[0]))
		hdr.Magic = WALMagic
		hdr.EntryCount = uint32(chunkLen)
		hdr.LSN = lsn
		hdr.TimestampNs = uint64(time.Now().UnixNano())
		hdr.CRC32C = crc

		if w.writePos+WALBlockSize > w.sizeBytes {
			w.writePos = 0
		}

		copy(w.mmapData[w.writePos:w.writePos+WALBlockSize], w.blockBuf[:])

		w.writePos += WALBlockSize
		offset += chunkLen
	}

	return lastLSN, nil
}

// Sync synchronise les données sur le support physique via fdatasync (zéro écriture d'inode).
func (w *WALFile) Sync() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	return syscall.Fdatasync(int(w.file.Fd()))
}

// Close libère la projection mémoire mmap et ferme le descripteur de fichier.
func (w *WALFile) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.mmapData != nil {
		_ = syscall.Munmap(w.mmapData)
		w.mmapData = nil
	}
	return w.file.Close()
}

// Replay relit les blocs valides directement depuis la mémoire mmapée selon l'ordre strict LSN.
func (w *WALFile) Replay(onSlot func(slot *Slot)) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	var validBlocks []blockEntry

	for offset := int64(0); offset+WALBlockSize <= w.sizeBytes; offset += WALBlockSize {
		hdr := (*WALHeader)(unsafe.Pointer(&w.mmapData[offset]))
		if hdr.Magic != WALMagic {
			continue
		}

		if hdr.EntryCount == 0 || hdr.EntryCount > MaxSlotsPerBlock {
			continue
		}

		dataLen := int64(hdr.EntryCount * 64)
		calcCRC := HardwareCRC32C(w.mmapData[offset+64 : offset+64+dataLen])
		if hdr.CRC32C != calcCRC {
			continue
		}

		validBlocks = append(validBlocks, blockEntry{
			lsn:        hdr.LSN,
			offset:     offset,
			entryCount: hdr.EntryCount,
		})
	}

	if len(validBlocks) == 0 {
		return 0, nil
	}

	sort.Slice(validBlocks, func(i, j int) bool {
		return validBlocks[i].lsn < validBlocks[j].lsn
	})

	startIdx := 0
	maxLSN := validBlocks[len(validBlocks)-1].lsn
	maxCapacityBlocks := uint64(w.sizeBytes / WALBlockSize)

	for i, b := range validBlocks {
		if maxLSN >= maxCapacityBlocks && b.lsn < (maxLSN-maxCapacityBlocks+1) {
			startIdx = i + 1
		}
	}

	replayedSlots := 0
	for i := startIdx; i < len(validBlocks); i++ {
		b := validBlocks[i]
		blockSlice := w.mmapData[b.offset:]

		for j := uint32(0); j < b.entryCount; j++ {
			s := (*Slot)(unsafe.Pointer(&blockSlice[64+j*64]))
			onSlot(s)
			replayedSlots++
		}
	}

	return replayedSlots, nil
}
