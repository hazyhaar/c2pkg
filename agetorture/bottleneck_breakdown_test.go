// SPDX-License-Identifier: Apache-2.0 OR MIT

package agetorture

import (
	"crypto/rand"
	"encoding/binary"
	"testing"

	"github.com/hazyhaar/c2pkg/c2chacha8"
	"github.com/hazyhaar/c2pkg/c2fused"
	"github.com/hazyhaar/c2pkg/c2poly1305"
)

// BenchmarkBottleneckBreakdown_64KiB décompose le temps exact passé dans chaque sous-composant
// sur un bloc nominal age de 64 KiB :
// 1. Dérivation de sous-clé (ChaCha block 0)
// 2. Chiffrement Keystream ChaCha20 pur
// 3. Authentification Poly1305 pure
// 4. Moteur Fusionné (ChaCha + Poly entrelacés)
func BenchmarkBottleneckBreakdown_64KiB(b *testing.B) {
	key := make([]byte, 32)
	rand.Read(key)
	payload := make([]byte, 64*1024)
	dst := make([]byte, 64*1024)
	rand.Read(payload)

	var nonce12 [12]byte
	binary.BigEndian.PutUint64(nonce12[3:11], 0)
	nonce12[11] = 0x01

	var polyKey [64]byte
	c2chacha8.C2chacha8_xor_blocks(polyKey[:], polyKey[:], 64, key, nonce12[:], 0)

	b.Run("1_Subkey_Derive_64B", func(b *testing.B) {
		b.SetBytes(64)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			c2chacha8.C2chacha8_xor_blocks(polyKey[:], polyKey[:], 64, key, nonce12[:], 0)
		}
	})

	b.Run("2_ChaCha20_Keystream_Only_64KiB", func(b *testing.B) {
		b.SetBytes(int64(len(payload)))
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			c2chacha8.C2chacha8_xor_blocks(dst, payload, uint64(len(payload)), key, nonce12[:], 1)
		}
	})

	b.Run("3_Poly1305_Auth_Only_64KiB", func(b *testing.B) {
		b.SetBytes(int64(len(payload)))
		b.ReportAllocs()
		var ctx c2poly1305.Crypto_poly1305_ctx
		var mac [16]byte
		for i := 0; i < b.N; i++ {
			c2poly1305.Crypto_poly1305_init(&ctx, polyKey[:32])
			c2poly1305.Crypto_poly1305_update(&ctx, payload, uint64(len(payload)))
			c2poly1305.Crypto_poly1305_final(&ctx, mac[:])
		}
	})

	b.Run("4_Fused_Poly_Update_Only_64KiB", func(b *testing.B) {
		b.SetBytes(int64(len(payload)))
		b.ReportAllocs()
		var st c2fused.C2fused_poly
		var mac [16]byte
		for i := 0; i < b.N; i++ {
			c2fused.C2fused_poly_init(&st, polyKey[:32])
			c2fused.C2fused_poly_update(&st, payload, uint64(len(payload)))
			c2fused.C2fused_poly_final(&st, mac[:])
		}
	})

	b.Run("5_Fused_Total_Lock_64KiB", func(b *testing.B) {
		b.SetBytes(int64(len(payload)))
		b.ReportAllocs()
		var st c2fused.C2fused_poly
		var mac [16]byte
		for i := 0; i < b.N; i++ {
			c2chacha8.C2chacha8_xor_blocks(polyKey[:], polyKey[:], 64, key, nonce12[:], 0)
			c2fused.C2fused_poly_init(&st, polyKey[:32])
			c2chacha8.C2chacha8_xor_blocks(dst, payload, uint64(len(payload)), key, nonce12[:], 1)
			c2fused.C2fused_poly_update(&st, dst, uint64(len(dst)))
			var lenBlock [16]byte
			c2fused.C2fused_poly_update(&st, lenBlock[:], 16)
			c2fused.C2fused_poly_final(&st, mac[:])
		}
	})
}
