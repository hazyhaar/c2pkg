package c2q55

import (
	"path/filepath"
	"testing"
)

// BenchmarkRing_Enqueue_Dequeue mesure le débit brut en mémoire vive L1 d'un anneau unique.
func BenchmarkRing_Enqueue_Dequeue(b *testing.B) {
	ring := NewRing(65536)
	payload := []byte("bench-payload-01")
	var out Slot

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		seq, _ := ring.Enqueue(uint64(i), uint64(i), 10, 5000000, 1, 0, payload)
		_, _ = ring.Dequeue(20, &out)
		_ = ring.Ack(seq)
	}
}

// BenchmarkPartitionRouter_Throughput mesure le débit agrégé sur 16 shards indépendants.
func BenchmarkPartitionRouter_Throughput(b *testing.B) {
	router := NewPartitionRouter(16, 65536)
	payload := []byte("bench-payload-01")
	var out Slot

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		idLow := uint64(i)
		shard := router.ShardForID(idLow)
		seq, _ := shard.Enqueue(idLow, uint64(i), 10, 5000000, 1, 0, payload)
		_, _ = shard.Dequeue(20, &out)
		_ = shard.Ack(seq)
	}
}

// BenchmarkLoserTree_Merge mesure le débit de fusion K-Way en mémoire L1.
func BenchmarkLoserTree_Merge(b *testing.B) {
	capacity := uint64(b.N + 1024)
	if capacity < 65536 {
		capacity = 65536
	}
	router := NewPartitionRouter(16, capacity)
	payload := []byte("bench-val")

	for i := 0; i < b.N; i++ {
		idLow := uint64(i)
		shard := router.ShardForID(idLow)
		_, _ = shard.Enqueue(idLow, uint64(i), 10, 5000000, 1, 0, payload)
	}

	merger := NewLoserTreeMerger(router.Shards())
	merger.Init(50)

	var out Slot
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		merger.Next(50, &out)
	}
}

// BenchmarkWAL_WriteBatch mesure le débit d'écriture sur le journal WAL.
func BenchmarkWAL_WriteBatch(b *testing.B) {
	tmpDir := b.TempDir()
	walPath := filepath.Join(tmpDir, "bench.wal")
	wal, err := OpenWAL(walPath, 64*1024*1024)
	if err != nil {
		b.Fatalf("OpenWAL failed: %v", err)
	}
	defer wal.Close()

	batch := make([]Slot, 60)
	for i := range batch {
		batch[i].Pack(uint64(i), uint64(i), 10, 5000000, 1, 0, []byte("bench-val"), SlotReady)
	}

	b.SetBytes(int64(len(batch) * 64))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = wal.WriteBatch(batch)
	}
}
