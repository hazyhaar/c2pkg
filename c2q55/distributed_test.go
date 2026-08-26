// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2q55

import (
	"bytes"
	"path/filepath"
	"testing"
	"time"
)

// TestDistributed_LargePayloadSlabs_ViaEngine teste la publication et consommation de 1 Mo via Engine.
func TestDistributed_LargePayloadSlabs_ViaEngine(t *testing.T) {
	tmpDir := t.TempDir()
	slabPath := filepath.Join(tmpDir, "engine_large.slab")

	opts := DefaultOptions()
	opts.NumShards = 4
	opts.SlabPath = slabPath
	opts.SlabSizeBytes = 32 * 1024 * 1024

	eng, err := NewEngine(opts)
	if err != nil {
		t.Fatalf("NewEngine failed: %v", err)
	}
	defer eng.Close()

	// 1. Charge utile réelle de 1 Mo (1 048 576 octets)
	largeBody := MakeRealisticDocumentPayload(1024 * 1024)

	err = eng.Publish(0x12345678, 0x9ABCDEF0, 1, 0, largeBody)
	if err != nil {
		t.Fatalf("Publish 1MB body failed: %v", err)
	}

	// 2. Consommation via Message complet et vérification bit-exacte
	var msg Message
	if !eng.ConsumeMessage(&msg) {
		t.Fatalf("ConsumeMessage failed to return message")
	}

	if len(msg.Body) != len(largeBody) {
		t.Fatalf("Body length mismatch: got %d, want %d", len(msg.Body), len(largeBody))
	}
	if !bytes.Equal(msg.Body, largeBody) {
		t.Fatalf("Body content corrupted after slab roundtrip")
	}
}

// TestDistributed_ConsumerGroup_DurableOffset_ProcessRestart teste la persistance réelle sur disque des offsets et la reprise après redémarrage complet.
func TestDistributed_ConsumerGroup_DurableOffset_ProcessRestart(t *testing.T) {
	tmpDir := t.TempDir()
	offsetDir := filepath.Join(tmpDir, "offsets_store")

	// Phase 1 : Initialisation de l'Engine 1 avec stockage persistant
	opts1 := DefaultOptions()
	opts1.NumShards = 4
	opts1.ShardCapacity = 1024
	opts1.OffsetDir = offsetDir

	eng1, err := NewEngine(opts1)
	if err != nil {
		t.Fatalf("Engine 1 start failed: %v", err)
	}

	group1 := eng1.GetOrCreateConsumerGroup("durable-order-workers")
	group1.RegisterConsumer("worker-1")

	// Publication de 100 messages
	for i := 0; i < 100; i++ {
		_ = eng1.Publish(uint64(i), uint64(i), 1, 0, []byte("ord-msg"))
	}

	// Consommation de 45 messages par le worker-1 et validation des offsets
	var msg Message
	for i := 0; i < 45; i++ {
		if ok, _ := eng1.ConsumeGroup("durable-order-workers", "worker-1", &msg); ok {
			_ = group1.CommitOffset(msg.Partition, uint64(i+1))
		}
	}

	lastP0 := group1.GetCommittedOffset(0)

	// Arrêt total de l'Engine 1 (simulation de fermeture/crash de processus)
	_ = eng1.Close()

	// Phase 2 : Démarrage d'un NOUVEL Engine 2 sur les MÊMES fichiers d'offsets
	opts2 := DefaultOptions()
	opts2.NumShards = 4
	opts2.ShardCapacity = 1024
	opts2.OffsetDir = offsetDir

	eng2, err := NewEngine(opts2)
	if err != nil {
		t.Fatalf("Engine 2 start failed: %v", err)
	}
	defer eng2.Close()

	// Création du groupe sur l'Engine 2 : les offsets doivent être rechargés du disque
	group2 := eng2.GetOrCreateConsumerGroup("durable-order-workers")

	loadedP0 := group2.GetCommittedOffset(0)
	if loadedP0 != lastP0 {
		t.Fatalf("DURABILITY FAILURE: Reopened group offset mismatch: got %d, want %d", loadedP0, lastP0)
	}
}

// TestDistributed_Slab_WatermarkProtection teste la protection absolue contre l'écrasement de corps vivants.
func TestDistributed_Slab_WatermarkProtection(t *testing.T) {
	tmpDir := t.TempDir()
	slabPath := filepath.Join(tmpDir, "watermark_test.slab")

	// Petite arène de 128 Ko
	const slabSize int64 = 128 * 1024
	slab, err := OpenSlabArena(slabPath, slabSize)
	if err != nil {
		t.Fatalf("OpenSlabArena failed: %v", err)
	}
	defer slab.Close()

	// Écriture de 3 blocs de 40 Ko (120 Ko utiles + alignements)
	body40K := MakeRealisticDocumentPayload(40 * 1024)
	desc1, err := slab.Write(body40K)
	if err != nil {
		t.Fatalf("Write block 1 failed: %v", err)
	}

	_, err = slab.Write(body40K)
	if err != nil {
		t.Fatalf("Write block 2 failed: %v", err)
	}

	_, err = slab.Write(body40K)
	if err != nil {
		t.Fatalf("Write block 3 failed: %v", err)
	}

	// 4ème écriture : l'arène est pleine, la ligne d'eau basse est à 0.
	// Doit retourner ErrSlabFull pour protéger les corps vivants (zéro écrasement silencieux) !
	_, err = slab.Write(body40K)
	if err != ErrSlabFull {
		t.Fatalf("Expected ErrSlabFull when capacity reached, got: %v", err)
	}

	// Relecture du bloc 1 pour prouver qu'il est 100% intact et non corrompu
	read1, err := slab.Read(desc1)
	if err != nil || !bytes.Equal(read1, body40K) {
		t.Fatalf("Block 1 corrupted by subsequent writes")
	}

	// Recyclage : on avance la ligne d'eau basse (les 40 premiers Ko ont été consommés et acquittés)
	slab.AdvanceWatermark(uint64(desc1.Length + 64))

	// L'écriture suivante doit maintenant réussir !
	_, err = slab.Write(body40K)
	if err != nil {
		t.Fatalf("Write failed after advancing watermark: %v", err)
	}
}

// TestDistributed_CompactTable_ReplayWAL teste la reconstruction d'état bit-exacte par rejeu de log.
func TestDistributed_CompactTable_ReplayWAL(t *testing.T) {
	tmpDir := t.TempDir()
	walPath := filepath.Join(tmpDir, "state_rebuild.wal")
	slabPath := filepath.Join(tmpDir, "state_rebuild.slab")

	opts := DefaultOptions()
	opts.NumShards = 4
	opts.WALPath = walPath
	opts.SlabPath = slabPath

	eng, err := NewEngine(opts)
	if err != nil {
		t.Fatalf("NewEngine failed: %v", err)
	}

	const numKeys = 100
	expectedPayloads := make([][]byte, numKeys)
	for i := 0; i < numKeys; i++ {
		key := uint64(1000 + i)
		val := MakeRealisticTelemetryPayload(i)
		expectedPayloads[i] = val
		_ = eng.Publish(key, key, 1, 0, val)
	}

	_ = eng.Sync()
	_ = eng.Close()

	// Réouverture d'une CompactTable vierge et reconstruction intégrale par rejeu WAL
	slab, err := OpenSlabArena(slabPath, 64*1024*1024)
	if err != nil {
		t.Fatalf("OpenSlabArena failed: %v", err)
	}
	defer slab.Close()

	freshTable := NewCompactTable(4, slab)
	replayed, err := freshTable.ReplayWAL(walPath, 64*1024*1024)
	if err != nil {
		t.Fatalf("ReplayWAL failed: %v", err)
	}
	if replayed != numKeys {
		t.Fatalf("Replayed keys mismatch: got %d, want %d", replayed, numKeys)
	}

	for i := 0; i < numKeys; i++ {
		key := uint64(1000 + i)
		expectedVal := expectedPayloads[i]
		recoveredVal, ok := freshTable.Get(key)
		if !ok {
			t.Fatalf("Key %d missing after WAL replay", key)
		}
		if !bytes.Equal(recoveredVal, expectedVal) {
			t.Fatalf("Key %d value mismatch after WAL replay", key)
		}
	}
}

// TestDistributed_QUICTransport_1MB_Over_Stream teste l'envoi de 100 messages 4 Ko + 1 message 1 Mo via quic-go.
func TestDistributed_QUICTransport_1MB_Over_Stream(t *testing.T) {
	tmpDir := t.TempDir()
	slabPath := filepath.Join(tmpDir, "quic_engine.slab")

	opts := DefaultOptions()
	opts.NumShards = 4
	opts.ShardCapacity = 2048
	opts.SlabPath = slabPath
	opts.SlabSizeBytes = 32 * 1024 * 1024

	eng, err := NewEngine(opts)
	if err != nil {
		t.Fatalf("NewEngine failed: %v", err)
	}
	defer eng.Close()

	table := NewCompactTable(4, eng.Slab())

	server, err := ListenQUICTransport("127.0.0.1:0", eng, table)
	if err != nil {
		t.Fatalf("ListenQUICTransport failed: %v", err)
	}
	defer server.Close()

	serverAddr := server.Addr()

	client, err := DialQUICTransport(serverAddr)
	if err != nil {
		t.Fatalf("DialQUICTransport failed: %v", err)
	}
	defer client.Close()

	// 1. Émettre 100 messages de 4 Ko
	payload4KB := MakeRealisticDocumentPayload(4096)
	for i := 0; i < 100; i++ {
		err := client.SendPublish(uint16(i%4), uint64(i), uint64(time.Now().UnixNano()), payload4KB, true)
		if err != nil {
			t.Fatalf("SendPublish 4KB failed at %d: %v", i, err)
		}
	}

	// 2. Émettre 1 message de 1 Mo (1 048 576 octets)
	payload1MB := MakeRealisticDocumentPayload(1024 * 1024)
	err = client.SendPublish(0, 0xCAFEBABE, uint64(time.Now().UnixNano()), payload1MB, true)
	if err != nil {
		t.Fatalf("SendPublish 1MB failed: %v", err)
	}

	if server.Received() < 101 {
		t.Fatalf("Server received count mismatch: got %d, want >= 101", server.Received())
	}

	// 3. Consommation et validation bit-exacte des messages reçus
	var msg Message
	count4KB := 0
	found1MB := false

	for i := 0; i < 101; i++ {
		if eng.ConsumeMessage(&msg) {
			if len(msg.Body) == 4096 && bytes.Equal(msg.Body, payload4KB) {
				count4KB++
			} else if len(msg.Body) == 1024*1024 && bytes.Equal(msg.Body, payload1MB) {
				found1MB = true
			}
		}
	}

	if count4KB != 100 {
		t.Fatalf("4KB messages validated mismatch: got %d, want 100", count4KB)
	}
	if !found1MB {
		t.Fatalf("1MB payload not found or corrupted over QUIC stream")
	}
}
