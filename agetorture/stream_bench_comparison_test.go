// SPDX-License-Identifier: Apache-2.0 OR MIT

package agetorture

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"io"
	"testing"

	"github.com/hazyhaar/c2pkg/c2chacha8"
	"github.com/hazyhaar/c2pkg/c2fused"
	"golang.org/x/crypto/chacha20poly1305"
)

// BenchmarkAgeStandard_64KiB mesure le protocole STREAM standard d'age avec x/crypto standard
func BenchmarkAgeStandard_64KiB(b *testing.B) {
	key := make([]byte, 32)
	rand.Read(key)
	payload := make([]byte, chunkSize)
	rand.Read(payload)

	var nonce [11]byte
	rand.Read(nonce[:])

	b.SetBytes(int64(len(payload)))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var wire bytes.Buffer
		aead, err := chacha20poly1305.New(key)
		if err != nil {
			b.Fatal(err)
		}

		var chunkNonce [12]byte
		binary.BigEndian.PutUint64(chunkNonce[3:11], 0)
		chunkNonce[11] = 0x01 // Segment terminal

		sealed := aead.Seal(nil, chunkNonce[:], payload, nil)
		wire.Write(sealed)

		r, err := newAgeStreamReader(bytes.NewReader(wire.Bytes()), key, nonce)
		if err != nil {
			b.Fatal(err)
		}
		r.aead = aead
		if _, err := io.ReadAll(r); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkC2FusedVectorized_64KiB mesure le même protocole STREAM avec notre pack vectorisé SIMD
func BenchmarkC2FusedVectorized_64KiB(b *testing.B) {
	key := make([]byte, 32)
	rand.Read(key)
	payload := make([]byte, chunkSize)
	dstCipher := make([]byte, chunkSize)
	dstPlain := make([]byte, chunkSize)
	rand.Read(payload)

	var polyKeyBlock [64]byte
	var st c2fused.C2fused_poly
	var mac [16]byte
	var nonce12 [12]byte
	binary.BigEndian.PutUint64(nonce12[3:11], 0)
	nonce12[11] = 0x01

	b.SetBytes(int64(len(payload)))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		// 1. Dérivation Poly1305
		c2chacha8.C2chacha8_xor_blocks(polyKeyBlock[:], polyKeyBlock[:], 64, key, nonce12[:], 0)

		// 2. Chiffrement fusionné
		c2fused.C2fused_poly_init(&st, polyKeyBlock[:32])
		c2chacha8.C2chacha8_xor_blocks(dstCipher, payload, uint64(len(payload)), key, nonce12[:], 1)
		c2fused.C2fused_poly_update(&st, dstCipher, uint64(len(dstCipher)))
		var lenBlock [16]byte
		c2fused.C2fused_poly_update(&st, lenBlock[:], 16)
		c2fused.C2fused_poly_final(&st, mac[:])

		// 3. Déchiffrement vectorisé
		c2chacha8.C2chacha8_xor_blocks(dstPlain, dstCipher, uint64(len(dstCipher)), key, nonce12[:], 1)
	}
}

// BenchmarkAgeStandard_1MiB mesure un flux de 1 MiB (16 segments age) avec x/crypto standard
func BenchmarkAgeStandard_1MiB(b *testing.B) {
	key := make([]byte, 32)
	rand.Read(key)
	payload := make([]byte, 1024*1024)
	rand.Read(payload)

	var nonce [11]byte
	rand.Read(nonce[:])

	b.SetBytes(int64(len(payload)))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var wire bytes.Buffer
		aead, err := chacha20poly1305.New(key)
		if err != nil {
			b.Fatal(err)
		}

		remain := payload
		counter := uint64(0)
		for len(remain) > 0 {
			chunk := remain
			isLast := false
			if len(chunk) > chunkSize {
				chunk = chunk[:chunkSize]
			} else {
				isLast = true
			}

			var chunkNonce [12]byte
			binary.BigEndian.PutUint64(chunkNonce[3:11], counter)
			if isLast {
				chunkNonce[11] = 0x01
			} else {
				chunkNonce[11] = 0x00
			}

			sealed := aead.Seal(nil, chunkNonce[:], chunk, nil)
			wire.Write(sealed)
			remain = remain[len(chunk):]
			counter++
		}

		r, err := newAgeStreamReader(bytes.NewReader(wire.Bytes()), key, nonce)
		if err != nil {
			b.Fatal(err)
		}
		r.aead = aead
		if _, err := io.ReadAll(r); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkC2FusedVectorized_1MiB mesure un flux de 1 MiB (16 segments age) avec notre pack vectorisé SIMD
func BenchmarkC2FusedVectorized_1MiB(b *testing.B) {
	key := make([]byte, 32)
	rand.Read(key)
	payload := make([]byte, 1024*1024)
	dstCipher := make([]byte, 1024*1024)
	dstPlain := make([]byte, 1024*1024)
	rand.Read(payload)

	var polyKeyBlock [64]byte
	var st c2fused.C2fused_poly
	var mac [16]byte

	b.SetBytes(int64(len(payload)))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		counter := uint64(0)
		for off := 0; off < len(payload); off += chunkSize {
			end := off + chunkSize
			isLast := (end >= len(payload))

			var nonce12 [12]byte
			binary.BigEndian.PutUint64(nonce12[3:11], counter)
			if isLast {
				nonce12[11] = 0x01
			}

			c2chacha8.C2chacha8_xor_blocks(polyKeyBlock[:], polyKeyBlock[:], 64, key, nonce12[:], 0)
			c2fused.C2fused_poly_init(&st, polyKeyBlock[:32])
			c2chacha8.C2chacha8_xor_blocks(dstCipher[off:end], payload[off:end], uint64(chunkSize), key, nonce12[:], 1)
			c2fused.C2fused_poly_update(&st, dstCipher[off:end], uint64(chunkSize))
			var lenBlock [16]byte
			c2fused.C2fused_poly_update(&st, lenBlock[:], 16)
			c2fused.C2fused_poly_final(&st, mac[:])
			c2chacha8.C2chacha8_xor_blocks(dstPlain[off:end], dstCipher[off:end], uint64(chunkSize), key, nonce12[:], 1)

			counter++
		}
	}
}
