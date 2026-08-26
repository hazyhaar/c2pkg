// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2ssim

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

func TestSSIM_Identity(t *testing.T) {
	const w = 64
	const h = 64
	const sz = w * h

	img1 := make([]byte, sz)
	img2 := make([]byte, sz)

	for i := range img1 {
		img1[i] = byte((i*17 + 43) & 0xFF)
		img2[i] = img1[i]
	}

	scoreMilli := C2_ssim_compute_milli(img1, img2, w, h, w)
	scoreQ16 := C2_ssim_compute_q16(img1, img2, w, h, w)

	if scoreMilli != 1000000 {
		t.Fatalf("score SSIM milli identité = %d, attendu 1000000", scoreMilli)
	}
	if scoreQ16 != 65536 {
		t.Fatalf("score SSIM Q16 identité = %d, attendu 65536", scoreQ16)
	}
}

func TestSSIM_DCGainPreservation(t *testing.T) {
	const w = 64
	const h = 64
	const sz = w * h

	img := make([]byte, sz)
	dst := make([]byte, sz)
	tmp := make([]uint16, sz)

	for i := range img {
		img[i] = 128
	}

	C2_gaussian_blur_2d(img, dst, tmp, w, h, w)

	for i, v := range dst {
		if v != 128 {
			t.Fatalf("pixel %d : valeur = %d, attendu 128 (gain continu dégradé)", i, v)
		}
	}
}

func TestSSIM_KAT(t *testing.T) {
	const kw = 128
	const kh = 128
	const ksz = kw * kh

	kimg1 := make([]byte, ksz)
	kimg2 := make([]byte, ksz)
	kblur := make([]byte, ksz)
	ktmp := make([]uint16, ksz)

	for i := 0; i < ksz; i++ {
		kimg1[i] = byte((i*59 + 13) & 0xFF)
		kimg2[i] = byte((i*67 + 29) & 0xFF)
	}

	C2_gaussian_blur_2d(kimg1, kblur, ktmp, kw, kh, kw)
	h := hashBytes(kblur)

	const expectedBlurHash uint64 = 0x03FE76258E335041
	if h != expectedBlurHash {
		t.Fatalf("KAT GaussBlur 128x128 Hash = 0x%016X, attendu 0x%016X (oracle C gcc -O2)", h, expectedBlurHash)
	}

	score := C2_ssim_compute_milli(kimg1, kimg2, kw, kh, kw)
	const expectedScore int64 = 353076
	if score != expectedScore {
		t.Fatalf("KAT SSIM 128x128 Score = %d, attendu %d (oracle C gcc -O2)", score, expectedScore)
	}
}

func TestSSIM_NoiseMonotonicity(t *testing.T) {
	const w = 64
	const h = 64
	const sz = w * h

	ref := make([]byte, sz)
	for i := range ref {
		ref[i] = byte((i*31 + 7) & 0xFF)
	}

	prevScore := int64(1000000)
	for noiseAmp := 1; noiseAmp <= 10; noiseAmp++ {
		noisy := make([]byte, sz)
		for i := range ref {
			val := int(ref[i]) + (i%3-1)*noiseAmp
			if val < 0 {
				val = 0
			}
			if val > 255 {
				val = 255
			}
			noisy[i] = byte(val)
		}

		score := C2_ssim_compute_milli(ref, noisy, w, h, w)
		if score >= prevScore {
			t.Fatalf("non-monotonie SSIM sous bruit amplitude %d : score=%d >= prev=%d", noiseAmp, score, prevScore)
		}
		prevScore = score
	}
}

func TestSSIM_ZeroAlloc(t *testing.T) {
	const w = 32
	const h = 32
	const sz = w * h

	img1 := make([]byte, sz)
	img2 := make([]byte, sz)
	dst := make([]byte, sz)
	tmp := make([]uint16, sz)

	allocsBlur := testing.AllocsPerRun(1000, func() {
		C2_gaussian_blur_2d(img1, dst, tmp, w, h, w)
	})
	if allocsBlur != 0 {
		t.Fatalf("C2_gaussian_blur_2d allocations = %.2f, attendu 0.0", allocsBlur)
	}

	allocsSSIM := testing.AllocsPerRun(100, func() {
		_ = C2_ssim_compute_milli(img1, img2, w, h, w)
	})
	if allocsSSIM != 0 {
		t.Fatalf("C2_ssim_compute_milli allocations = %.2f, attendu 0.0", allocsSSIM)
	}
}

func TestSSIMVsCOracle(t *testing.T) {
	tmpBin := filepath.Join(t.TempDir(), "ssim_c_oracle")
	srcDir := filepath.Join("..", "..", "sources")
	if _, err := os.Stat(filepath.Join(srcDir, "c2_ssim_gaussian.c")); err != nil {
		srcDir = filepath.Join("sources")
	}

	cmdBuild := exec.Command("gcc", "-O2", "-mavx2", "-Wall", "-Wextra", "-std=gnu99",
		"-I", srcDir,
		filepath.Join(srcDir, "test_ssim_oracle.c"),
		filepath.Join(srcDir, "c2_ssim_gaussian.c"),
		filepath.Join(srcDir, "c2_ssim_gaussian_avx2.c"),
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

	if !bytes.Contains(out, []byte("0x03FE76258E335041")) {
		t.Fatalf("l'oracle C n'a pas produit le hash GaussBlur attendu : %s", string(out))
	}
	if !bytes.Contains(out, []byte("353076")) {
		t.Fatalf("l'oracle C n'a pas produit le score KAT SSIM attendu : %s", string(out))
	}
	t.Logf("Sortie Oracle C (gcc -O2 -mavx2):\n%s", string(out))
}

func BenchmarkSSIM_GaussianBlur_1080p(b *testing.B) {
	const w = 1920
	const h = 1080
	const sz = w * h
	src := make([]byte, sz)
	dst := make([]byte, sz)
	tmp := make([]uint16, sz)

	b.SetBytes(int64(sz))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		C2_gaussian_blur_2d(src, dst, tmp, w, h, w)
	}
}

func BenchmarkSSIM_Compute_64x64(b *testing.B) {
	const w = 64
	const h = 64
	const sz = w * h
	img1 := make([]byte, sz)
	img2 := make([]byte, sz)

	b.SetBytes(int64(sz * 2))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = C2_ssim_compute_milli(img1, img2, w, h, w)
	}
}
