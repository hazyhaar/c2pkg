// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2chacha8

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

// Vecteur RFC 8439 §2.4.2, 114 octets (huit blocs couvrent le texte entier).
func TestRFC8439Blocks(t *testing.T) {
	key := make([]byte, 32)
	for i := range key {
		key[i] = byte(i)
	}
	nonce := make([]byte, 12)
	nonce[7] = 0x4a
	msg := []byte("Ladies and Gentlemen of the class of '99: If I could offer you only one tip for the future, sunscreen would be it.")
	want, _ := hex.DecodeString("6e2e359a2568f98041ba0728dd0d6981e97e7aec1d4360c20a27afccfd9fae0bf91b65c5524733ab8f593dabcd62b3571639d624e65152ab8f530c359f0861d807ca0dbf500d6a6156a38e088a22b65e52bc514d16ccf806818ce91ab77937365af90bbf74a35be6b40b8eedf2785e42874d")
	out := make([]byte, 114)
	if got := C2chacha8_xor_blocks(out, msg, 114, key, nonce, 1); got != 114 {
		t.Fatalf("n rendu %d", got)
	}
	if !bytes.Equal(out, want) {
		t.Fatalf("RFC 8439 : obtenu %x", out)
	}
}

// Parité contre golang.org/x/crypto/chacha20 pour toutes les longueurs 0..512,
// clé, nonce, compteur et message aléatoires ; en place et hors place.
// x/crypto refuse (panique) un compteur dont l'avance dépasse 2^32 : les
// compteurs sont donc bornés ici à [0, 2^32−9] ; le débordement par voie
// (counter+[0…7] modulo 2^32) est prouvé contre l'oracle C par
// TestCounterWrapVsCOracle, x/crypto ne pouvant pas en être l'oracle.
func TestParityXCrypto(t *testing.T) {
	for n := 0; n <= 512; n++ {
		for rep := 0; rep < 4; rep++ {
			key := make([]byte, 32)
			nonce := make([]byte, 12)
			msg := make([]byte, n)
			rand.Read(key)
			rand.Read(nonce)
			rand.Read(msg)
			var ctrB [4]byte
			rand.Read(ctrB[:])
			ctr := uint32(ctrB[0]) | uint32(ctrB[1])<<8 | uint32(ctrB[2])<<16 | uint32(ctrB[3])<<24
			if ctr > 0xFFFFFFF7 {
				ctr -= 8
			}

			c, err := chacha20.NewUnauthenticatedCipher(key, nonce)
			if err != nil {
				t.Fatal(err)
			}
			c.SetCounter(ctr)
			want := make([]byte, n)
			c.XORKeyStream(want, msg)

			out := make([]byte, n)
			C2chacha8_xor_blocks(out, msg, uint64(n), key, nonce, ctr)
			if !bytes.Equal(out, want) {
				t.Fatalf("n=%d ctr=%d : %x != %x", n, ctr, out, want)
			}
			inPlace := append([]byte(nil), msg...)
			C2chacha8_xor_blocks(inPlace, inPlace, uint64(n), key, nonce, ctr)
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

// Oracle gcc -O2 -mavx2 : la même source C que celle transpilée, compilée nativement.
func TestChacha8VsCOracle(t *testing.T) {
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("gcc absent : oracle C non jouable")
	}
	srcDir := filepath.Join("..", "..", "sources", "c2chacha8")
	if _, err := os.Stat(filepath.Join(srcDir, "c2chacha8_simd.c")); err != nil {
		t.Skipf("source oracle absente : %v", err)
	}
	tmpBin := filepath.Join(t.TempDir(), "c2chacha8_oracle")
	build := exec.Command("gcc", "-O2", "-mavx2", "-Wall", "-Wextra", "-std=gnu99", "-I", srcDir,
		filepath.Join(srcDir, "test_c2chacha8_oracle.c"), filepath.Join(srcDir, "c2chacha8_simd.c"), "-o", tmpBin)
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("compilation oracle C : %v\n%s", err, out)
	}
	out, err := exec.Command(tmpBin).Output()
	if err != nil {
		t.Fatalf("exécution oracle C : %v", err)
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) != 22 {
		t.Fatalf("oracle C : %d lignes, attendu 22 (rfc + 18 tailles + 3 débordements)", len(lines))
	}
	rfcWant := "6e2e359a2568f98041ba0728dd0d6981e97e7aec1d4360c20a27afccfd9fae0bf91b65c5524733ab8f593dabcd62b3571639d624e65152ab8f530c359f0861d807ca0dbf500d6a6156a38e088a22b65e52bc514d16ccf806818ce91ab77937365af90bbf74a35be6b40b8eedf2785e42874d"
	if !strings.HasPrefix(lines[0], "rfc:") {
		t.Fatalf("première ligne oracle : %q", lines[0])
	}
	keyRFC := make([]byte, 32)
	for i := range keyRFC {
		keyRFC[i] = byte(i)
	}
	nonceRFC := make([]byte, 12)
	nonceRFC[7] = 0x4a
	msgRFC := []byte("Ladies and Gentlemen of the class of '99: If I could offer you only one tip for the future, sunscreen would be it.")
	gotRFC := make([]byte, 114)
	C2chacha8_xor_blocks(gotRFC, msgRFC, 114, keyRFC, nonceRFC, 1)
	if hex.EncodeToString(gotRFC) != strings.TrimPrefix(lines[0], "rfc:") {
		t.Fatalf("rfc : Go %x, oracle C %s", gotRFC, strings.TrimPrefix(lines[0], "rfc:"))
	}
	if hex.EncodeToString(gotRFC) != rfcWant {
		t.Fatalf("rfc : Go %x ≠ RFC", gotRFC)
	}
	var g lcg
	wraps := 0
	for _, line := range lines[1:] {
		if strings.HasPrefix(line, "wrap:") {
			// Débordement du compteur par voie : counter + [0..7] modulo 2^32,
			// prouvé contre le C seul (x/crypto refuse ces compteurs).
			var ctr uint32
			var hexout string
			if _, err := fmt.Sscanf(line, "wrap:%08x:%s", &ctr, &hexout); err != nil {
				t.Fatalf("ligne wrap illisible %q : %v", line, err)
			}
			g.seed(uint64(ctr))
			key := make([]byte, 32)
			for k := range key {
				key[k] = g.next()
			}
			nonce := make([]byte, 12)
			for k := range nonce {
				nonce[k] = g.next()
			}
			in := make([]byte, 512)
			for k := range in {
				in[k] = g.next()
			}
			got := make([]byte, 512)
			C2chacha8_xor_blocks(got, in, 512, key, nonce, ctr)
			if hex.EncodeToString(got) != strings.TrimSpace(hexout) {
				t.Fatalf("wrap ctr=%08x : Go diverge de l'oracle C", ctr)
			}
			wraps++
			continue
		}
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
		C2chacha8_xor_blocks(got, in, uint64(n), key, nonce, ctr)
		if hex.EncodeToString(got) != strings.TrimSpace(hexout) {
			t.Fatalf("n=%d : Go %x, oracle C %s", n, got, hexout)
		}
	}
	if wraps != 3 {
		t.Fatalf("débordements vérifiés contre le C : %d, attendu 3", wraps)
	}
}

func TestZeroAlloc(t *testing.T) {
	key := make([]byte, 32)
	nonce := make([]byte, 12)
	in := make([]byte, 512)
	out := make([]byte, 512)
	allocs := testing.AllocsPerRun(200, func() {
		C2chacha8_xor_blocks(out, in, 512, key, nonce, 1)
	})
	if allocs != 0 {
		t.Fatalf("allocations = %.2f, attendu 0", allocs)
	}
}

func BenchmarkChacha8Blocks(b *testing.B) {
	key := make([]byte, 32)
	nonce := make([]byte, 12)
	in := make([]byte, 512)
	out := make([]byte, 512)
	rand.Read(key)
	rand.Read(nonce)
	rand.Read(in)
	b.Run("emis_8blocs", func(b *testing.B) {
		b.SetBytes(512)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			C2chacha8_xor_blocks(out, in, 512, key, nonce, 1)
		}
	})
	b.Run("xcrypto_chacha20", func(b *testing.B) {
		b.SetBytes(512)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			c, _ := chacha20.NewUnauthenticatedCipher(key, nonce)
			c.SetCounter(1)
			c.XORKeyStream(out, in)
		}
	})
}
