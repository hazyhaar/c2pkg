package c2painter

import (
	"math/rand"
	"testing"
)

func TestSIMD_Div255_BitExactness(t *testing.T) {
	// Vérification exhaustive sur l'intégralité du domaine [0, 255*255 = 65025]
	for val := uint32(0); val <= 255*255; val++ {
		expected := (val + 127) / 255
		got := Div255(val)
		if got != expected {
			t.Fatalf("Disparité Div255 pour t=%d: attendu %d, obtenu %d", val, expected, got)
		}
	}
}

func TestSIMD_BlendPixel_BitExactVsC2BlendPixel(t *testing.T) {
	// 1. Cas limites particuliers
	edgeColors := []uint32{
		0x00000000, // Transparent total
		0xFF000000, // Noir opaque
		0xFFFFFFFF, // Blanc opaque
		0x80FFFFFF, // Blanc 50%
		0x80123456, // Couleur arbitraire 50%
		0x01123456, // Quasi-transparent
		0xFE123456, // Quasi-opaque
	}

	for _, src := range edgeColors {
		for _, dst := range edgeColors {
			expected := C2_blend_pixel(dst, src)
			got := BlendPixelSIMD(dst, src)
			if got != expected {
				t.Fatalf("Disparité BlendPixelSIMD(dst=0x%08x, src=0x%08x): attendu 0x%08x, obtenu 0x%08x", dst, src, expected, got)
			}
		}
	}

	// 2. Épreuve pseudo-aléatoire massive (100 000 paires)
	rng := rand.New(rand.NewSource(42))
	for i := 0; i < 100000; i++ {
		dst := rng.Uint32()
		src := rng.Uint32()

		expected := C2_blend_pixel(dst, src)
		got := BlendPixelSIMD(dst, src)
		if got != expected {
			t.Fatalf("Disparité aléatoire BlendPixelSIMD(dst=0x%08x, src=0x%08x): attendu 0x%08x, obtenu 0x%08x", dst, src, expected, got)
		}
	}
}

func TestSIMD_FillRectSIMD_ParityVsC2FillRect(t *testing.T) {
	testSizes := []struct {
		w, h int
	}{
		{32, 32},
		{64, 64},
		{127, 83}, // Dimensions non alignées
		{100, 100},
	}

	testColors := []uint32{
		PackRGBA(255, 0, 0, 255),   // Opaque rouge
		PackRGBA(0, 255, 0, 128),   // Translucide vert
		PackRGBA(0, 0, 255, 64),    // Translucide bleu faible
		PackRGBA(255, 255, 255, 0), // Transparent
	}

	for _, sz := range testSizes {
		for _, col := range testColors {
			// Surface scalaire oracle
			surfRef := NewSurface(sz.w, sz.h)
			pRef := NewPainter(surfRef)
			pRef.Clear(PackRGBA(10, 20, 30, 255))
			C2_fill_rect(&pRef.ctx, 5, 7, sz.w-12, sz.h-15, col)

			// Surface vectorielle SIMD
			surfSIMD := NewSurface(sz.w, sz.h)
			pSIMD := NewPainter(surfSIMD)
			pSIMD.Clear(PackRGBA(10, 20, 30, 255))
			FillRectSIMD(&pSIMD.ctx, 5, 7, sz.w-12, sz.h-15, col)

			// Contrôle bit-exact pixel à pixel
			for y := 0; y < sz.h; y++ {
				for x := 0; x < sz.w; x++ {
					refPix := surfRef.Pixels[y*sz.w+x]
					simdPix := surfSIMD.Pixels[y*sz.w+x]
					if refPix != simdPix {
						t.Fatalf("Divergence FillRectSIMD à (%d,%d) pour sz=%dx%d col=0x%08x: ref=0x%08x simd=0x%08x",
							x, y, sz.w, sz.h, col, refPix, simdPix)
					}
				}
			}
		}
	}
}

func TestSIMD_BlendPixels8AVX2(t *testing.T) {
	var dst [8]uint32
	var src [8]uint32

	for i := 0; i < 8; i++ {
		dst[i] = PackRGBA(uint8(i*20), uint8(i*30), uint8(i*40), 255)
		src[i] = PackRGBA(uint8(255-i*20), uint8(255-i*30), 100, 128)
	}

	var dstRef [8]uint32
	copy(dstRef[:], dst[:])

	for i := 0; i < 8; i++ {
		dstRef[i] = C2_blend_pixel(dstRef[i], src[i])
	}

	BlendPixels8AVX2(dst[:], src[:])

	for i := 0; i < 8; i++ {
		if dst[i] != dstRef[i] {
			t.Fatalf("Divergence BlendPixels8AVX2 à l'index %d: attendu 0x%08x, obtenu 0x%08x", i, dstRef[i], dst[i])
		}
	}
}

func BenchmarkFillRect_Scalar(b *testing.B) {
	surf := NewSurface(1920, 1080)
	p := NewPainter(surf)
	col := PackRGBA(255, 128, 64, 180)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C2_fill_rect(&p.ctx, 100, 100, 800, 600, col)
	}
}

func BenchmarkFillRect_SIMD(b *testing.B) {
	surf := NewSurface(1920, 1080)
	p := NewPainter(surf)
	col := PackRGBA(255, 128, 64, 180)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FillRectSIMD(&p.ctx, 100, 100, 800, 600, col)
	}
}
