// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2q55

import (
	"fmt"
	"math"
	"path/filepath"
	"sync"
	"sync/atomic"
	"testing"
)

// TestWAL_WrapAroundMonotonicReplay_RED teste l'enroulement cyclique sur un micro-WAL (64 Ko)
// et vérifie que le rejeu restitue exactement le sous-ensemble monotone valide du dernier tour.
func TestWAL_WrapAroundMonotonicReplay_RED(t *testing.T) {
	tmpDir := t.TempDir()
	walPath := filepath.Join(tmpDir, "micro_wrap.wal")

	// Micro-WAL de 64 Ko (exactement 16 blocs de 4 Ko, 16 * 60 = 960 slots max par tour pour batchSize=60)
	walSize := int64(64 * 1024)
	wal, err := OpenWAL(walPath, walSize)
	if err != nil {
		t.Fatalf("OpenWAL failed: %v", err)
	}

	const totalMsgs = 3000
	batchSize := 60
	for i := 0; i < totalMsgs; i += batchSize {
		batch := make([]Slot, batchSize)
		for j := 0; j < batchSize; j++ {
			msgID := uint64(i + j)
			batch[j].Pack(msgID, msgID, 100, 5000, 1, 0, []byte(fmt.Sprintf("m-%04d", msgID)), SlotReady)
		}
		_, err := wal.WriteBatch(batch)
		if err != nil {
			t.Fatalf("WriteBatch failed at %d: %v", i, err)
		}
	}
	_ = wal.Sync()
	_ = wal.Close()

	walReopen, err := OpenWAL(walPath, walSize)
	if err != nil {
		t.Fatalf("Reopen WAL failed: %v", err)
	}
	defer walReopen.Close()

	var prevID uint64 = 0
	replayedCount := 0

	n, err := walReopen.Replay(func(s *Slot) {
		if replayedCount > 0 {
			if s.IDLow <= prevID {
				t.Fatalf("VIOLATION DE MONOTONICITÉ WAL APRÈS ENROULEMENT à l'index %d: prevID=%d, curID=%d",
					replayedCount, prevID, s.IDLow)
			}
		}
		prevID = s.IDLow
		replayedCount++
	})
	if err != nil {
		t.Fatalf("Replay failed: %v", err)
	}

	if n != 960 {
		t.Fatalf("Assertion rejeu enroulement échouée : attendu 960 slots du dernier tour, obtenu %d", n)
	}
}

// TestRing_Uint64BoundaryWrap_RED teste le franchissement de la frontière MaxUint64 (2^64 - 1 -> 0).
func TestRing_Uint64BoundaryWrap_RED(t *testing.T) {
	ring := NewRing(64)
	payload := []byte("boundary-msg")

	ring.producer.head.Store(math.MaxUint64 - 5)
	ring.consumer.tail.Store(math.MaxUint64 - 5)

	for i := 0; i < 20; i++ {
		_, err := ring.Enqueue(uint64(i), uint64(i), 10, 5000, 1, 0, payload)
		if err != nil {
			t.Fatalf("Enqueue failed during boundary wrap at %d: %v", i, err)
		}
	}

	var out Slot
	for i := 0; i < 20; i++ {
		seq, ok := ring.Dequeue(20, &out)
		if !ok {
			t.Fatalf("CONSUMER FROZEN: Dequeue failed during boundary wrap at %d", i)
		}
		_ = ring.Ack(seq)
		if out.IDLow != uint64(i) {
			t.Fatalf("ID mismatch during boundary wrap: got %d, want %d", out.IDLow, i)
		}
	}
}

// TestEngine_ParallelConsume_NoMutexLock_RED teste la non-régression et l'exactitude de consommation multi-consommateurs.
func TestEngine_ParallelConsume_NoMutexLock_RED(t *testing.T) {
	opts := DefaultOptions()
	opts.NumShards = 16
	opts.ShardCapacity = 8192

	eng, err := NewEngine(opts)
	if err != nil {
		t.Fatalf("NewEngine failed: %v", err)
	}
	defer eng.Close()

	const total = 10000
	for i := 0; i < total; i++ {
		id := uint64(i)
		if err := eng.Publish(id, id, 10, 1, []byte("par-msg")); err != nil {
			t.Fatalf("Publish failed at %d: %v", i, err)
		}
	}

	var wg sync.WaitGroup
	var consumedCount sync.Map
	var totalReceived atomic.Uint64

	for c := 0; c < 16; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var out Slot
			for {
				if eng.Consume(&out) {
					if _, loaded := consumedCount.LoadOrStore(out.IDLow, true); loaded {
						t.Errorf("DOUBLON DÉTECTÉ: ID %d consommé plusieurs fois", out.IDLow)
					}
					totalReceived.Add(1)
				} else {
					break
				}
			}
		}()
	}
	wg.Wait()

	if totalReceived.Load() != total {
		t.Fatalf("Assertion de consommation parallèle échouée : attendu %d messages uniques, reçu %d", total, totalReceived.Load())
	}
}
