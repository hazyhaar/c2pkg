// SPDX-License-Identifier: Apache-2.0 OR MIT

package agetorture

import (
	"crypto/rand"
	"encoding/binary"
	"runtime"
	"testing"

	"github.com/hazyhaar/c2pkg/c2chacha8"
	"github.com/hazyhaar/c2pkg/c2fused"
	"golang.org/x/crypto/chacha20poly1305"
)

// TestPerfHardwareCounters_MicroarchProbe exécute perf stat sur les deux moteurs
// pour quantifier précisément : IPC, Cache Misses L1D/LLC, Branch Misses, Stalls d'exécution et Déversements.
func TestPerfHardwareCounters_MicroarchProbe(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping hardware counter probe in short mode")
	}

	key := make([]byte, 32)
	rand.Read(key)
	payload := make([]byte, 64*1024)
	rand.Read(payload)

	t.Run("AgeStandard_xCrypto_Loop", func(t *testing.T) {
		aead, err := chacha20poly1305.New(key)
		if err != nil {
			t.Fatal(err)
		}
		var chunkNonce [12]byte
		binary.BigEndian.PutUint64(chunkNonce[3:11], 1)
		chunkNonce[11] = 0x01

		for i := 0; i < 50000; i++ {
			_ = aead.Seal(nil, chunkNonce[:], payload, nil)
		}
	})

	t.Run("C2Fused_SIMD_Loop", func(t *testing.T) {
		var polyKeyBlock [64]byte
		var st c2fused.C2fused_poly
		var mac [16]byte
		var nonce12 [12]byte
		dstCipher := make([]byte, len(payload))

		for i := 0; i < 50000; i++ {
			c2chacha8.C2chacha8_xor_blocks(polyKeyBlock[:], polyKeyBlock[:], 64, key, nonce12[:], 0)
			c2fused.C2fused_poly_init(&st, polyKeyBlock[:32])
			c2chacha8.C2chacha8_xor_blocks(dstCipher, payload, uint64(len(payload)), key, nonce12[:], 1)
			c2fused.C2fused_poly_update(&st, dstCipher, uint64(len(dstCipher)))
			var lenBlock [16]byte
			c2fused.C2fused_poly_update(&st, lenBlock[:], 16)
			c2fused.C2fused_poly_final(&st, mac[:])
		}
		runtime.KeepAlive(dstCipher)
	})
}
