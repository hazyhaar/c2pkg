package c2q55

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"syscall"
)

const (
	OffsetFileMagic uint32 = 0x43324F46 // "C2OF"
	OffsetHeaderLen int    = 64
)

// OffsetStore gère la persistance synchrone sur disque des décalages (offsets) de groupes de consommateurs.
type OffsetStore struct {
	baseDir string
	files   map[string]*os.File
	mmaps   map[string][]byte
	mu      sync.RWMutex
}

// OpenOffsetStore initialise le répertoire de stockage persistant des offsets.
func OpenOffsetStore(baseDir string) (*OffsetStore, error) {
	if err := os.MkdirAll(baseDir, 0700); err != nil {
		return nil, fmt.Errorf("c2q55/offset: mkdir failed: %w", err)
	}

	return &OffsetStore{
		baseDir: baseDir,
		files:   make(map[string]*os.File),
		mmaps:   make(map[string][]byte),
	}, nil
}

func (s *OffsetStore) getOrCreateGroupMmap(groupName string, numPartitions int) ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if m, ok := s.mmaps[groupName]; ok {
		return m, nil
	}

	filePath := filepath.Join(s.baseDir, fmt.Sprintf("%s.offsets", groupName))
	fileSize := int64(OffsetHeaderLen + numPartitions*8)

	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return nil, fmt.Errorf("c2q55/offset: open failed: %w", err)
	}

	fi, err := f.Stat()
	if err != nil {
		_ = f.Close()
		return nil, err
	}

	if fi.Size() < fileSize {
		if err := f.Truncate(fileSize); err != nil {
			_ = f.Close()
			return nil, err
		}
	}

	mmapData, err := syscall.Mmap(int(f.Fd()), 0, int(fileSize), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		_ = f.Close()
		return nil, fmt.Errorf("c2q55/offset: mmap failed: %w", err)
	}

	// Initialisation de l'en-tête si nouveau
	magic := binary.LittleEndian.Uint32(mmapData[0:4])
	if magic != OffsetFileMagic {
		binary.LittleEndian.PutUint32(mmapData[0:4], OffsetFileMagic)
		binary.LittleEndian.PutUint32(mmapData[4:8], 1) // Version 1
		binary.LittleEndian.PutUint32(mmapData[8:12], uint32(numPartitions))
		copy(mmapData[12:44], []byte(groupName))
		_ = syscall.Fdatasync(int(f.Fd()))
	}

	s.files[groupName] = f
	s.mmaps[groupName] = mmapData

	return mmapData, nil
}

// CommitOffset persiste de façon synchrone le décalage d'une partition.
func (s *OffsetStore) CommitOffset(groupName string, partition int, offset uint64) error {
	m, err := s.getOrCreateGroupMmap(groupName, partition+1)
	if err != nil {
		return err
	}

	pos := OffsetHeaderLen + partition*8
	if pos+8 > len(m) {
		return errors.New("c2q55/offset: partition out of range")
	}

	binary.LittleEndian.PutUint64(m[pos:pos+8], offset)

	s.mu.RLock()
	f := s.files[groupName]
	s.mu.RUnlock()

	if f != nil {
		return syscall.Fdatasync(int(f.Fd()))
	}
	return nil
}

// LoadOffsets relit l'ensemble des offsets persistés d'un groupe au démarrage.
func (s *OffsetStore) LoadOffsets(groupName string, numPartitions int) ([]uint64, error) {
	m, err := s.getOrCreateGroupMmap(groupName, numPartitions)
	if err != nil {
		return nil, err
	}

	offsets := make([]uint64, numPartitions)
	for p := 0; p < numPartitions; p++ {
		pos := OffsetHeaderLen + p*8
		if pos+8 <= len(m) {
			offsets[p] = binary.LittleEndian.Uint64(m[pos : pos+8])
		}
	}

	return offsets, nil
}

// Close libère toutes les projections mémoire et synchronise les descripteurs.
func (s *OffsetStore) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for name, m := range s.mmaps {
		_ = syscall.Munmap(m)
		if f, ok := s.files[name]; ok {
			_ = syscall.Fdatasync(int(f.Fd()))
			_ = f.Close()
		}
	}
	s.mmaps = make(map[string][]byte)
	s.files = make(map[string]*os.File)
	return nil
}
