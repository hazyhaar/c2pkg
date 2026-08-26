// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2q55

import (
	"fmt"
	"path/filepath"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// TestTorture_ConcurrentProducersConsumers teste 128 goroutines concurrentes sous le race detector.
func TestTorture_ConcurrentProducersConsumers(t *testing.T) {
	numShards := uint64(16)
	router := NewPartitionRouter(numShards, 4096)

	numProducers := 64
	numConsumers := 64
	msgsPerProducer := 1000

	var totalProduced atomic.Uint64
	var totalConsumed atomic.Uint64

	var wg sync.WaitGroup
	startSignal := make(chan struct{})

	// Producteurs
	for p := 0; p < numProducers; p++ {
		wg.Add(1)
		go func(prodID int) {
			defer wg.Done()
			<-startSignal

			for i := 0; i < msgsPerProducer; i++ {
				idLow := uint64(prodID*100000 + i)
				idHigh := uint64(time.Now().UnixNano())
				shard := router.ShardForID(idLow)

				for {
					_, err := shard.Enqueue(idLow, idHigh, 10, 5000000, 1, 0, []byte("data"))
					if err == nil {
						totalProduced.Add(1)
						break
					}
					time.Sleep(10 * time.Microsecond)
				}
			}
		}(p)
	}

	// Consommateurs
	for c := 0; c < numConsumers; c++ {
		wg.Add(1)
		go func(consID int) {
			defer wg.Done()
			<-startSignal

			var out Slot
			for totalConsumed.Load() < uint64(numProducers*msgsPerProducer) {
				shardIdx := uint64(consID) % numShards
				shard := router.shards[shardIdx]
				if seq, ok := shard.Dequeue(100, &out); ok {
					_ = shard.Ack(seq)
					totalConsumed.Add(1)
				} else {
					time.Sleep(5 * time.Microsecond)
				}
			}
		}(c)
	}

	close(startSignal)
	wg.Wait()

	if totalProduced.Load() != uint64(numProducers*msgsPerProducer) {
		t.Fatalf("Total produced mismatch: %d", totalProduced.Load())
	}
	if totalConsumed.Load() != totalProduced.Load() {
		t.Fatalf("Total consumed mismatch: %d vs %d", totalConsumed.Load(), totalProduced.Load())
	}
}

// TestTorture_BackpressureBurst teste la saturation brutale d'un anneau sans consommateur.
func TestTorture_BackpressureBurst(t *testing.T) {
	ring := NewRing(128)
	payload := []byte("burst-msg")

	inserted := 0
	dropped := 0
	for i := 0; i < 500; i++ {
		_, err := ring.Enqueue(uint64(i), 100, 10, 5000, 1, 0, payload)
		if err == nil {
			inserted++
		} else {
			dropped++
		}
	}

	if inserted != 128 {
		t.Fatalf("Inserted count mismatch: got %d, want 128", inserted)
	}
	if dropped != (500 - 128) {
		t.Fatalf("Dropped count mismatch: got %d, want %d", dropped, 500-128)
	}
	if ring.Drops() != uint64(dropped) {
		t.Fatalf("Ring.Drops counter mismatch: got %d, want %d", ring.Drops(), dropped)
	}

	var out Slot
	for i := 0; i < 64; i++ {
		seq, ok := ring.Dequeue(50, &out)
		if !ok {
			t.Fatalf("Dequeue failed at %d", i)
		}
		_ = ring.Ack(seq)
	}

	reinserted := 0
	for i := 0; i < 64; i++ {
		_, err := ring.Enqueue(uint64(1000+i), 100, 10, 5000, 1, 0, payload)
		if err == nil {
			reinserted++
		}
	}
	if reinserted != 64 {
		t.Fatalf("Reinserted count mismatch: got %d, want 64", reinserted)
	}
}

// TestTorture_PoisonPillAndDLQ teste l'interception et l'éjection de messages corrompus/toxiques.
func TestTorture_PoisonPillAndDLQ(t *testing.T) {
	ring := NewRing(64)
	payload := []byte("poison-payload")

	_, err := ring.Enqueue(999, 100, 10, 5000, 1, 0, payload)
	if err != nil {
		t.Fatalf("Enqueue failed: %v", err)
	}

	var out Slot
	// Tentative 1 (visible à t=10, extrait à t=20, reprogrammé pour t=40)
	deqSeq, ok := ring.Dequeue(20, &out)
	if !ok {
		t.Fatalf("Dequeue 1 failed")
	}
	_ = ring.Nack(deqSeq, 3, 20, 20)

	// Tentative 2 (extrait à t=40, reprogrammé pour t=60)
	deqSeq, ok = ring.Dequeue(40, &out)
	if !ok {
		t.Fatalf("Dequeue 2 failed")
	}
	_ = ring.Nack(deqSeq, 3, 20, 40)

	// Tentative 3 (extrait à t=60, dépassement maxRetries=3 -> bascule en SlotDLQ)
	deqSeq, ok = ring.Dequeue(60, &out)
	if !ok {
		t.Fatalf("Dequeue 3 failed")
	}
	_ = ring.Nack(deqSeq, 3, 20, 60)

	// Le slot est désormais en DLQ, Dequeue ne doit plus rien trouver même à t=100
	if _, ok := ring.Dequeue(100, &out); ok {
		t.Fatalf("Slot should be in DLQ and ignored by normal Dequeue")
	}
}

// TestTorture_CrashAndWALReplay teste la persistance et le rejeu après fermeture.
func TestTorture_CrashAndWALReplay(t *testing.T) {
	tmpDir := t.TempDir()
	walPath := filepath.Join(tmpDir, "crash.wal")

	opts := DefaultOptions()
	opts.WALPath = walPath
	opts.NumShards = 4
	opts.ShardCapacity = 1024

	eng, err := NewEngine(opts)
	if err != nil {
		t.Fatalf("NewEngine failed: %v", err)
	}

	for i := 0; i < 2000; i++ {
		idLow := uint64(i)
		idHigh := uint64(1000 + i)
		p := []byte(fmt.Sprintf("msg-%04d", i))
		if err := eng.Publish(idLow, idHigh, 10, 1, p); err != nil {
			t.Fatalf("Publish failed at %d: %v", i, err)
		}
	}

	_ = eng.Close()

	w, err := OpenWAL(walPath, opts.WALSizeBytes)
	if err != nil {
		t.Fatalf("OpenWAL failed: %v", err)
	}
	defer w.Close()

	totalReplayed := 0
	n, err := w.Replay(func(s *Slot) {
		totalReplayed++
	})
	if err != nil {
		t.Fatalf("Replay failed: %v", err)
	}

	if n != 2000 || totalReplayed != 2000 {
		t.Fatalf("Crash recovery mismatch: replayed %d items (n=%d), want 2000", totalReplayed, n)
	}
}
