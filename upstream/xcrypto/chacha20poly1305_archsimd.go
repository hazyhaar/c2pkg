// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build archsimd && goexperiment.simd && amd64 && !purego

package chacha20poly1305

import (
	c2simd "github.com/hazyhaar/c2pkg/c2fused"
	"golang.org/x/crypto/internal/alias"
	"golang.org/x/crypto/internal/poly1305"
)

// Greffe archsimd : le chemin IETF (nonce 12 octets) est délégué au noyau
// pur Go AVX2 de c2simd ; XChaCha20 continue de passer par HChaCha20 puis
// par ce chemin, exactement comme pour l'assembleur et le générique.
// Ce fichier remplace chacha20poly1305_amd64.go sous le tag `archsimd`
// (le tag de l'assembleur est complété de `&& !archsimd` par harness.sh).

func (c *chacha20poly1305) seal(dst, nonce, plaintext, additionalData []byte) []byte {
	ret, out := sliceForAppend(dst, len(plaintext)+poly1305.TagSize)
	ciphertext, tag := out[:len(plaintext)], out[len(plaintext):]
	if alias.InexactOverlap(out, plaintext) {
		panic("chacha20poly1305: invalid buffer overlap of output and input")
	}
	if alias.AnyOverlap(out, additionalData) {
		panic("chacha20poly1305: invalid buffer overlap of output and additional data")
	}
	var mac [16]byte
	if _, err := c2simd.AEADSubkeyLockDst(ciphertext, &mac, c.key[:], nonce, additionalData, plaintext); err != nil {
		panic("chacha20poly1305: archsimd seal: " + err.Error())
	}
	copy(tag, mac[:])
	return ret
}

func (c *chacha20poly1305) open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error) {
	tag := ciphertext[len(ciphertext)-16:]
	ciphertext = ciphertext[:len(ciphertext)-16]

	ret, out := sliceForAppend(dst, len(ciphertext))
	if alias.InexactOverlap(out, ciphertext) {
		panic("chacha20poly1305: invalid buffer overlap of output and input")
	}
	if alias.AnyOverlap(out, additionalData) {
		panic("chacha20poly1305: invalid buffer overlap of output and additional data")
	}
	if _, err := c2simd.AEADSubkeyUnlockDst(out, c.key[:], nonce, additionalData, ciphertext, tag); err != nil {
		for i := range out {
			out[i] = 0
		}
		return nil, errOpen
	}
	return ret, nil
}
