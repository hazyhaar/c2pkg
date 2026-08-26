// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2fused

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hazyhaar/c2pkg/c2chacha8"
	"github.com/hazyhaar/c2pkg/c2poly1305"
	"golang.org/x/crypto/chacha20poly1305"
)

func TestRFC8439Poly(t *testing.T) {
	key, _ := hex.DecodeString("85d6be7857556d337f4452fe42d506a80103808afb0db2fd4abff6af4149f51b")
	msg := []byte("Cryptographic Forum Research Group")
	want, _ := hex.DecodeString("a8061dc1305136c6c22b8baf0c0127a9")
	var st C2fused_poly
	C2fused_poly_init(&st, key)
	C2fused_poly_update(&st, msg, uint64(len(msg)))
	var mac [16]byte
	C2fused_poly_final(&st, mac[:])
	if !bytes.Equal(mac[:], want) {
		t.Fatalf("RFC 8439 §2.5.2 : obtenu %x, attendu %x", mac, want)
	}
}

func rfc8439AEAD() (key, nonce, ad, msg, tag []byte) {
	key = make([]byte, 32)
	for i := range key {
		key[i] = byte(0x80 + i)
	}
	nonce = make([]byte, 12)
	copy(nonce, []byte{0x07, 0, 0, 0, 0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47})
	ad = []byte{0x50, 0x51, 0x52, 0x53, 0xc0, 0xc1, 0xc2, 0xc3, 0xc4, 0xc5, 0xc6, 0xc7}
	msg = []byte("Ladies and Gentlemen of the class of '99: If I could offer you only one tip for the future, sunscreen would be it.")
	tag, _ = hex.DecodeString("1ae10b594f09e26a7e902ecbd0600691")
	return
}

func polyPad(st *C2fused_poly, n int) {
	if rem := n & 15; rem != 0 {
		var z [16]byte
		C2fused_poly_update(st, z[:], uint64(16-rem))
	}
}

func aeadFused(out, tag, in, ad, key, nonce []byte) {
	var zeros, polykey [64]byte
	c2chacha8.C2chacha8_xor_blocks(polykey[:], zeros[:], 64, key, nonce, 0)
	var st C2fused_poly
	C2fused_poly_init(&st, polykey[:])
	C2fused_poly_update(&st, ad, uint64(len(ad)))
	polyPad(&st, len(ad))
	ctr := uint32(1)
	var prev []byte
	n := len(in)
	for off := 0; off < n; off += 512 {
		chunk := n - off
		if chunk > 512 {
			chunk = 512
		}
		if prev != nil && chunk == 512 {
			C2fused_seal_blocks(out[off:], in[off:], uint64(chunk), key, nonce, ctr, &st, prev)
		} else {
			if prev != nil {
				C2fused_poly_blocks(&st, prev, 32)
			}
			c2chacha8.C2chacha8_xor_blocks(out[off:], in[off:], uint64(chunk), key, nonce, ctr)
		}
		prev = out[off : off+chunk]
		if chunk != 512 {
			C2fused_poly_update(&st, prev, uint64(chunk))
			prev = nil
		}
		ctr += 8
	}
	if prev != nil {
		C2fused_poly_blocks(&st, prev, 32)
	}
	polyPad(&st, n)
	var lens [16]byte
	ln := uint64(len(ad))
	for i := 0; i < 8; i++ {
		lens[i] = byte(ln >> (8 * i))
		lens[8+i] = byte(uint64(n) >> (8 * i))
	}
	C2fused_poly_update(&st, lens[:], 16)
	C2fused_poly_final(&st, tag)
}

func aeadSeq(out, tag, in, ad, key, nonce []byte) {
	var zeros, polykey [64]byte
	c2chacha8.C2chacha8_xor_blocks(polykey[:], zeros[:], 64, key, nonce, 0)
	var st C2fused_poly
	C2fused_poly_init(&st, polykey[:])
	C2fused_poly_update(&st, ad, uint64(len(ad)))
	polyPad(&st, len(ad))
	ctr := uint32(1)
	n := len(in)
	for off := 0; off < n; off += 512 {
		chunk := n - off
		if chunk > 512 {
			chunk = 512
		}
		c2chacha8.C2chacha8_xor_blocks(out[off:], in[off:], uint64(chunk), key, nonce, ctr)
		ctr += 8
		C2fused_poly_update(&st, out[off:off+chunk], uint64(chunk))
	}
	polyPad(&st, n)
	var lens [16]byte
	ln := uint64(len(ad))
	for i := 0; i < 8; i++ {
		lens[i] = byte(ln >> (8 * i))
		lens[8+i] = byte(uint64(n) >> (8 * i))
	}
	C2fused_poly_update(&st, lens[:], 16)
	C2fused_poly_final(&st, tag)
}

func TestRFC8439AEAD(t *testing.T) {
	key, nonce, ad, msg, wantTag := rfc8439AEAD()
	out := make([]byte, len(msg))
	tag := make([]byte, 16)
	aeadFused(out, tag, msg, ad, key, nonce)
	if !bytes.Equal(tag, wantTag) {
		t.Fatalf("RFC 8439 §2.8.2 tag : obtenu %x, attendu %x", tag, wantTag)
	}
	outS := make([]byte, len(msg))
	tagS := make([]byte, 16)
	aeadSeq(outS, tagS, msg, ad, key, nonce)
	if !bytes.Equal(out, outS) || !bytes.Equal(tag, tagS) {
		t.Fatalf("RFC 8439 §2.8.2 fusionné ≠ séquentiel")
	}
}

func TestParityXCrypto(t *testing.T) {
	for k := 0; k < 512; k++ {
		n := (k % 16) * 512
		if n == 0 {
			n = 64
		}
		key := make([]byte, 32)
		nonce := make([]byte, 12)
		ad := make([]byte, k%17)
		msg := make([]byte, n)
		rand.Read(key)
		rand.Read(nonce)
		rand.Read(ad)
		rand.Read(msg)
		aead, err := chacha20poly1305.New(key)
		if err != nil {
			t.Fatal(err)
		}
		want := aead.Seal(nil, nonce, msg, ad)
		out := make([]byte, n)
		tag := make([]byte, 16)
		aeadFused(out, tag, msg, ad, key, nonce)
		got := append(append([]byte(nil), out...), tag...)
		if !bytes.Equal(got, want) {
			t.Fatalf("k=%d n=%d : diverge de x/crypto", k, n)
		}
	}
}

func TestParity4000(t *testing.T) {
	type lcg struct{ s uint64 }
	next := func(l *lcg) byte {
		l.s = l.s*6364136223846793005 + 1442695040888963407
		return byte(l.s >> 56)
	}
	for draw := 0; draw < 4000; draw++ {
		var l lcg
		l.s = uint64(draw+1)*0x9E3779B97F4A7C15 + 1
		n := (1 + int(next(&l))%16) * 512
		if n > 8192 {
			n = 8192
		}
		adn := int(next(&l) % 14)
		key := make([]byte, 32)
		nonce := make([]byte, 12)
		ad := make([]byte, adn)
		in := make([]byte, n)
		for i := range key {
			key[i] = next(&l)
		}
		for i := range nonce {
			nonce[i] = next(&l)
		}
		for i := range ad {
			ad[i] = next(&l)
		}
		for i := range in {
			in[i] = next(&l)
		}
		outF := make([]byte, n)
		outS := make([]byte, n)
		tagF := make([]byte, 16)
		tagS := make([]byte, 16)
		aeadFused(outF, tagF, in, ad, key, nonce)
		aeadSeq(outS, tagS, in, ad, key, nonce)
		if !bytes.Equal(outF, outS) || !bytes.Equal(tagF, tagS) {
			t.Fatalf("draw %d n=%d fusionné ≠ séquentiel", draw, n)
		}
		var pst c2poly1305.Crypto_poly1305_ctx
		var zeros, polykey [64]byte
		c2chacha8.C2chacha8_xor_blocks(polykey[:], zeros[:], 64, key, nonce, 0)
		c2poly1305.Crypto_poly1305_init(&pst, polykey[:])
		c2poly1305.Crypto_poly1305_update(&pst, ad, uint64(adn))
		if rem := adn & 15; rem != 0 {
			var z [16]byte
			c2poly1305.Crypto_poly1305_update(&pst, z[:], uint64(16-rem))
		}
		ctr := uint32(1)
		outC := make([]byte, n)
		for off := 0; off < n; off += 512 {
			chunk := n - off
			if chunk > 512 {
				chunk = 512
			}
			c2chacha8.C2chacha8_xor_blocks(outC[off:], in[off:], uint64(chunk), key, nonce, ctr)
			ctr += 8
			c2poly1305.Crypto_poly1305_update(&pst, outC[off:off+chunk], uint64(chunk))
		}
		if rem := n & 15; rem != 0 {
			var z [16]byte
			c2poly1305.Crypto_poly1305_update(&pst, z[:], uint64(16-rem))
		}
		var lens [16]byte
		ln := uint64(adn)
		for i := 0; i < 8; i++ {
			lens[i] = byte(ln >> (8 * i))
			lens[8+i] = byte(uint64(n) >> (8 * i))
		}
		c2poly1305.Crypto_poly1305_update(&pst, lens[:], 16)
		var tagC [16]byte
		c2poly1305.Crypto_poly1305_final(&pst, tagC[:])
		if !bytes.Equal(outF, outC) || !bytes.Equal(tagF, tagC[:]) {
			t.Fatalf("draw %d n=%d ≠ c2chacha8+c2poly1305", draw, n)
		}
	}
}

func TestFusedVsCOracle(t *testing.T) {
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("gcc absent : oracle C non jouable")
	}
	srcDir := filepath.Join("..", "..", "sources", "c2fused")
	chachaDir := filepath.Join("..", "..", "sources", "c2chacha8")
	for _, p := range []string{
		filepath.Join(srcDir, "c2fused_simd.c"),
		filepath.Join(srcDir, "test_c2fused_oracle.c"),
		filepath.Join(chachaDir, "c2chacha8_simd.c"),
	} {
		if _, err := os.Stat(p); err != nil {
			t.Skipf("source oracle absente : %s", p)
		}
	}
	tmpBin := filepath.Join(t.TempDir(), "c2fused_oracle")
	build := exec.Command("gcc", "-O2", "-mavx2", "-Wall", "-Wextra", "-std=gnu99",
		"-I", srcDir, "-I", chachaDir,
		filepath.Join(srcDir, "test_c2fused_oracle.c"),
		filepath.Join(srcDir, "c2fused_simd.c"),
		filepath.Join(chachaDir, "c2chacha8_simd.c"),
		"-o", tmpBin)
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("compilation oracle C : %v\n%s", err, out)
	}
	out, err := exec.Command(tmpBin).Output()
	if err != nil {
		t.Fatalf("exécution oracle C : %v\n%s", err, out)
	}
	if !strings.Contains(string(out), "draws:4000 ok") {
		t.Fatalf("oracle C incomplet :\n%s", out)
	}
	key, nonce, ad, msg, wantTag := rfc8439AEAD()
	got := make([]byte, len(msg))
	tag := make([]byte, 16)
	aeadFused(got, tag, msg, ad, key, nonce)
	if !bytes.Equal(tag, wantTag) {
		t.Fatalf("tag Go %x ≠ RFC / oracle", tag)
	}
}

func TestZeroAlloc(t *testing.T) {
	key := make([]byte, 32)
	nonce := make([]byte, 12)
	in := make([]byte, 512)
	out := make([]byte, 512)
	prev := make([]byte, 512)
	var st C2fused_poly
	C2fused_poly_init(&st, key)
	allocs := testing.AllocsPerRun(200, func() {
		C2fused_seal_blocks(out, in, 512, key, nonce, 1, &st, prev)
	})
	if allocs != 0 {
		t.Fatalf("allocations = %.2f, attendu 0", allocs)
	}
}
