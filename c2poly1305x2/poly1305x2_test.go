// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2poly1305x2

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

	"github.com/hazyhaar/c2pkg/c2poly1305"
	"golang.org/x/crypto/poly1305"
)

func TestRFC8439Vector(t *testing.T) {
	key, _ := hex.DecodeString("85d6be7857556d337f4452fe42d506a80103808afb0db2fd4abff6af4149f51b")
	msg := []byte("Cryptographic Forum Research Group")
	want, _ := hex.DecodeString("a8061dc1305136c6c22b8baf0c0127a9")
	var mac [16]byte
	Crypto_poly1305x2(mac[:], msg, uint64(len(msg)), key)
	if !bytes.Equal(mac[:], want) {
		t.Fatalf("RFC 8439 : obtenu %x, attendu %x", mac, want)
	}
}

func TestParityXCrypto(t *testing.T) {
	for n := 0; n <= 1088; n++ {
		var key [32]byte
		msg := make([]byte, n)
		rand.Read(key[:])
		rand.Read(msg)
		var want [16]byte
		poly1305.Sum(&want, msg, &key)

		var got [16]byte
		Crypto_poly1305x2(got[:], msg, uint64(n), key[:])
		if got != want {
			t.Fatalf("n=%d une passe : %x != %x", n, got, want)
		}

		var ctx Crypto_poly1305x2_ctx
		Crypto_poly1305x2_init(&ctx, key[:])
		for off := 0; off < n; {
			step := 1 + (off*7+n)%29
			if off+step > n {
				step = n - off
			}
			Crypto_poly1305x2_update(&ctx, msg[off:off+step], uint64(step))
			off += step
		}
		Crypto_poly1305x2_final(&ctx, got[:])
		if got != want {
			t.Fatalf("n=%d fragmenté : %x != %x", n, got, want)
		}
	}
}

func TestParityC2Poly1305(t *testing.T) {
	for i := 0; i < 4000; i++ {
		var nBuf [2]byte
		rand.Read(nBuf[:])
		n := int(nBuf[0]) | int(nBuf[1])<<8
		n %= 4097
		var key [32]byte
		msg := make([]byte, n)
		rand.Read(key[:])
		rand.Read(msg)
		var want, got [16]byte
		c2poly1305.Crypto_poly1305(want[:], msg, uint64(n), key[:])
		Crypto_poly1305x2(got[:], msg, uint64(n), key[:])
		if got != want {
			t.Fatalf("tirage %d n=%d : x2 %x != mono %x", i, n, got, want)
		}
	}
}

type lcg struct{ s uint64 }

func (l *lcg) seed(n uint64) { l.s = n*0x9E3779B97F4A7C15 + 1 }
func (l *lcg) next() byte {
	l.s = l.s*6364136223846793005 + 1442695040888963407
	return byte(l.s >> 56)
}

func oracleSrcDir(t *testing.T) string {
	t.Helper()
	srcDir := filepath.Join("..", "..", "sources", "c2poly1305x2")
	for _, name := range []string{"c2poly1305x2.c", "c2poly1305x2.h", "test_c2poly1305x2_oracle.c"} {
		if _, err := os.Stat(filepath.Join(srcDir, name)); err != nil {
			t.Skipf("source oracle absente : %s", name)
		}
	}
	return srcDir
}

func compileOracle(t *testing.T, extra ...string) string {
	t.Helper()
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("gcc absent : oracle C non jouable")
	}
	srcDir := oracleSrcDir(t)
	tmpBin := filepath.Join(t.TempDir(), "c2poly1305x2_oracle")
	args := append([]string{}, extra...)
	args = append(args, "-Wall", "-Wextra", "-std=gnu99", "-I", srcDir,
		filepath.Join(srcDir, "test_c2poly1305x2_oracle.c"),
		filepath.Join(srcDir, "c2poly1305x2.c"),
		"-o", tmpBin)
	build := exec.Command("gcc", args...)
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("compilation oracle C : %v\n%s", err, out)
	}
	return tmpBin
}

func checkOracleOutput(t *testing.T, out []byte) {
	t.Helper()
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) != 25 {
		t.Fatalf("oracle C : %d lignes, attendu 25", len(lines))
	}
	if !strings.HasPrefix(lines[0], "rfc:") {
		t.Fatalf("première ligne oracle sans préfixe rfc : %q", lines[0])
	}
	rfcHex := strings.TrimPrefix(lines[0], "rfc:")
	var rfcMac [16]byte
	key, _ := hex.DecodeString("85d6be7857556d337f4452fe42d506a80103808afb0db2fd4abff6af4149f51b")
	msg := []byte("Cryptographic Forum Research Group")
	Crypto_poly1305x2(rfcMac[:], msg, uint64(len(msg)), key)
	if hex.EncodeToString(rfcMac[:]) != rfcHex {
		t.Fatalf("RFC : Go %x, oracle C %s", rfcMac, rfcHex)
	}
	var g lcg
	for _, line := range lines[1:] {
		var n int
		var hexmac string
		if _, err := fmt.Sscanf(line, "%d:%s", &n, &hexmac); err != nil {
			t.Fatalf("ligne oracle illisible %q : %v", line, err)
		}
		g.seed(uint64(n))
		k := make([]byte, 32)
		for i := range k {
			k[i] = g.next()
		}
		m := make([]byte, n)
		for i := range m {
			m[i] = g.next()
		}
		var mac [16]byte
		Crypto_poly1305x2(mac[:], m, uint64(n), k)
		if hex.EncodeToString(mac[:]) != hexmac {
			t.Fatalf("n=%d : Go %x, oracle C %s", n, mac, hexmac)
		}
	}
}

func TestPoly1305x2VsCOracle(t *testing.T) {
	tmpBin := compileOracle(t, "-O2")
	out, err := exec.Command(tmpBin).Output()
	if err != nil {
		t.Fatalf("exécution oracle C : %v", err)
	}
	checkOracleOutput(t, out)
}

func TestOracleASanUBSan(t *testing.T) {
	tmpBin := compileOracle(t, "-O2", "-fsanitize=address,undefined")
	cmd := exec.Command(tmpBin)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("oracle ASan/UBSan : %v\n%s", err, out)
	}
	if bytes.Contains(out, []byte("ERROR:")) || bytes.Contains(out, []byte("runtime error:")) {
		t.Fatalf("sanitizer non muet :\n%s", out)
	}
	checkOracleOutput(t, out)
}

func TestZeroAlloc(t *testing.T) {
	key := make([]byte, 32)
	msg := make([]byte, 64)
	var mac [16]byte
	allocs := testing.AllocsPerRun(200, func() {
		var ctx Crypto_poly1305x2_ctx
		Crypto_poly1305x2_init(&ctx, key)
		Crypto_poly1305x2_update(&ctx, msg, 64)
		Crypto_poly1305x2_final(&ctx, mac[:])
	})
	if allocs != 0 {
		t.Fatalf("allocations = %.2f, attendu 0", allocs)
	}
}

func BenchmarkPoly1305x2(b *testing.B) {
	for _, n := range []int{64, 1350, 8192} {
		key := make([]byte, 32)
		msg := make([]byte, n)
		rand.Read(key)
		rand.Read(msg)
		var mac [16]byte
		b.Run(fmt.Sprintf("emis_mono/%d", n), func(b *testing.B) {
			b.SetBytes(int64(n))
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				var ctx c2poly1305.Crypto_poly1305_ctx
				c2poly1305.Crypto_poly1305_init(&ctx, key)
				c2poly1305.Crypto_poly1305_update(&ctx, msg, uint64(n))
				c2poly1305.Crypto_poly1305_final(&ctx, mac[:])
			}
		})
		b.Run(fmt.Sprintf("emis_x2/%d", n), func(b *testing.B) {
			b.SetBytes(int64(n))
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				var ctx Crypto_poly1305x2_ctx
				Crypto_poly1305x2_init(&ctx, key)
				Crypto_poly1305x2_update(&ctx, msg, uint64(n))
				Crypto_poly1305x2_final(&ctx, mac[:])
			}
		})
		b.Run(fmt.Sprintf("xcrypto_asm/%d", n), func(b *testing.B) {
			var k [32]byte
			copy(k[:], key)
			b.SetBytes(int64(n))
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				poly1305.Sum(&mac, msg, &k)
			}
		})
	}
}
