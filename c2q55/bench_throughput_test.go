// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2q55

import (
	"path/filepath"
	"testing"
)

// =========================================================================
// BENCHMARK D'INGESTION ET CONSOMMATION RÉELLES (AVEC CHARGES STRUCTURÉES)
// =========================================================================

// BenchmarkComp_C2Q55_PublishMonoCore mesure le cycle complet Publish + Consume sur 1 cœur avec charges réelles.
func BenchmarkComp_C2Q55_PublishMonoCore(b *testing.B) {
	opts := DefaultOptions()
	opts.NumShards = 16
	opts.ShardCapacity = 65536

	eng, err := NewEngine(opts)
	if err != nil {
		b.Fatalf("NewEngine failed: %v", err)
	}
	defer eng.Close()

	payload := MakeRealisticOrderPayload(42)
	var out Slot

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		id := uint64(i)
		_ = eng.Publish(id, id, 10, 1, payload)
		_ = eng.Consume(&out)
	}
}

// BenchmarkComp_C2Q55_PublishMultiCore_16Workers mesure le cycle complet Publish + Consume en parallèle sur 32 threads.
func BenchmarkComp_C2Q55_PublishMultiCore_16Workers(b *testing.B) {
	opts := DefaultOptions()
	opts.NumShards = 16
	opts.ShardCapacity = 262144

	eng, err := NewEngine(opts)
	if err != nil {
		b.Fatalf("NewEngine failed: %v", err)
	}
	defer eng.Close()

	payload := MakeRealisticOrderPayload(101)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		var id uint64 = 0
		var out Slot
		for pb.Next() {
			id++
			_ = eng.Publish(id, id, 10, 1, payload)
			_ = eng.Consume(&out)
		}
	})
}

// BenchmarkComp_C2Q55_WAL_GroupCommit_4KB mesure le débit d'écriture par micro-lots 4K (63 slots avec hash réel).
func BenchmarkComp_C2Q55_WAL_GroupCommit_4KB(b *testing.B) {
	tmpDir := b.TempDir()
	walPath := filepath.Join(tmpDir, "comp_c2q.wal")
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

// BenchmarkComp_C2Q55_WAL_HardwareFDatasync mesure le coût d'une barrière matérielle fdatasync.
func BenchmarkComp_C2Q55_WAL_HardwareFDatasync(b *testing.B) {
	tmpDir := b.TempDir()
	walPath := filepath.Join(tmpDir, "comp_c2q_sync.wal")
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
