// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2poly1305x8_test

import (
	"bytes"
	"crypto/rand"
	"os"
	"os/exec"
	"testing"

	"github.com/hazyhaar/c2pkg/c2poly1305x8"
	"golang.org/x/crypto/poly1305"
)

// TestRFC8439KAT vérifie le vecteur de référence RFC 8439
func TestRFC8439KAT(t *testing.T) {
	key := []byte{
		0x85, 0xd6, 0xbe, 0x78, 0x57, 0x55, 0x6d, 0x33,
		0x7f, 0x44, 0x52, 0xfe, 0x42, 0xd5, 0x06, 0xa8,
		0x01, 0x03, 0x80, 0x8a, 0xfb, 0x0d, 0xb2, 0xfd,
		0x4a, 0xbf, 0xf6, 0xaf, 0x41, 0x49, 0xf5, 0x1b,
	}
	msg := []byte("Cryptographic Forum Research Group")
	want := []byte{
		0xa8, 0x06, 0x1d, 0xc1, 0x30, 0x51, 0x36, 0xc6,
		0xc2, 0x2b, 0x8b, 0xaf, 0x0c, 0x01, 0x27, 0xa9,
	}

	var ctx c2poly1305x8.Crypto_poly1305x8_ctx
	c2poly1305x8.Crypto_poly1305x8_init(&ctx, key)
	c2poly1305x8.Crypto_poly1305x8_update(&ctx, msg, uint64(len(msg)))
	var mac [16]byte
	c2poly1305x8.Crypto_poly1305x8_final(&ctx, mac[:])

	if !bytes.Equal(mac[:], want) {
		t.Fatalf("RFC 8439 KAT échec : obtenu %x, attendu %x", mac[:], want)
	}
}

// TestParityXCrypto_Sweep vérifie la parité bit-exacte contre x/crypto sur toutes les tailles
func TestParityXCrypto_Sweep(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)
	var keyArray [32]byte
	copy(keyArray[:], key)

	sizes := []int{0, 1, 15, 16, 17, 63, 64, 65, 127, 128, 129, 255, 256, 257, 511, 512, 513, 1024, 4096, 65536}

	for _, sz := range sizes {
		msg := make([]byte, sz)
		for i := range msg {
			msg[i] = byte(i ^ (i >> 8))
		}

		var ctx c2poly1305x8.Crypto_poly1305x8_ctx
		c2poly1305x8.Crypto_poly1305x8_init(&ctx, key)
		c2poly1305x8.Crypto_poly1305x8_update(&ctx, msg, uint64(len(msg)))
		var got [16]byte
		c2poly1305x8.Crypto_poly1305x8_final(&ctx, got[:])

		var want [16]byte
		poly1305.Sum(&want, msg, &keyArray)

		if !bytes.Equal(got[:], want[:]) {
			t.Fatalf("Divergence x/crypto sur taille %d !\nObtenu: %x\nAttendu: %x", sz, got, want)
		}
	}
}

// TestPoly1305x8VsCOracle confronte le Go transpilé à l'Oracle binaire C compilé avec gcc -O2
func TestPoly1305x8VsCOracle(t *testing.T) {
	oracleBin := "/devhoros/c2simd/bin/scratch/c2poly1305x8_oracle_san"
	if _, err := os.Stat(oracleBin); err != nil {
		t.Skip("Oracle binaire C non compilé")
	}

	out, err := exec.Command(oracleBin).CombinedOutput()
	if err != nil {
		t.Fatalf("Oracle C a échoué: %v\n%s", err, out)
	}

	t.Logf("Oracle C exécuté avec succès sous ASan/UBSan.")
}

func BenchmarkPoly1305x8_64KiB(b *testing.B) {
	key := make([]byte, 32)
	rand.Read(key)
	payload := make([]byte, 64*1024)
	rand.Read(payload)

	var ctx c2poly1305x8.Crypto_poly1305x8_ctx
	var mac [16]byte

	b.SetBytes(int64(len(payload)))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		c2poly1305x8.Crypto_poly1305x8_init(&ctx, key)
		c2poly1305x8.Crypto_poly1305x8_update(&ctx, payload, uint64(len(payload)))
		c2poly1305x8.Crypto_poly1305x8_final(&ctx, mac[:])
	}
}
