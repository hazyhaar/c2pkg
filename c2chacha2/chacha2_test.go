// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2chacha2

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

	"github.com/hazyhaar/c2pkg/c2chacha1"
	"golang.org/x/crypto/chacha20"
)

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
	if got := C2chacha2_xor_blocks(out, msg, 114, key, nonce, 1); got != 114 {
		t.Fatalf("n rendu %d", got)
	}
	if !bytes.Equal(out, want) {
		t.Fatalf("RFC 8439 : obtenu %x", out)
	}
}

func TestParityXCrypto(t *testing.T) {
	for n := 0; n <= 128; n++ {
		for rep := 0; rep < 8; rep++ {
			key := make([]byte, 32)
			nonce := make([]byte, 12)
			msg := make([]byte, n)
			rand.Read(key)
			rand.Read(nonce)
			rand.Read(msg)
			var ctrB [4]byte
			rand.Read(ctrB[:])
			ctr := uint32(ctrB[0]) | uint32(ctrB[1])<<8 | uint32(ctrB[2])<<16 | uint32(ctrB[3])<<24
			if ctr > 0xFFFFFFFE {
				ctr -= 2
			}

			c, err := chacha20.NewUnauthenticatedCipher(key, nonce)
			if err != nil {
				t.Fatal(err)
			}
			c.SetCounter(ctr)
			want := make([]byte, n)
			c.XORKeyStream(want, msg)

			out := make([]byte, n)
			C2chacha2_xor_blocks(out, msg, uint64(n), key, nonce, ctr)
			if !bytes.Equal(out, want) {
				t.Fatalf("n=%d ctr=%d : %x != %x", n, ctr, out, want)
			}
			inPlace := append([]byte(nil), msg...)
			C2chacha2_xor_blocks(inPlace, inPlace, uint64(n), key, nonce, ctr)
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

func TestChacha2VsCOracle(t *testing.T) {
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("gcc absent : oracle C non jouable")
	}
	srcDir := filepath.Join("..", "..", "sources", "c2chacha2")
	if _, err := os.Stat(filepath.Join(srcDir, "c2chacha2_simd.c")); err != nil {
		t.Skipf("source oracle absente : %v", err)
	}
	tmpBin := filepath.Join(t.TempDir(), "c2chacha2_oracle")
	build := exec.Command("gcc", "-O2", "-mavx2", "-Wall", "-Wextra", "-std=gnu99", "-I", srcDir,
		filepath.Join(srcDir, "test_c2chacha2_oracle.c"), filepath.Join(srcDir, "c2chacha2_simd.c"), "-o", tmpBin)
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("compilation oracle C : %v\n%s", err, out)
	}
	out, err := exec.Command(tmpBin).Output()
	if err != nil {
		t.Fatalf("exécution oracle C : %v", err)
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) != 17 {
		t.Fatalf("oracle C : %d lignes, attendu 17 (rfc + 14 tailles + 2 débordements)", len(lines))
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
	C2chacha2_xor_blocks(gotRFC, msgRFC, 114, keyRFC, nonceRFC, 1)
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
			in := make([]byte, 128)
			for k := range in {
				in[k] = g.next()
			}
			got := make([]byte, 128)
			C2chacha2_xor_blocks(got, in, 128, key, nonce, ctr)
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
		C2chacha2_xor_blocks(got, in, uint64(n), key, nonce, ctr)
		if hex.EncodeToString(got) != strings.TrimSpace(hexout) {
			t.Fatalf("n=%d : Go %x, oracle C %s", n, got, hexout)
		}
	}
	if wraps != 2 {
		t.Fatalf("débordements vérifiés contre le C : %d, attendu 2", wraps)
	}
}

func TestZeroAlloc(t *testing.T) {
	key := make([]byte, 32)
	nonce := make([]byte, 12)
	in := make([]byte, 128)
	out := make([]byte, 128)
	allocs := testing.AllocsPerRun(200, func() {
		C2chacha2_xor_blocks(out, in, 128, key, nonce, 1)
	})
	if allocs != 0 {
		t.Fatalf("allocations = %.2f, attendu 0", allocs)
	}
}

func BenchmarkChacha2Blocks(b *testing.B) {
	key := make([]byte, 32)
	nonce := make([]byte, 12)
	in := make([]byte, 128)
	out := make([]byte, 128)
	rand.Read(key)
	rand.Read(nonce)
	rand.Read(in)
	b.Run("emis_2blocs/128", func(b *testing.B) {
		b.SetBytes(128)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			C2chacha2_xor_blocks(out, in, 128, key, nonce, 1)
		}
	})
	zero64 := make([]byte, 64)
	plain := make([]byte, 64)
	copy(plain, in[:64])
	poly := make([]byte, 64)
	ks := make([]byte, 64)
	fusedIn := make([]byte, 128)
	fusedOut := make([]byte, 128)
	copy(fusedIn[64:], plain)
	b.Run("c2chacha1x2/poly+64", func(b *testing.B) {
		b.SetBytes(128)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			c2chacha1.C2chacha1_xor_block(poly, zero64, 64, key, nonce, 0)
			c2chacha1.C2chacha1_xor_block(ks, plain, 64, key, nonce, 1)
		}
	})
	b.Run("emis_2blocs/fused128", func(b *testing.B) {
		b.SetBytes(128)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			C2chacha2_xor_blocks(fusedOut, fusedIn, 128, key, nonce, 0)
		}
	})
	rawKS := make([]byte, 128)
	zero128 := make([]byte, 128)
	b.Run("emis_2blocs/ks128+xor", func(b *testing.B) {
		b.SetBytes(128)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			C2chacha2_xor_blocks(rawKS, zero128, 128, key, nonce, 0)
			for k := 0; k < 64; k++ {
				ks[k] = plain[k] ^ rawKS[64+k]
			}
		}
	})
	b.Run("xcrypto_chacha20/128", func(b *testing.B) {
		b.SetBytes(128)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			c, _ := chacha20.NewUnauthenticatedCipher(key, nonce)
			c.SetCounter(1)
			c.XORKeyStream(out, in)
		}
	})
}
