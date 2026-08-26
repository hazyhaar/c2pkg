// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2swizzle

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func hashBytes(data []byte) uint64 {
	var h uint64 = 0xcbf29ce484222325
	for _, b := range data {
		h ^= uint64(b)
		h *= 0x100000001b3
	}
	return h
}

func TestSwizzle_UnitAndInvolution(t *testing.T) {
	pxRGBA := uint32(0xAA332211) // A=0xAA, B=0x33, G=0x22, R=0x11
	pxBGRA := C2_swizzle_pixel(pxRGBA)
	pxExpected := uint32(0xAA112233) // A=0xAA, R=0x11, G=0x22, B=0x33

	if pxBGRA != pxExpected {
		t.Fatalf("C2_swizzle_pixel 0x%08X != 0x%08X (obtenu 0x%08X)", pxRGBA, pxExpected, pxBGRA)
	}
	if C2_swizzle_pixel(pxBGRA) != pxRGBA {
		t.Fatalf("C2_swizzle_pixel involution non respectée")
	}

	pairRGBA := uint64(0xFF44332200000000) | uint64(0xAA332211)
	pairBGRA := C2_swizzle_pair(pairRGBA)
	pairExpected := uint64(0xFF22334400000000) | uint64(0xAA112233)

	if pairBGRA != pairExpected {
		t.Fatalf("C2_swizzle_pair 0x%016X != 0x%016X", pairBGRA, pairExpected)
	}
}

func TestSwizzle_KAT(t *testing.T) {
	const katPixels = 100000
	const katBytes = katPixels * 4
	katBuf := make([]byte, katBytes)
	katOut := make([]byte, katBytes)

	for i := range katBuf {
		katBuf[i] = byte((i*131 + 17) & 0xFF)
	}

	C2_swizzle_rgba_bgra(katBuf, katOut, katPixels)
	h := hashBytes(katOut)

	const expectedHash uint64 = 0xB3FE429A3D4FA625
	if h != expectedHash {
		t.Fatalf("KAT RGBA->BGRA Hash = 0x%016X, attendu 0x%016X (oracle C gcc -O2)", h, expectedHash)
	}
}

func TestSwizzle_BoundarySizes(t *testing.T) {
	sizes := []int{
		0, 1, 2, 3, 4, 5, 7, 8, 9, 15, 16, 17, 31, 32, 33, 63, 64, 65, 127, 128, 255, 256, 1024, 16384,
	}

	for _, n := range sizes {
		byteLen := n * 4
		if n == 0 {
			continue
		}

		src := make([]byte, byteLen)
		dst := make([]byte, byteLen)
		dstInplace := make([]byte, byteLen)

		for i := range src {
			src[i] = byte((i*37 + 101) & 0xFF)
		}

		C2_swizzle_rgba_bgra(src, dst, uint64(n))

		// Vérification pixel par pixel
		for i := 0; i < n; i++ {
			off := i * 4
			if dst[off+0] != src[off+2] ||
				dst[off+1] != src[off+1] ||
				dst[off+2] != src[off+0] ||
				dst[off+3] != src[off+3] {
				t.Fatalf("taille %d : discordance pixel %d", n, i)
			}
		}

		// In-place
		copy(dstInplace, src)
		C2_swizzle_rgba_bgra_inplace(dstInplace, uint64(n))
		if !bytes.Equal(dst, dstInplace) {
			t.Fatalf("taille %d : divergence in-place vs out-of-place", n)
		}

		// Double application in-place
		C2_swizzle_rgba_bgra_inplace(dstInplace, uint64(n))
		if !bytes.Equal(src, dstInplace) {
			t.Fatalf("taille %d : involution in-place échouée", n)
		}
	}
}

func TestSwizzle_ARGBConversions(t *testing.T) {
	const n = 1000
	srcARGB := make([]byte, n*4)
	dstRGBA := make([]byte, n*4)
	dstBackARGB := make([]byte, n*4)
	dstBGRA := make([]byte, n*4)

	for i := range srcARGB {
		srcARGB[i] = byte(i & 0xFF)
	}

	C2_swizzle_argb_to_rgba(srcARGB, dstRGBA, n)
	C2_swizzle_rgba_to_argb(dstRGBA, dstBackARGB, n)

	if !bytes.Equal(srcARGB, dstBackARGB) {
		t.Fatalf("roundtrip ARGB <-> RGBA invalide")
	}

	C2_swizzle_argb_to_bgra(srcARGB, dstBGRA, n)
	for i := 0; i < n; i++ {
		off := i * 4
		if dstBGRA[off+0] != srcARGB[off+3] ||
			dstBGRA[off+1] != srcARGB[off+2] ||
			dstBGRA[off+2] != srcARGB[off+1] ||
			dstBGRA[off+3] != srcARGB[off+0] {
			t.Fatalf("pixel %d conversion ARGB -> BGRA invalide", i)
		}
	}
}

func TestSwizzle_ZeroAlloc(t *testing.T) {
	src := make([]byte, 1024*4)
	dst := make([]byte, 1024*4)

	allocs := testing.AllocsPerRun(1000, func() {
		C2_swizzle_rgba_bgra(src, dst, 1024)
	})
	if allocs != 0 {
		t.Fatalf("C2_swizzle_rgba_bgra allocations = %.2f, attendu 0.0", allocs)
	}

	allocsInplace := testing.AllocsPerRun(1000, func() {
		C2_swizzle_rgba_bgra_inplace(src, 1024)
	})
	if allocsInplace != 0 {
		t.Fatalf("C2_swizzle_rgba_bgra_inplace allocations = %.2f, attendu 0.0", allocsInplace)
	}
}

func TestSwizzleVsCOracle(t *testing.T) {
	tmpBin := filepath.Join(t.TempDir(), "swizzle_c_oracle")
	srcDir := filepath.Join("..", "..", "sources")
	if _, err := os.Stat(filepath.Join(srcDir, "c2_swizzle_simd.c")); err != nil {
		srcDir = filepath.Join("sources")
	}

	cmdBuild := exec.Command("gcc", "-O2", "-mavx2", "-Wall", "-Wextra", "-std=gnu99",
		"-I", srcDir,
		filepath.Join(srcDir, "test_swizzle_oracle.c"),
		filepath.Join(srcDir, "c2_swizzle_simd.c"),
		filepath.Join(srcDir, "c2_swizzle_simd_avx2.c"),
		"-o", tmpBin,
	)
	if out, err := cmdBuild.CombinedOutput(); err != nil {
		t.Fatalf("échec de compilation oracle C gcc -O2: %v\nSortie: %s", err, string(out))
	}

	cmdRun := exec.Command(tmpBin)
	out, err := cmdRun.CombinedOutput()
	if err != nil {
		t.Fatalf("échec d'exécution oracle C: %v\nSortie: %s", err, string(out))
	}

	if !bytes.Contains(out, []byte("0xB3FE429A3D4FA625")) {
		t.Fatalf("l'oracle C n'a pas produit le hash KAT attendu : %s", string(out))
	}
	t.Logf("Sortie Oracle C (gcc -O2 -mavx2):\n%s", string(out))
}

func BenchmarkSwizzle_RGBA_BGRA_1080p(b *testing.B) {
	const numPixels = 1920 * 1080 // Trame Full HD 1080p (~8.3 MB)
	src := make([]byte, numPixels*4)
	dst := make([]byte, numPixels*4)

	b.SetBytes(numPixels * 4)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		C2_swizzle_rgba_bgra(src, dst, numPixels)
	}
}

func BenchmarkSwizzle_RGBA_BGRA_Inplace_1080p(b *testing.B) {
	const numPixels = 1920 * 1080
	buf := make([]byte, numPixels*4)

	b.SetBytes(numPixels * 4)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		C2_swizzle_rgba_bgra_inplace(buf, numPixels)
	}
}
