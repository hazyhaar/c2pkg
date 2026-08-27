package c2raster

import (
	"encoding/binary"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

type bezierTestCase struct {
	name   string
	x0, y0 float32
	x1, y1 float32
	x2, y2 float32
	w, h   int
}

type roundedRectTestCase struct {
	name      string
	w, h      int
	stride    int
	x, y      int
	rw, rh    int
	radius    int
	color     uint32
	baseColor uint32
}

func TestRasterVsCOracle(t *testing.T) {
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("gcc non disponible")
	}

	srcC, err := filepath.Abs(filepath.Join("sources", "c2raster.c"))
	if err != nil {
		t.Fatal(err)
	}
	srcH, err := filepath.Abs(filepath.Join("sources", "c2raster.h"))
	if err != nil {
		t.Fatal(err)
	}

	t.Run("BezierQuad_BitExact_Vs_COracle", func(t *testing.T) {
		tests := []bezierTestCase{
			{
				name: "DiagonalCurve",
				x0:   10, y0: 10,
				x1: 60, y1: 10,
				x2: 110, y2: 110,
				w: 128, h: 128,
			},
			{
				name: "SteepCurvature",
				x0:   20, y0: 100,
				x1: 64, y1: 10,
				x2: 108, y2: 100,
				w: 128, h: 128,
			},
			{
				name: "HorizontalFlat",
				x0:   10, y0: 64,
				x1: 64, y1: 64,
				x2: 118, y2: 64,
				w: 128, h: 128,
			},
			{
				name: "BoundaryClipping",
				x0:   -10, y0: -10,
				x1: 50, y1: 10,
				x2: 140, y2: 140,
				w: 128, h: 128,
			},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				dir := t.TempDir()
				mainC := fmt.Sprintf(`#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include "%s"
#include "%s"

int main(int argc, char **argv) {
    if (argc < 2) return 1;
    const char *out_path = argv[1];
    int w = %d;
    int h = %d;
    uint8_t *cov = (uint8_t *)calloc(w * h, sizeof(uint8_t));
    if (!cov) return 1;

    rasterize_bezier_quad(%ff, %ff, %ff, %ff, %ff, %ff, cov, w, h);

    FILE *f = fopen(out_path, "wb");
    if (!f) return 1;
    fwrite(cov, sizeof(uint8_t), w * h, f);
    fclose(f);
    free(cov);
    return 0;
}
`, srcH, srcC, tc.w, tc.h, tc.x0, tc.y0, tc.x1, tc.y1, tc.x2, tc.y2)

				mainPath := filepath.Join(dir, "oracle_main.c")
				if err := os.WriteFile(mainPath, []byte(mainC), 0o644); err != nil {
					t.Fatal(err)
				}

				bin := filepath.Join(dir, "oracle_bin")
				cmd := exec.Command("gcc", "-std=c99", "-Wall", "-Wextra", "-Werror", "-O2", "-o", bin, mainPath)
				if out, err := cmd.CombinedOutput(); err != nil {
					t.Fatalf("Compilation gcc oracle: %v\n%s", err, out)
				}

				cOutPath := filepath.Join(dir, "oracle_coverage.bin")
				if out, err := exec.Command(bin, cOutPath).CombinedOutput(); err != nil {
					t.Fatalf("Exécution oracle gcc: %v\n%s", err, out)
				}

				cBytes, err := os.ReadFile(cOutPath)
				if err != nil {
					t.Fatal(err)
				}

				goCov := make([]byte, tc.w*tc.h)
				Rasterize_bezier_quad(tc.x0, tc.y0, tc.x1, tc.y1, tc.x2, tc.y2, goCov, tc.w, tc.h)

				diffCount := 0
				for y := 0; y < tc.h; y++ {
					for x := 0; x < tc.w; x++ {
						idx := y*tc.w + x
						if goCov[idx] != cBytes[idx] {
							if diffCount < 5 {
								t.Errorf("Différence bit-exacte Bézier à (%d, %d): Go=%d vs C=%d", x, y, goCov[idx], cBytes[idx])
							}
							diffCount++
						}
					}
				}

				if diffCount > 0 {
					t.Fatalf("Échec de parité bit-exacte Bézier: %d divergences sur %d pixels", diffCount, tc.w*tc.h)
				}
			})
		}
	})

	t.Run("RoundedRect_BitExact_Vs_COracle", func(t *testing.T) {
		tests := []roundedRectTestCase{
			{
				name:      "SolidOpaque",
				w:         128,
				h:         128,
				stride:    128,
				x:         16,
				y:         16,
				rw:        96,
				rh:        64,
				radius:    12,
				color:     0xFF1080FF,
				baseColor: 0xFF202020,
			},
			{
				name:      "TranslucentBlend",
				w:         128,
				h:         128,
				stride:    128,
				x:         20,
				y:         20,
				rw:        80,
				rh:        80,
				radius:    20,
				color:     0x8000FF80,
				baseColor: 0xFF101010,
			},
			{
				name:      "PillShape",
				w:         128,
				h:         128,
				stride:    128,
				x:         10,
				y:         30,
				rw:        108,
				rh:        40,
				radius:    20,
				color:     0xFFFFFFFF,
				baseColor: 0xFF000000,
			},
			{
				name:      "ZeroRadius",
				w:         128,
				h:         128,
				stride:    128,
				x:         15,
				y:         15,
				rw:        60,
				rh:        40,
				radius:    0,
				color:     0xFFFFAA00,
				baseColor: 0xFF333333,
			},
			{
				name:      "ClippingBounds",
				w:         128,
				h:         128,
				stride:    128,
				x:         -10,
				y:         -10,
				rw:        80,
				rh:        80,
				radius:    16,
				color:     0xFF00AACC,
				baseColor: 0xFF151515,
			},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				dir := t.TempDir()
				mainC := fmt.Sprintf(`#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include "%s"
#include "%s"

int main(int argc, char **argv) {
    if (argc < 2) return 1;
    const char *out_path = argv[1];
    int w = %d;
    int h = %d;
    int stride = %d;
    uint32_t *fb = (uint32_t *)malloc(w * h * sizeof(uint32_t));
    if (!fb) return 1;

    for (int i = 0; i < w * h; i++) {
        fb[i] = 0x%08Xu;
    }

    rasterize_rounded_rect(fb, stride, %d, %d, %d, %d, %d, 0x%08Xu);

    FILE *f = fopen(out_path, "wb");
    if (!f) return 1;
    fwrite(fb, sizeof(uint32_t), w * h, f);
    fclose(f);
    free(fb);
    return 0;
}
`, srcH, srcC, tc.w, tc.h, tc.stride, tc.baseColor, tc.x, tc.y, tc.rw, tc.rh, tc.radius, tc.color)

				mainPath := filepath.Join(dir, "oracle_main.c")
				if err := os.WriteFile(mainPath, []byte(mainC), 0o644); err != nil {
					t.Fatal(err)
				}

				bin := filepath.Join(dir, "oracle_bin")
				cmd := exec.Command("gcc", "-std=c99", "-Wall", "-Wextra", "-Werror", "-O2", "-o", bin, mainPath)
				if out, err := cmd.CombinedOutput(); err != nil {
					t.Fatalf("Compilation gcc oracle: %v\n%s", err, out)
				}

				cOutPath := filepath.Join(dir, "oracle_pixels.bin")
				if out, err := exec.Command(bin, cOutPath).CombinedOutput(); err != nil {
					t.Fatalf("Exécution oracle gcc: %v\n%s", err, out)
				}

				cBytes, err := os.ReadFile(cOutPath)
				if err != nil {
					t.Fatal(err)
				}

				cPts := make([]uint32, tc.w*tc.h)
				for i := 0; i < len(cPts); i++ {
					cPts[i] = binary.LittleEndian.Uint32(cBytes[i*4 : (i+1)*4])
				}

				goFb := make([]uint32, tc.w*tc.h)
				for i := range goFb {
					goFb[i] = tc.baseColor
				}

				Rasterize_rounded_rect(goFb, tc.stride, tc.x, tc.y, tc.rw, tc.rh, tc.radius, tc.color)

				diffCount := 0
				for y := 0; y < tc.h; y++ {
					for x := 0; x < tc.w; x++ {
						idx := y*tc.stride + x
						if goFb[idx] != cPts[idx] {
							if diffCount < 5 {
								t.Errorf("Différence bit-exacte Rectangle Arrondi à (%d, %d): Go=0x%08X vs C=0x%08X", x, y, goFb[idx], cPts[idx])
							}
							diffCount++
						}
					}
				}

				if diffCount > 0 {
					t.Fatalf("Échec de parité bit-exacte Rectangle Arrondi: %d divergences sur %d pixels", diffCount, tc.w*tc.h)
				}
			})
		}
	})

	t.Run("CompositeScene_BitExact_Vs_COracle", func(t *testing.T) {
		const w = 160
		const h = 160
		const stride = 160

		dir := t.TempDir()
		mainC := fmt.Sprintf(`#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include "%s"
#include "%s"

int main(int argc, char **argv) {
    if (argc < 2) return 1;
    const char *out_path = argv[1];
    uint32_t *fb = (uint32_t *)malloc(%d * %d * sizeof(uint32_t));
    if (!fb) return 1;

    for (int i = 0; i < %d * %d; i++) {
        fb[i] = 0xFF18181Bu;
    }

    /* Plusieurs rectangles arrondis */
    rasterize_rounded_rect(fb, %d, 10, 10, 60, 40, 8, 0xFF3D3D44u);
    rasterize_rounded_rect(fb, %d, 80, 10, 70, 50, 15, 0xCCFF5500u);
    rasterize_rounded_rect(fb, %d, 20, 70, 120, 70, 24, 0xAA00CC88u);

    FILE *f = fopen(out_path, "wb");
    if (!f) return 1;
    fwrite(fb, sizeof(uint32_t), %d * %d, f);
    fclose(f);
    free(fb);
    return 0;
}
`, srcH, srcC, w, h, w, h, stride, stride, stride, w, h)

		mainPath := filepath.Join(dir, "composite_oracle.c")
		if err := os.WriteFile(mainPath, []byte(mainC), 0o644); err != nil {
			t.Fatal(err)
		}

		bin := filepath.Join(dir, "composite_bin")
		cmd := exec.Command("gcc", "-std=c99", "-Wall", "-Wextra", "-Werror", "-O2", "-o", bin, mainPath)
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("Compilation gcc oracle composite: %v\n%s", err, out)
		}

		cOutPath := filepath.Join(dir, "composite_pixels.bin")
		if out, err := exec.Command(bin, cOutPath).CombinedOutput(); err != nil {
			t.Fatalf("Exécution oracle composite: %v\n%s", err, out)
		}

		cBytes, err := os.ReadFile(cOutPath)
		if err != nil {
			t.Fatal(err)
		}

		cPts := make([]uint32, w*h)
		for i := 0; i < len(cPts); i++ {
			cPts[i] = binary.LittleEndian.Uint32(cBytes[i*4 : (i+1)*4])
		}

		goFb := make([]uint32, w*h)
		for i := range goFb {
			goFb[i] = 0xFF18181B
		}

		Rasterize_rounded_rect(goFb, stride, 10, 10, 60, 40, 8, 0xFF3D3D44)
		Rasterize_rounded_rect(goFb, stride, 80, 10, 70, 50, 15, 0xCCFF5500)
		Rasterize_rounded_rect(goFb, stride, 20, 70, 120, 70, 24, 0xAA00CC88)

		diffCount := 0
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				idx := y*stride + x
				if goFb[idx] != cPts[idx] {
					if diffCount < 5 {
						t.Errorf("Différence composite à (%d, %d): Go=0x%08X vs C=0x%08X", x, y, goFb[idx], cPts[idx])
					}
					diffCount++
				}
			}
		}

		if diffCount > 0 {
			t.Fatalf("Échec composite: %d divergences sur %d pixels", diffCount, w*h)
		}
	})
}
