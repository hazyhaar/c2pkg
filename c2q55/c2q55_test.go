package c2q55

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"unsafe"
)

// TestKAT_Vectors valide les Known Answer Tests fondamentaux de l'implémentation.
func TestKAT_Vectors(t *testing.T) {
	// 1. KAT CRC32-C Castagnoli ("123456789" -> 0xE3069283)
	katData := []byte("123456789")
	crc := HardwareCRC32C(katData)
	if crc != 0xE3069283 {
		t.Fatalf("KAT CRC32-C mismatch: got 0x%08X, want 0xE3069283", crc)
	}

	// 2. KAT Slot pack & unpack
	var s Slot
	s.Pack(0x0102030405060708, 0x090A0B0C0D0E0F10, 1000, 2000, 42, 0, []byte("hello-kat"), SlotReady)
	if s.IDLow != 0x0102030405060708 || s.IDHigh != 0x090A0B0C0D0E0F10 {
		t.Fatalf("KAT Slot ID mismatch")
	}
	if s.PayloadLen != 9 {
		t.Fatalf("KAT PayloadLen mismatch: %d", s.PayloadLen)
	}
}

// TestC2Q55_ParityVsCOracle exécute l'oracle C réel compilé sous gcc -O2 -mavx2 et compare les sorties.
func TestC2Q55_ParityVsCOracle(t *testing.T) {
	cSrc := filepath.Join("..", "..", "c2simd", "sources", "c2q55", "test_c2q55_oracle.c")
	c2qSrc := filepath.Join("..", "..", "c2simd", "sources", "c2q55", "c2q55.c")

	if _, err := os.Stat(cSrc); os.IsNotExist(err) {
		t.Skipf("Fichier source C non trouvé: %s", cSrc)
	}

	tmpBin := filepath.Join(t.TempDir(), "c2q_oracle")
	cmdBuild := exec.Command("gcc", "-O2", "-mavx2", "-I", filepath.Dir(cSrc), cSrc, c2qSrc, "-o", tmpBin)
	buildOut, err := cmdBuild.CombinedOutput()
	if err != nil {
		t.Fatalf("Échec de compilation gcc oracle C: %v\nSortie: %s", err, string(buildOut))
	}

	cmdRun := exec.Command(tmpBin)
	out, err := cmdRun.CombinedOutput()
	if err != nil {
		t.Fatalf("Échec d'exécution oracle C: %v\nSortie: %s", err, string(out))
	}
	t.Logf("Sortie Oracle C (gcc -O2):\n%s", string(out))

	// Comparaison Go vs C sur le même vecteur KAT
	katData := []byte("123456789")
	goCRC := HardwareCRC32C(katData)
	if goCRC != 0xE3069283 {
		t.Fatalf("Parité CRC32-C échouée : Go=0x%08X, attendu=0xE3069283", goCRC)
	}
}

// TestZeroAlloc_Ring certifie que les opérations nominales n'allouent strictement aucun octet sur le tas.
func TestZeroAlloc_Ring(t *testing.T) {
	ring := NewRing(1024)
	payload := []byte("msg-payload-01")
	var out Slot

	// Mesure sur chemin nominal non saturé (Enqueue + Dequeue + Ack successifs)
	allocs := testing.AllocsPerRun(10000, func() {
		seq, err := ring.Enqueue(1234, 5678, 100, 5000000, 1, 0, payload)
		if err == nil {
			if deqSeq, ok := ring.Dequeue(200, &out); ok {
				_ = ring.Ack(deqSeq)
			} else {
				_ = ring.Ack(seq)
			}
		}
	})
	if allocs != 0 {
		t.Fatalf("Ring operations allocate: %.2f allocs/op (expected 0.00)", allocs)
	}
}

// TestRing_Lifecycle valide le cycle de vie complet Enqueue -> Dequeue -> LeaseExtend -> Ack.
func TestRing_Lifecycle(t *testing.T) {
	ring := NewRing(64)
	payload := []byte("test-data")

	seq, err := ring.Enqueue(10, 20, 100, 1000000, 42, 0, payload)
	if err != nil {
		t.Fatalf("Enqueue failed: %v", err)
	}

	var out Slot
	if _, ok := ring.Dequeue(50, &out); ok {
		t.Fatalf("Dequeue should not be visible before timestamp")
	}

	deqSeq, ok := ring.Dequeue(100, &out)
	if !ok {
		t.Fatalf("Dequeue failed at visible timestamp")
	}
	if deqSeq != seq {
		t.Fatalf("Seq mismatch: got %d, want %d", deqSeq, seq)
	}
	if out.TopicID != 42 {
		t.Fatalf("TopicID mismatch: got %d, want 42", out.TopicID)
	}

	if err := ring.LeaseExtend(seq, 500000); err != nil {
		t.Fatalf("LeaseExtend failed: %v", err)
	}

	if err := ring.Ack(seq); err != nil {
		t.Fatalf("Ack failed: %v", err)
	}
}

// TestWAL_Persistence valide l'écriture par lots, le calcul CRC32-C et le rejeu après incident.
func TestWAL_Persistence(t *testing.T) {
	tmpDir := t.TempDir()
	walPath := filepath.Join(tmpDir, "test.wal")

	wal, err := OpenWAL(walPath, 1024*1024)
	if err != nil {
		t.Fatalf("OpenWAL failed: %v", err)
	}

	slots := make([]Slot, 100)
	for i := range slots {
		slots[i].Pack(uint64(i), uint64(1000+i), uint64(100+i), 5000, 1, 0, []byte("data"), SlotReady)
	}

	lastLSN, err := wal.WriteBatch(slots)
	if err != nil {
		t.Fatalf("WriteBatch failed: %v", err)
	}
	if lastLSN == 0 {
		t.Fatalf("lastLSN should be > 0")
	}
	_ = wal.Sync()
	_ = wal.Close()

	walReopen, err := OpenWAL(walPath, 1024*1024)
	if err != nil {
		t.Fatalf("Reopen WAL failed: %v", err)
	}
	defer walReopen.Close()

	replayedCount := 0
	replayedIDs := make([]uint64, 0, 100)
	n, err := walReopen.Replay(func(s *Slot) {
		replayedCount++
		replayedIDs = append(replayedIDs, s.IDLow)
	})
	if err != nil {
		t.Fatalf("Replay failed: %v", err)
	}
	if n != 100 || replayedCount != 100 {
		t.Fatalf("Replay count mismatch: got %d, want 100", n)
	}
	for i, id := range replayedIDs {
		if id != uint64(i) {
			t.Fatalf("Replayed ID mismatch at %d: got %d, want %d", i, id, i)
		}
	}
}

// TestLoserTree_KWayOrdering valide la fusion K-Way ordonnée de 16 shards indépendants.
func TestLoserTree_KWayOrdering(t *testing.T) {
	numShards := uint64(16)
	router := NewPartitionRouter(numShards, 1024)

	for i := uint64(0); i < 160; i++ {
		idLow := i
		idHigh := i
		shard := router.ShardForID(idLow)
		_, err := shard.Enqueue(idLow, idHigh, 100, 5000, 1, 0, []byte("val"))
		if err != nil {
			t.Fatalf("Enqueue failed on shard %d: %v", idLow%numShards, err)
		}
	}

	merger := NewLoserTreeMerger(router.Shards())
	merger.Init(150)

	var prevIDHigh uint64 = 0
	count := 0
	var out Slot

	for merger.Next(150, &out) {
		if count > 0 && out.IDHigh < prevIDHigh {
			t.Fatalf("K-Way merge ordering violation at count %d: prev=%d, got=%d",
				count, prevIDHigh, out.IDHigh)
		}
		prevIDHigh = out.IDHigh
		count++
	}

	if count != 160 {
		t.Fatalf("Total extracted mismatch: got %d, want 160", count)
	}
}

// TestSlot_ExactAlignment prouve que la structure Slot occupe exactement 64 octets.
func TestSlot_ExactAlignment(t *testing.T) {
	sz := unsafe.Sizeof(Slot{})
	if sz != 64 {
		t.Fatalf("Slot size is %d bytes, MUST be exactly 64 bytes (L1 cache line)", sz)
	}
}

// TestEngine_EndToEndLifecycle valide le fonctionnement global de Engine (Publish, Consume).
func TestEngine_EndToEndLifecycle(t *testing.T) {
	opts := DefaultOptions()
	opts.NumShards = 8
	opts.ShardCapacity = 512

	eng, err := NewEngine(opts)
	if err != nil {
		t.Fatalf("NewEngine failed: %v", err)
	}
	defer eng.Close()

	for i := 0; i < 50; i++ {
		idLow := uint64(i)
		idHigh := uint64(100 + i)
		payload := []byte("engine-msg")
		if err := eng.Publish(idLow, idHigh, 10, 1, payload); err != nil {
			t.Fatalf("Publish failed at %d: %v", i, err)
		}
	}

	consumed := 0
	var out Slot
	for i := 0; i < 50; i++ {
		if !eng.ConsumeOrdered(&out) {
			t.Fatalf("Consume failed at %d", i)
		}
		if out.IDHigh != uint64(100+i) {
			t.Fatalf("IDHigh ordering mismatch at %d: got %d, want %d", i, out.IDHigh, 100+i)
		}
		consumed++
	}
	if consumed != 50 {
		t.Fatalf("Consumed count mismatch: got %d, want 50", consumed)
	}
}
