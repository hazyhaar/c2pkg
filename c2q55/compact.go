// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2q55

import (
	"sync"
	"time"
)

// KeyEntry représente une entrée clé-valeur avec expiration TTL.
type KeyEntry struct {
	IDLow       uint64
	IDHigh      uint64
	ExpireAtNs  uint64
	Flags       uint32
	PayloadLen  uint32
	Payload     [16]byte
	ExternalVal []byte
}

// CompactTable implémente un magasin clé-valeur compacté par clé (équivalent Kafka Compacted Topic / Redis KV).
type CompactTable struct {
	numShards int
	shardMask uint64
	shards    []*compactShard
	slab      *SlabArena
}

type compactShard struct {
	entries map[uint64]KeyEntry
	mu      sync.RWMutex
}

// NewCompactTable initialise une table compactée par clé.
func NewCompactTable(numShards int, slab *SlabArena) *CompactTable {
	if numShards <= 0 {
		numShards = 16
	}

	capPow2 := 1
	for capPow2 < numShards {
		capPow2 <<= 1
	}

	shards := make([]*compactShard, capPow2)
	for i := range shards {
		shards[i] = &compactShard{
			entries: make(map[uint64]KeyEntry),
		}
	}

	return &CompactTable{
		numShards: capPow2,
		shardMask: uint64(capPow2 - 1),
		shards:    shards,
		slab:      slab,
	}
}

// Set stocke ou met à jour une clé (avec TTL optionnel en nanosecondes).
func (t *CompactTable) Set(key uint64, val []byte, ttlNs uint64) error {
	nowNs := uint64(time.Now().UnixNano())
	var expireAt uint64 = 0
	if ttlNs > 0 {
		expireAt = nowNs + ttlNs
	}

	shardIdx := key & t.shardMask
	s := t.shards[shardIdx]

	entry := KeyEntry{
		IDLow:      key,
		IDHigh:     nowNs,
		ExpireAtNs: expireAt,
	}

	if len(val) <= 16 {
		entry.PayloadLen = uint32(len(val))
		copy(entry.Payload[:], val)
	} else if t.slab != nil {
		desc, err := t.slab.Write(val)
		if err != nil {
			return err
		}
		desc.Encode(&entry.Payload)
		entry.PayloadLen = uint32(len(val))
		entry.Flags |= FlagExternalPayload
	} else {
		entry.ExternalVal = make([]byte, len(val))
		copy(entry.ExternalVal, val)
		entry.PayloadLen = uint32(len(val))
	}

	s.mu.Lock()
	s.entries[key] = entry
	s.mu.Unlock()

	return nil
}

// Get lit la dernière valeur associée à une clé et vérifie l'expiration TTL.
func (t *CompactTable) Get(key uint64) ([]byte, bool) {
	shardIdx := key & t.shardMask
	s := t.shards[shardIdx]

	nowNs := uint64(time.Now().UnixNano())

	s.mu.RLock()
	entry, exists := s.entries[key]
	s.mu.RUnlock()

	if !exists {
		return nil, false
	}

	// Vérification de l'expiration TTL
	if entry.ExpireAtNs > 0 && entry.ExpireAtNs <= nowNs {
		s.mu.Lock()
		delete(s.entries, key)
		s.mu.Unlock()
		return nil, false
	}

	if entry.Flags&FlagExternalPayload != 0 && t.slab != nil {
		desc := DecodeSlabDescriptor(&entry.Payload)
		val, err := t.slab.Read(desc)
		if err != nil {
			return nil, false
		}
		return val, true
	}

	if len(entry.ExternalVal) > 0 {
		return entry.ExternalVal, true
	}

	res := make([]byte, entry.PayloadLen)
	copy(res, entry.Payload[:entry.PayloadLen])
	return res, true
}

// Delete supprime une clé (tombstone).
func (t *CompactTable) Delete(key uint64) bool {
	shardIdx := key & t.shardMask
	s := t.shards[shardIdx]

	s.mu.Lock()
	_, exists := s.entries[key]
	if exists {
		delete(s.entries, key)
	}
	s.mu.Unlock()

	return exists
}

// ApplySlot ingère un Slot de journal pour reconstituer la table d'état (Event Sourcing).
func (t *CompactTable) ApplySlot(slot *Slot) error {
	if slot.PayloadLen <= 16 && slot.Flags&FlagExternalPayload == 0 {
		return t.Set(slot.IDLow, slot.Payload[:slot.PayloadLen], 0)
	}
	if t.slab != nil && slot.Flags&FlagExternalPayload != 0 {
		desc := DecodeSlabDescriptor(&slot.Payload)
		val, err := t.slab.Read(desc)
		if err != nil {
			return err
		}
		return t.Set(slot.IDLow, val, 0)
	}
	return nil
}

// ReplayWAL rejoue l'intégralité d'un fichier WAL pour reconstruire l'état bit-exact.
func (t *CompactTable) ReplayWAL(walPath string, walSize int64) (int, error) {
	wal, err := OpenWAL(walPath, walSize)
	if err != nil {
		return 0, err
	}
	defer wal.Close()

	var count int
	var applyErr error
	n, err := wal.Replay(func(slot *Slot) {
		if applyErr != nil {
			return
		}
		if err := t.ApplySlot(slot); err != nil {
			applyErr = err
		} else {
			count++
		}
	})
	if err != nil {
		return count, err
	}
	if applyErr != nil {
		return count, applyErr
	}
	return n, nil
}
