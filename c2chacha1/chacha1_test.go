// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2chacha1

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"golang.org/x/crypto/chacha20"
)

// Vecteur RFC 8439 §2.4.2, bloc 1 (64 premiers octets du chiffré).
func TestRFC8439Block1(t *testing.T) {
	key := make([]byte, 32)
	for i := range key {
		key[i] = byte(i)
	}
	nonce := make([]byte, 12)
	nonce[7] = 0x4a
	msg := []byte("Ladies and Gentlemen of the class of '99: If I could offer you only one tip for the future, sunscreen would be it.")
	want, _ := hex.DecodeString("6e2e359a2568f98041ba0728dd0d6981e97e7aec1d4360c20a27afccfd9fae0bf91b65c5524733ab8f593dabcd62b3571639d624e65152ab8f530c359f0861d8")
	out := make([]byte, 64)
	if got := C2chacha1_xor_block(out, msg[:64], 64, key, nonce, 1); got != 64 {
		t.Fatalf("n rendu %d", got)
	}
	if !bytes.Equal(out, want) {
		t.Fatalf("RFC 8439 : obtenu %x", out)
	}
}

// Parité contre golang.org/x/crypto/chacha20 pour toutes les longueurs 0..64,
// clé, nonce, compteur et message aléatoires ; en place et hors place.
func TestParityXCrypto(t *testing.T) {
	for n := 0; n <= 64; n++ {
		for rep := 0; rep < 8; rep++ {
			key := make([]byte, 32)
			nonce := make([]byte, 12)
			msg := make([]byte, n)
			var ctrB [4]byte
			rand.Read(key)
			rand.Read(nonce)
			rand.Read(msg)
			rand.Read(ctrB[:])
			ctr := uint32(ctrB[0]) | uint32(ctrB[1])<<8 | uint32(ctrB[2])<<16 | uint32(ctrB[3])<<24

			c, err := chacha20.NewUnauthenticatedCipher(key, nonce)
			if err != nil {
				t.Fatal(err)
			}
			c.SetCounter(ctr)
			want := make([]byte, n)
			c.XORKeyStream(want, msg)

			out := make([]byte, n)
			C2chacha1_xor_block(out, msg, uint64(n), key, nonce, ctr)
			if !bytes.Equal(out, want) {
				t.Fatalf("n=%d ctr=%d : %x != %x", n, ctr, out, want)
			}
			inPlace := append([]byte(nil), msg...)
			C2chacha1_xor_block(inPlace, inPlace, uint64(n), key, nonce, ctr)
			if !bytes.Equal(inPlace, want) {
				t.Fatalf("n=%d en place : divergence", n)
			}
		}
	}
}

type lcg struct{ s uint64 }

func (l *lcg) seed(n uint64) { l.s = n*0x9E3779B97F4A7C15 + 1 }
func (l *lcg) next() byte {
	l.s = l.s*6364136223846793005 + 1442695040888963407
	return byte(l.s >> 56)
}

// Oracle gcc -O2 : la même source C que celle transpilée, compilée nativement.
func TestChacha1VsCOracle(t *testing.T) {
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("gcc absent : oracle C non jouable")
	}
	srcDir := filepath.Join("..", "..", "sources", "c2chacha1")
	if _, err := os.Stat(filepath.Join(srcDir, "c2chacha1_simd.c")); err != nil {
		t.Skipf("source oracle absente : %v", err)
	}
	tmpBin := filepath.Join(t.TempDir(), "c2chacha1_oracle")
	build := exec.Command("gcc", "-O2", "-mssse3", "-Wall", "-Wextra", "-std=gnu99", "-I", srcDir,
		filepath.Join(srcDir, "test_c2chacha1_oracle.c"), filepath.Join(srcDir, "c2chacha1_simd.c"), "-o", tmpBin)
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("compilation oracle C : %v\n%s", err, out)
	}
	out, err := exec.Command(tmpBin).Output()
	if err != nil {
		t.Fatalf("exécution oracle C : %v", err)
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) != 13 {
		t.Fatalf("oracle C : %d lignes, attendu 13", len(lines))
	}
	var g lcg
	for _, line := range lines[1:] {
		var n int
		var hexout string
		if _, err := fmt.Sscanf(line+" ", "%d:%s", &n, &hexout); err != nil && n != 0 {
			t.Fatalf("ligne oracle illisible %q : %v", line, err)
		}
		g.seed(uint64(n))
		key := make([]byte, 32)
		for k := range key {
			key[k] = g.next()
		}
		nonce := make([]byte, 12)
		for k := range nonce {
			nonce[k] = g.next()
		}
		ctr := uint32(g.next()) | uint32(g.next())<<8
		in := make([]byte, n)
		for k := range in {
			in[k] = g.next()
		}
		got := make([]byte, n)
		C2chacha1_xor_block(got, in, uint64(n), key, nonce, ctr)
		if hex.EncodeToString(got) != strings.TrimSpace(hexout) {
			t.Fatalf("n=%d : Go %x, oracle C %s", n, got, hexout)
		}
	}
}

func TestZeroAlloc(t *testing.T) {
	key := make([]byte, 32)
	nonce := make([]byte, 12)
	in := make([]byte, 64)
	out := make([]byte, 64)
	allocs := testing.AllocsPerRun(200, func() {
		C2chacha1_xor_block(out, in, 64, key, nonce, 1)
	})
	if allocs != 0 {
		t.Fatalf("allocations = %.2f, attendu 0", allocs)
	}
}

func BenchmarkChacha1Block(b *testing.B) {
	key := make([]byte, 32)
	nonce := make([]byte, 12)
	in := make([]byte, 64)
	out := make([]byte, 64)
	rand.Read(key)
	rand.Read(nonce)
	rand.Read(in)
	b.Run("emis_1bloc/64", func(b *testing.B) {
		b.SetBytes(64)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			C2chacha1_xor_block(out, in, 64, key, nonce, 1)
		}
	})
	b.Run("xcrypto_chacha20/64", func(b *testing.B) {
		b.SetBytes(64)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			c, _ := chacha20.NewUnauthenticatedCipher(key, nonce)
			c.SetCounter(1)
			c.XORKeyStream(out, in)
		}
	})
}
