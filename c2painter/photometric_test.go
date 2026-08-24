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
