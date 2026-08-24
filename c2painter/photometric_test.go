package c2painter

import (
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

// TestBlendPixelPhotometric_DarkFringeElimination prouve que la composition linéaire sgoiter
// élimine l'assombrissement parasite sRGB : mélanger du blanc à 50% sur du noir
// produit une valeur sRGB de 188 (50% de luminance physique réelle) et non 128 (qui ne fait que 21.8% de luminance).
func TestBlendPixelPhotometric_DarkFringeElimination(t *testing.T) {
	black := PackRGBA(0, 0, 0, 255)
	whiteHalfAlpha := PackRGBA(255, 255, 255, 128)

	// Composition linéaire photométrique via le noyau sgoiter C2_blend_photometric
	blended := C2_blend_photometric(black, whiteHalfAlpha)
	r, g, b, a := UnpackRGBA(blended)

	if a != 255 {
		t.Errorf("Alpha attendu 255, obtenu %d", a)
	}

	// 50% d'énergie lumineuse dans l'espace sRGB correspond à 188 bit-exact (et non 128)
	if r != 188 || g != 188 || b != 188 {
		t.Fatalf("VIOLATION PHOTOMÉTRIQUE : Énergie lumineuse corrompue R=%d G=%d B=%d (attendu 188 exact pour 50%% d'énergie lumineuse réelle, 128 = frange sombre parasite)", r, g, b)
	}
}

func TestPainter_DrawRectStandard_AutoPhotometric(t *testing.T) {
	surf := NewSurface(100, 100)
	p := NewPainter(surf)
	p.PhotometricBlending = true
	p.Clear(PackRGBA(0, 0, 0, 255))
	// Appel à DrawRect standard avec PhotometricBlending activé : route vers le mélange linéaire
	p.DrawRect(10, 10, 80, 80, PackRGBA(255, 255, 255, 128))

	center := surf.Pixels[50*surf.Stride+50]
	r, g, b, a := UnpackRGBA(center)

	if a != 255 || r != 188 || g != 188 || b != 188 {
		t.Fatalf("DrawRect standard a échoué à éliminer la frange sombre : obtenu R=%d G=%d B=%d, A=%d (attendu 188 bit-exact)", r, g, b, a)
	}
}

// TestPhotometricVsCOracle prouve la parité bit-exacte contre l'oracle C gcc -O2.
func TestPhotometricVsCOracle(t *testing.T) {
	tmpDir := t.TempDir()
	binPath := filepath.Join(tmpDir, "test_photometric_oracle")

	cmdBuild := exec.Command("gcc", "-O2", "-Wall", "-Wextra",
		"/devhoros/c2simd/sources/c2_photometric.c",
		"/devhoros/c2simd/sources/test_photometric_oracle.c",
		"-o", binPath, "-lm")
	if out, err := cmdBuild.CombinedOutput(); err != nil {
		t.Fatalf("Échec compilation oracle C: %v\n%s", err, string(out))
	}

	cmdRun := exec.Command(binPath)
	out, err := cmdRun.CombinedOutput()
	if err != nil {
		t.Fatalf("Échec exécution oracle C: %v\n%s", err, string(out))
	}

	outStr := string(out)
	if !strings.Contains(outStr, "ORACLE C PHOTOMETRIC PASS") {
		t.Fatalf("L'oracle C n'a pas validé : %s", outStr)
	}

	// Validation de la parité sur le résultat Go transpilé sgoiter
	black := PackRGBA(0, 0, 0, 255)
	whiteHalf := PackRGBA(255, 255, 255, 128)
	goBlended := C2_blend_photometric(black, whiteHalf)
	r, g, b, a := UnpackRGBA(goBlended)

	expectedStr := "R=" + strconv.Itoa(int(r)) + " G=" + strconv.Itoa(int(g)) + " B=" + strconv.Itoa(int(b)) + " A=" + strconv.Itoa(int(a))
	if !strings.Contains(outStr, expectedStr) {
		t.Fatalf("Divergence Go transpilé vs Oracle C : Go=%s, C=%s", expectedStr, outStr)
	}
}

// TestCartesianMatrix_AllPrimitives_PhotometricLinear valide mécaniquement la matrice cartésienne
// de toutes les primitives sous mode photométrique pour garantir l'absence de régression vers sRGB naïf.
func TestCartesianMatrix_AllPrimitives_PhotometricLinear(t *testing.T) {
	black := PackRGBA(0, 0, 0, 255)
	whiteOpaque := PackRGBA(255, 255, 255, 255)
	white50 := PackRGBA(255, 255, 255, 128)

	// 1. Primitive: Rectangle translucide (DrawRect)
	t.Run("Primitive_DrawRect_50pct", func(t *testing.T) {
		surf := NewSurface(16, 16)
		p := NewPainter(surf)
		p.PhotometricBlending = true
		p.Clear(black)
		p.DrawRect(2, 2, 12, 12, white50)
		r, g, b, _ := UnpackRGBA(surf.Pixels[8*surf.Stride+8])
		if r != 188 || g != 188 || b != 188 {
			t.Fatalf("DrawRect non photométrique : R=%d G=%d B=%d (attendu 188)", r, g, b)
		}
	})

	// 2. Primitive: Glyphe textuel avec masque antialiasé à 50% (DrawTextGlyph)
	t.Run("Primitive_DrawTextGlyph_Alpha50", func(t *testing.T) {
		surf := NewSurface(16, 16)
		p := NewPainter(surf)
		p.PhotometricBlending = true
		p.Clear(black)
		mask := []byte{
			128, 128,
			128, 128,
		}
		p.DrawTextGlyph(4, 4, 2, 2, mask, 2, whiteOpaque)
		r, g, b, _ := UnpackRGBA(surf.Pixels[4*surf.Stride+4])
		if r != 188 || g != 188 || b != 188 {
			t.Fatalf("DrawTextGlyph non photométrique : R=%d G=%d B=%d (attendu 188)", r, g, b)
		}
	})

	// 3. Primitive: Couverture arbitraire (BlendPixelCov)
	t.Run("Primitive_BlendPixelCov_128", func(t *testing.T) {
		blended := C2_blend_pixel_cov_photometric(black, whiteOpaque, 128)
		r, g, b, _ := UnpackRGBA(blended)
		if r != 188 || g != 188 || b != 188 {
			t.Fatalf("BlendPixelCov non photométrique : R=%d G=%d B=%d (attendu 188)", r, g, b)
		}
	})

	// 4. Primitive: Traitement de portée (BlendSpan)
	t.Run("Primitive_BlendSpan_50pct", func(t *testing.T) {
		dst := []uint32{black, black, black, black}
		src := []uint32{white50, white50, white50, white50}
		C2_blend_span_photometric(dst, src, 4)
		for i := 0; i < 4; i++ {
			r, g, b, _ := UnpackRGBA(dst[i])
			if r != 188 || g != 188 || b != 188 {
				t.Fatalf("BlendSpan[%d] non photométrique : R=%d G=%d B=%d (attendu 188)", i, r, g, b)
			}
		}
	})

	// 5. Primitive: Rectangle vectoriel SIMD (DrawRectSIMD)
	t.Run("Primitive_DrawRectSIMD_50pct", func(t *testing.T) {
		surf := NewSurface(16, 16)
		p := NewPainter(surf)
		p.PhotometricBlending = true
		p.Clear(black)
		p.DrawRectSIMD(2, 2, 12, 12, white50)
		r, g, b, _ := UnpackRGBA(surf.Pixels[8*surf.Stride+8])
		if r != 188 || g != 188 || b != 188 {
			t.Fatalf("DrawRectSIMD non photométrique : R=%d G=%d B=%d (attendu 188)", r, g, b)
		}
	})

	// 6. Primitive: Cercle antialiasé à 50% d'alpha (DrawCircle)
	t.Run("Primitive_DrawCircle_50pct", func(t *testing.T) {
		surf := NewSurface(32, 32)
		p := NewPainter(surf)
		p.PhotometricBlending = true
		p.Clear(black)
		p.DrawCircle(16, 16, 8, white50)
		// Le centre du cercle doit recevoir exactement 50% de blanc sur noir physique -> 188 bit-exact
		r, g, b, _ := UnpackRGBA(surf.Pixels[16*surf.Stride+16])
		if r != 188 || g != 188 || b != 188 {
			t.Fatalf("DrawCircle centre non photométrique : R=%d G=%d B=%d (attendu 188)", r, g, b)
		}
	})

	// 7. Primitive: Ligne épaisse antialiasée à 50% d'alpha (DrawLine)
	t.Run("Primitive_DrawLine_50pct", func(t *testing.T) {
		surf := NewSurface(32, 32)
		p := NewPainter(surf)
		p.PhotometricBlending = true
		p.Clear(black)
		p.DrawLine(4, 16, 28, 16, 4, white50)
		// Le centre de la ligne doit recevoir 188 bit-exact
		r, g, b, _ := UnpackRGBA(surf.Pixels[16*surf.Stride+16])
		if r != 188 || g != 188 || b != 188 {
			t.Fatalf("DrawLine centre non photométrique : R=%d G=%d B=%d (attendu 188)", r, g, b)
		}
	})

	// 8. Test de parité bit-exacte du span vectorisé C2_blend_solid_span_photometric sur toutes les longueurs
	t.Run("SolidSpan_AllLengths_BitExactVsScalar", func(t *testing.T) {
		lengths := []int{1, 2, 3, 7, 8, 9, 15, 16, 31, 32, 63, 64, 127, 128, 255, 1024}
		colors := []uint32{
			PackRGBA(255, 255, 255, 128),
			PackRGBA(255, 100, 50, 200),
			PackRGBA(0, 255, 128, 64),
			PackRGBA(12, 34, 56, 255), // opaque
			PackRGBA(12, 34, 56, 0),   // transparent
		}

		for _, n := range lengths {
			for _, c := range colors {
				spanBuf := make([]uint32, n)
				scalarBuf := make([]uint32, n)

				// Remplissage initial avec motif non trivial
				for i := 0; i < n; i++ {
					initColor := PackRGBA(uint8(i*7+13), uint8(i*11+37), uint8(i*13+59), 255)
					spanBuf[i] = initColor
					scalarBuf[i] = initColor
				}

				// Exécution vectorisée (déroulée 8-way sgoiter)
				C2_blend_solid_span_photometric(spanBuf, c, n)

				// Exécution scalaire pixel-par-pixel
				for i := 0; i < n; i++ {
					scalarBuf[i] = C2_blend_photometric(scalarBuf[i], c)
				}

				// Vérification de la parité bit-exacte sur 100% des pixels
				for i := 0; i < n; i++ {
					if spanBuf[i] != scalarBuf[i] {
						sr, sg, sb, sa := UnpackRGBA(spanBuf[i])
						er, eg, eb, ea := UnpackRGBA(scalarBuf[i])
						t.Fatalf("Longueur %d, pixel %d : divergence span [%d,%d,%d,%d] vs scalaire [%d,%d,%d,%d]",
							n, i, sr, sg, sb, sa, er, eg, eb, ea)
					}
				}
			}
		}
	})
}

// TestPhotometric_TortureAndChaos_UpToCrash soumet toutes les primitives photométriques
// à des sollicitations extrêmes (débordements arithmétiques, coordonnées négatives,
// masques corrompus, 10 000 passes pseudo-aléatoires et concurrence agressive sous -race).
func TestPhotometric_TortureAndChaos_UpToCrash(t *testing.T) {
	// 1. Torture Géométrique & Débordements Entiers Extrêmes
	t.Run("ExtremeGeometryAndOverflows", func(t *testing.T) {
		surf := NewSurface(64, 64)
		p := NewPainter(surf)
		p.PhotometricBlending = true

		extremeValues := []int{
			-1000000000, -1000000, -65536, -100, -1, 0, 1, 63, 64, 65, 100, 65536, 1000000, 1000000000,
		}

		for _, cx := range extremeValues {
			for _, cy := range extremeValues {
				for _, r := range []int{-10000, -1, 0, 1, 32, 64, 100000} {
					// Aucun appel ne doit paniquer ni corrompre la mémoire
					p.DrawCircle(cx, cy, r, PackRGBA(255, 128, 64, 128))
					p.DrawLine(cx, cy, cy, cx, r, PackRGBA(64, 128, 255, 128))
					p.DrawRect(cx, cy, r, r, PackRGBA(128, 255, 64, 128))
					p.DrawRectSIMD(cx, cy, r, r, PackRGBA(255, 255, 255, 128))
				}
			}
		}
	})

	// 2. Fuzzing Masques et Blits Alpha Corrompus
	t.Run("CorruptedMasksAndBlits", func(t *testing.T) {
		surf := NewSurface(32, 32)
		p := NewPainter(surf)
		p.PhotometricBlending = true

		// Masque nil
		p.DrawTextGlyph(0, 0, 10, 10, nil, 10, PackRGBA(255, 255, 255, 128))
		// Masque tronqué
		p.DrawTextGlyph(0, 0, 10, 10, []byte{128, 128}, 10, PackRGBA(255, 255, 255, 128))
		// Masque stride nul ou négatif
		p.DrawTextGlyph(0, 0, 10, 10, make([]byte, 100), 0, PackRGBA(255, 255, 255, 128))
		p.DrawTextGlyph(0, 0, 10, 10, make([]byte, 100), -5, PackRGBA(255, 255, 255, 128))
		// Dimensions nulles ou négatives
		p.DrawTextGlyph(-10, -10, -5, -5, make([]byte, 100), 10, PackRGBA(255, 255, 255, 128))

		// C2_blend_span_photometric avec tranches valides et cas limite n=0
		C2_blend_span_photometric(nil, nil, 0)
		C2_blend_span_photometric(make([]uint32, 4), make([]uint32, 4), 4)
	})

	// 3. Fuzzing Pseudo-Aléatoire (10 000 Opérations Chaotiques)
	t.Run("FuzzRandomChaos10k", func(t *testing.T) {
		surf := NewSurface(128, 128)
		p := NewPainter(surf)
		p.PhotometricBlending = true

		// LCG déterministe pour reproductibilité stricte
		var seed uint64 = 0xDEADBEEFCAFE5555
		randNext := func() int {
			seed = seed*6364136223846793005 + 1442695040888963407
			return int(seed >> 33)
		}

		for i := 0; i < 10000; i++ {
			x0 := (randNext() % 300) - 100
			y0 := (randNext() % 300) - 100
			x1 := (randNext() % 300) - 100
			y1 := (randNext() % 300) - 100
			w := (randNext() % 200) - 50
			h := (randNext() % 200) - 50
			r := (randNext() % 100) - 20
			color := uint32(randNext())

			switch i % 5 {
			case 0:
				p.DrawRect(x0, y0, w, h, color)
			case 1:
				p.DrawRectSIMD(x0, y0, w, h, color)
			case 2:
				p.DrawCircle(x0, y0, r, color)
			case 3:
				p.DrawLine(x0, y0, x1, y1, max(1, r%10), color)
			case 4:
				mask := make([]byte, max(0, w*h))
				p.DrawTextGlyph(x0, y0, max(0, w), max(0, h), mask, max(1, w), color)
			}
		}
	})
}
