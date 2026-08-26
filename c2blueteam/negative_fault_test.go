// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2blueteam

import (
	"testing"
)

// TestNegativeFaultInjection_NullGuards valide formellement la résistance aux pointeurs nuls
// et entrées corrompues sans aucune panique Go (fermeture de P0-01, P0-02, P1-03, P1-16).
func TestNegativeFaultInjection_NullGuards(t *testing.T) {
	// 1. Canal unitaire - pointeurs nuls
	if res := C2bt_channel_write(nil, nil); res != -1 {
		t.Errorf("C2bt_channel_write(nil, nil) = %d, attendu -1", res)
	}
	var ev Probe_event_t
	if res := C2bt_channel_write(nil, &ev); res != -1 {
		t.Errorf("C2bt_channel_write(nil, &ev) = %d, attendu -1", res)
	}
	var ch Probe_channel_t
	if res := C2bt_channel_write(&ch, nil); res != -1 {
		t.Errorf("C2bt_channel_write(&ch, nil) = %d, attendu -1", res)
	}

	if res := C2bt_channel_read(nil, nil); res != -1 {
		t.Errorf("C2bt_channel_read(nil, nil) = %d, attendu -1", res)
	}
	if res := C2bt_channel_read(&ch, nil); res != -1 {
		t.Errorf("C2bt_channel_read(&ch, nil) = %d, attendu -1", res)
	}
	if res := C2bt_channel_read(nil, &ev); res != -1 {
		t.Errorf("C2bt_channel_read(nil, &ev) = %d, attendu -1", res)
	}

	// 2. Lecture par lot - bornes invalides et négatives
	outBatch := make([]Probe_event_t, 10)
	if res := C2bt_channel_read_batch(nil, outBatch, 10); res != 0 {
		t.Errorf("C2bt_channel_read_batch(nil, ...) = %d, attendu 0", res)
	}
	if res := C2bt_channel_read_batch(&ch, nil, 10); res != 0 {
		t.Errorf("C2bt_channel_read_batch(&ch, nil, ...) = %d, attendu 0", res)
	}
	if res := C2bt_channel_read_batch(&ch, outBatch, 0); res != 0 {
		t.Errorf("C2bt_channel_read_batch(&ch, out, 0) = %d, attendu 0", res)
	}
	if res := C2bt_channel_read_batch(&ch, outBatch, -5); res != 0 {
		t.Errorf("C2bt_channel_read_batch(&ch, out, -5) = %d, attendu 0", res)
	}

	// 3. Calculateur d'entropie - tranche nulle, taille nulle et dépassement de borne
	if ent := C2bt_calc_entropy_8_8(nil, 0); ent != 0 {
		t.Errorf("C2bt_calc_entropy_8_8(nil, 0) = %d, attendu 0", ent)
	}
	if ent := C2bt_calc_entropy_8_8([]byte{}, 0); ent != 0 {
		t.Errorf("C2bt_calc_entropy_8_8([], 0) = %d, attendu 0", ent)
	}
	if ent := C2bt_calc_entropy_8_8([]byte("test"), 0); ent != 0 {
		t.Errorf("C2bt_calc_entropy_8_8(data, 0) = %d, attendu 0", ent)
	}
	// Épreuve de résistance au débordement (len_ > len(data)) sans panique
	smallBuf := []byte("secret")
	if ent := C2bt_calc_entropy_8_8(smallBuf, 1000); ent == 0 {
		t.Errorf("C2bt_calc_entropy_8_8(smallBuf, 1000) = 0, attendu calcul valide sans panique")
	}
	var prof C2bt_entropy_profile_t
	if res := C2bt_profile_payload(smallBuf, 500, &prof); res != 0 || prof.Len_ != uint64(len(smallBuf)) {
		t.Errorf("C2bt_profile_payload(smallBuf, 500) : res=%d, len=%d (attendu %d)", res, prof.Len_, len(smallBuf))
	}

	// 4. Moteur de règles - tranches nulles, count négatif et dépassement de bornes
	if count := C2bt_eval_rules_batch(outBatch[:2], outBatch[:2], 100); count < 0 {
		t.Errorf("C2bt_eval_rules_batch(short, short, 100) = %d", count)
	}
	if count := C2bt_eval_rules_batch(nil, nil, 0); count != 0 {
		t.Errorf("C2bt_eval_rules_batch(nil, nil, 0) = %d, attendu 0", count)
	}
	if count := C2bt_eval_rules_batch(outBatch, nil, 10); count != 0 {
		t.Errorf("C2bt_eval_rules_batch(out, nil, 10) = %d, attendu 0", count)
	}
	if count := C2bt_eval_rules_batch(nil, outBatch, 10); count != 0 {
		t.Errorf("C2bt_eval_rules_batch(nil, out, 10) = %d, attendu 0", count)
	}
	if count := C2bt_eval_rules_batch(outBatch, outBatch, -1); count != 0 {
		t.Errorf("C2bt_eval_rules_batch(out, out, -1) = %d, attendu 0", count)
	}

	// 5. Contexte - pointeurs nuls
	if res := C2bt_init_inplace(nil, nil); res != -1 {
		t.Errorf("C2bt_init_inplace(nil, nil) = %d, attendu -1", res)
	}
	if res := C2bt_start(nil); res != -1 {
		t.Errorf("C2bt_start(nil) = %d, attendu -1", res)
	}
	if res := C2bt_stop(nil); res != -1 {
		t.Errorf("C2bt_stop(nil) = %d, attendu -1", res)
	}
}

// TestZeroAllocationPath vérifie strictement la doctrine d'allocation zéro (0 B/op) sur le chemin chaud
func TestZeroAllocationPath(t *testing.T) {
	var ch Probe_channel_t
	C2bt_channel_init(&ch)

	var ev Probe_event_t
	ev.Subsystem = C2BT_SUB_PROC
	ev.Action = C2BT_ACT_EXEC
	copy(ev.Payload[:], "ls")

	// 1. Vérification 0 allocation sur écriture canal SPSC
	allocsWrite := testing.AllocsPerRun(1000, func() {
		C2bt_channel_write(&ch, &ev)
		var outEv Probe_event_t
		C2bt_channel_read(&ch, &outEv)
	})
	if allocsWrite != 0 {
		t.Errorf("C2bt_channel_write/read alloue %.2f allocs/op, attendu 0.0", allocsWrite)
	}

	// 2. Vérification 0 allocation sur calcul d'entropie
	buf := make([]byte, 256)
	for i := range buf {
		buf[i] = byte(i)
	}
	allocsEntropy := testing.AllocsPerRun(1000, func() {
		_ = C2bt_calc_entropy_8_8(buf, 256)
	})
	if allocsEntropy != 0 {
		t.Errorf("C2bt_calc_entropy_8_8 alloue %.2f allocs/op, attendu 0.0", allocsEntropy)
	}

	// 3. Vérification 0 allocation sur évaluation de règles par lot
	inBatch := make([]Probe_event_t, 8)
	outBatch := make([]Probe_event_t, 8)
	for i := range inBatch {
		inBatch[i].Subsystem = C2BT_SUB_PROC
		inBatch[i].Action = C2BT_ACT_EXEC
		copy(inBatch[i].Payload[:], "/usr/bin/curl -O https://example.com")
	}
	allocsRules := testing.AllocsPerRun(1000, func() {
		_ = C2bt_eval_rules_batch(inBatch, outBatch, 8)
	})
	if allocsRules != 0 {
		t.Errorf("C2bt_eval_rules_batch alloue %.2f allocs/op, attendu 0.0", allocsRules)
	}
}
