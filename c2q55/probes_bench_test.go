// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2q55

import (
	"path/filepath"
	"sync/atomic"
	"testing"
)

var SinkCRC uint32

// =========================================================================
// STRATE 1 : MICRO-NOYAUX DE CALCUL & CRC (0-Memory Raw Latency)
// =========================================================================

func BenchmarkStrat1_Scalar_ScanBlock_8Slots(b *testing.B) {
	var block C2q_meta_block_t
	for i := 0; i < 8; i++ {
		block.Visible_at[i] = uint64(100 + i*10)
		block.State[i] = SlotReady
	}
	now := uint64(145)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = C2q_scan_block(&block, now)
	}
}

func BenchmarkStrat1_CRC32C_Hardware_4KB(b *testing.B) {
	// Données de télémétrie structurées réelles de 4 032 octets
	block := MakeRealisticDocumentPayload(4032)

	b.SetBytes(4032)
	b.ReportAllocs()
	b.ResetTimer()

	var acc uint32
	for i := 0; i < b.N; i++ {
		block[0] = byte(i)
		acc += HardwareCRC32C(block)
	}
	SinkCRC = acc
}

// =========================================================================
// STRATE 2 : ANNEAU MONO-SHARD LOCK-FREE (L1 Cache-Line Contention)
// =========================================================================

func BenchmarkStrat2_Ring_SingleCore_EnqueueDequeue(b *testing.B) {
	ring := NewRing(65536)
	payload := []byte("usr_sess_9401ab8")
	var out Slot

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		seq, _ := ring.Enqueue(uint64(i), uint64(i), 10, 5000000, 1, 0, payload)
		_, _ = ring.Dequeue(20, &out)
		_ = ring.Ack(seq)
	}
}

func BenchmarkStrat2_Ring_Parallel_EnqueueWithDrain(b *testing.B) {
	ring := NewRing(262144)
	payload := []byte("usr_sess_9401ab8")

	var stopDrain atomic.Bool
	for c := 0; c < 4; c++ {
		go func() {
			var out Slot
			for !stopDrain.Load() {
				if seq, ok := ring.Dequeue(1000000000000, &out); ok {
					_ = ring.Ack(seq)
				}
			}
		}()
	}
	defer stopDrain.Store(true)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		var i uint64 = 0
		for pb.Next() {
			i++
			_, _ = ring.Enqueue(i, i, 10, 5000000, 1, 0, payload)
		}
	})
}

// =========================================================================
// STRATE 3 : ROUTEUR MULTI-SHARDS & CONCURRENCE (Scalabilité Multicœur Drainée)
// =========================================================================

func BenchmarkStrat3_Router_MultiCore_PublishWithDrain(b *testing.B) {
	opts := DefaultOptions()
	opts.NumShards = 16
	opts.ShardCapacity = 262144

	eng, err := NewEngine(opts)
	if err != nil {
		b.Fatalf("NewEngine failed: %v", err)
	}
	defer eng.Close()

	payload := MakeRealisticTelemetryPayload(42)

	var stopDrain atomic.Bool
	for c := 0; c < 8; c++ {
		go func() {
			var out Slot
			for !stopDrain.Load() {
				eng.Consume(&out)
			}
		}()
	}
	defer stopDrain.Store(true)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		var id uint64 = 0
		for pb.Next() {
			id++
			_ = eng.Publish(id, id, 10, 1, payload)
		}
	})
}

func BenchmarkStrat3_Router_MultiCore_ConsumeLockFree(b *testing.B) {
	opts := DefaultOptions()
	opts.NumShards = 16
	opts.ShardCapacity = 262144

	eng, err := NewEngine(opts)
	if err != nil {
		b.Fatalf("NewEngine failed: %v", err)
	}
	defer eng.Close()

	payload := MakeRealisticTelemetryPayload(101)
	for i := uint64(0); i < 100000; i++ {
		_ = eng.Publish(i, i, 10, 1, payload)
	}

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		var out Slot
		for pb.Next() {
			eng.Consume(&out)
		}
	})
}

// =========================================================================
// STRATE 4 : FUSION MONOTONE K-WAY LOSER TREE (O(log K) L1D Cache Residency)
// =========================================================================

func BenchmarkStrat4_LoserTree_K16_OrderedMerge(b *testing.B) {
	capacity := uint64(b.N + 1024)
	if capacity < 65536 {
		capacity = 65536
	}
	router := NewPartitionRouter(16, capacity)
	payload := []byte("ord_tx_01a02fe")

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

func BenchmarkStrat4_LoserTree_K64_OrderedMerge(b *testing.B) {
	capacity := uint64(b.N + 1024)
	if capacity < 65536 {
		capacity = 65536
	}
	router := NewPartitionRouter(64, capacity)
	payload := []byte("ord_tx_01a02fe")

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

// =========================================================================
// STRATE 5 : PERSISTENCE WAL DISQUE (Group Commit & NVMe Flush)
// =========================================================================

func BenchmarkStrat5_WAL_GroupCommit_PageCache(b *testing.B) {
	tmpDir := b.TempDir()
	walPath := filepath.Join(tmpDir, "bench_pagecache.wal")
	wal, err := OpenWAL(walPath, 64*1024*1024)
	if err != nil {
		b.Fatalf("OpenWAL failed: %v", err)
	}
	defer wal.Close()

	batch := make([]Slot, 63)
	for i := range batch {
		payload := MakeRealisticOrderPayload(i)
		batch[i].Pack(uint64(i), uint64(i), 10, 5000000, 1, 0, payload, SlotReady)
	}

	b.SetBytes(4096)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = wal.WriteBatch(batch)
	}
}

func BenchmarkStrat5_WAL_SingleFSync_Latency(b *testing.B) {
	tmpDir := b.TempDir()
	walPath := filepath.Join(tmpDir, "bench_fsync.wal")
	wal, err := OpenWAL(walPath, 64*1024*1024)
	if err != nil {
		b.Fatalf("OpenWAL failed: %v", err)
	}
	defer wal.Close()

	batch := make([]Slot, 63)
	for i := range batch {
		payload := MakeRealisticOrderPayload(i)
		batch[i].Pack(uint64(i), uint64(i), 10, 5000000, 1, 0, payload, SlotReady)
	}

	b.SetBytes(4096)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = wal.WriteBatch(batch)
		_ = wal.Sync()
	}
}
