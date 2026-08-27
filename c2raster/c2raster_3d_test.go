// SPDX-License-Identifier: Apache-2.0 OR MIT
package c2raster

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

type tri3dTestCase struct {
	name   string
	tri    Triangle3D
	w, h   int
	stride int
}

func TestRaster3DVsCOracle(t *testing.T) {
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("gcc non disponible sur l'hôte")
	}

	srcC, err := filepath.Abs(filepath.Join("sources", "c2raster_3d.c"))
	if err != nil {
		t.Fatal(err)
	}
	srcH, err := filepath.Abs(filepath.Join("sources", "c2raster_3d.h"))
	if err != nil {
		t.Fatal(err)
	}

	tests := []tri3dTestCase{
		{
			name: "FrontFacingEquilateral",
			w:    128, h: 128, stride: 128,
			tri: Triangle3D{
				V0: Vertex3D{X: 64, Y: 20, Z: 0.2, InvW: 1.0, Color: 0xFF0000FF}, // Rouge (RGBA: R=0xFF)
				V1: Vertex3D{X: 104, Y: 90, Z: 0.5, InvW: 0.8, Color: 0xFF00FF00}, // Vert
				V2: Vertex3D{X: 24, Y: 90, Z: 0.8, InvW: 0.5, Color: 0xFFFF0000},  // Bleu
			},
		},
		{
			name: "SteepPerspectiveTriangle",
			w:    128, h: 128, stride: 128,
			tri: Triangle3D{
				V0: Vertex3D{X: 10, Y: 10, Z: 0.05, InvW: 2.0, Color: 0xFFFFFFFF},
				V1: Vertex3D{X: 118, Y: 30, Z: 0.95, InvW: 0.1, Color: 0xFF102030},
				V2: Vertex3D{X: 50, Y: 120, Z: 0.50, InvW: 0.5, Color: 0xFF808080},
			},
		},
		{
			name: "BoundaryClippedTriangle",
			w:    64, h: 64, stride: 64,
			tri: Triangle3D{
				V0: Vertex3D{X: -20, Y: -10, Z: 0.1, InvW: 1.0, Color: 0xFFAA00AA},
				V1: Vertex3D{X: 80, Y: 10, Z: 0.4, InvW: 0.8, Color: 0xFF00AAAA},
				V2: Vertex3D{X: 30, Y: 80, Z: 0.7, InvW: 0.6, Color: 0xFFAAAA00},
			},
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
    if (argc < 3) return 1;
    const char *out_color = argv[1];
    const char *out_depth = argv[2];
    int w = %d;
    int h = %d;
    int stride = %d;

    uint32_t *colors = (uint32_t *)calloc(stride * h, sizeof(uint32_t));
    float *depths = (float *)malloc(stride * h * sizeof(float));
    for (int i = 0; i < stride * h; i++) depths[i] = 1.0f;

    c2_triangle3d_t tri = {
        .v0 = { .x = %ff, .y = %ff, .z = %ff, .inv_w = %ff, .color = 0x%X },
        .v1 = { .x = %ff, .y = %ff, .z = %ff, .inv_w = %ff, .color = 0x%X },
        .v2 = { .x = %ff, .y = %ff, .z = %ff, .inv_w = %ff, .color = 0x%X },
    };

    c2_rasterize_triangle3d(&tri, colors, depths, w, h, stride, 0, 0, w, h);

    FILE *fc = fopen(out_color, "wb");
    if (!fc) return 1;
    fwrite(colors, sizeof(uint32_t), stride * h, fc);
    fclose(fc);

    FILE *fd = fopen(out_depth, "wb");
    if (!fd) return 1;
    fwrite(depths, sizeof(float), stride * h, fd);
    fclose(fd);

    free(colors);
    free(depths);
    return 0;
}
`, srcH, srcC, tc.w, tc.h, tc.stride,
				tc.tri.V0.X, tc.tri.V0.Y, tc.tri.V0.Z, tc.tri.V0.InvW, tc.tri.V0.Color,
				tc.tri.V1.X, tc.tri.V1.Y, tc.tri.V1.Z, tc.tri.V1.InvW, tc.tri.V1.Color,
				tc.tri.V2.X, tc.tri.V2.Y, tc.tri.V2.Z, tc.tri.V2.InvW, tc.tri.V2.Color,
			)

			mainPath := filepath.Join(dir, "oracle_main.c")
			if err := os.WriteFile(mainPath, []byte(mainC), 0600); err != nil {
				t.Fatal(err)
			}

			binPath := filepath.Join(dir, "oracle_bin")
			cmdBuild := exec.Command("gcc", "-O2", "-Wall", "-Wextra", mainPath, "-o", binPath, "-lm")
			if out, err := cmdBuild.CombinedOutput(); err != nil {
				t.Fatalf("échec compilation oracle C: %v, sortie: %s", err, string(out))
			}

			outColor := filepath.Join(dir, "color.bin")
			outDepth := filepath.Join(dir, "depth.bin")
			cmdRun := exec.Command(binPath, outColor, outDepth)
			if out, err := cmdRun.CombinedOutput(); err != nil {
				t.Fatalf("échec exécution oracle C: %v, sortie: %s", err, string(out))
			}

			cColorBytes, err := os.ReadFile(outColor)
			if err != nil {
				t.Fatal(err)
			}
			cDepthBytes, err := os.ReadFile(outDepth)
			if err != nil {
				t.Fatal(err)
			}

			// Exécution Go
			ctx := NewRasterContext3D(tc.w, tc.h, tc.stride)
			ctx.Clear(0, 1.0)
			ctx.RasterizeTriangle(&tc.tri)

			var goColorBuf bytes.Buffer
			for _, c := range ctx.ColorBuffer {
				_ = binary.Write(&goColorBuf, binary.LittleEndian, c)
			}

			var goDepthBuf bytes.Buffer
			for _, d := range ctx.DepthBuffer {
				_ = binary.Write(&goDepthBuf, binary.LittleEndian, d)
			}

			if !bytes.Equal(goColorBuf.Bytes(), cColorBytes) {
				t.Errorf("divergence bit-exacte du tampon couleur sur %s", tc.name)
			}

			if !bytes.Equal(goDepthBuf.Bytes(), cDepthBytes) {
				t.Errorf("divergence bit-exacte du tampon profondeur Z sur %s", tc.name)
			}
		})
	}
}

func TestRaster3DSharedEdgeWatertight(t *testing.T) {
	// Deux triangles adjacents formant un quad rectangle parfait [10, 110] x [10, 110]
	w, h := 128, 128
	ctx := NewRasterContext3D(w, h, w)
	ctx.Clear(0, 1.0)

	// Triangle 1 : (10,10) -> (110,10) -> (10,110)
	t1 := Triangle3D{
		V0: Vertex3D{X: 10, Y: 10, Z: 0.5, InvW: 1.0, Color: 0xFF0000FF},
		V1: Vertex3D{X: 110, Y: 10, Z: 0.5, InvW: 1.0, Color: 0xFF0000FF},
		V2: Vertex3D{X: 10, Y: 110, Z: 0.5, InvW: 1.0, Color: 0xFF0000FF},
	}

	// Triangle 2 : (110,10) -> (110,110) -> (10,110)
	t2 := Triangle3D{
		V0: Vertex3D{X: 110, Y: 10, Z: 0.5, InvW: 1.0, Color: 0xFF00FF00},
		V1: Vertex3D{X: 110, Y: 110, Z: 0.5, InvW: 1.0, Color: 0xFF00FF00},
		V2: Vertex3D{X: 10, Y: 110, Z: 0.5, InvW: 1.0, Color: 0xFF00FF00},
	}

	ctx.RasterizeTriangle(&t1)
	ctx.RasterizeTriangle(&t2)

	// Vérification de couverture complète sans aucun trou dans [10, 109] x [10, 109]
	holes := 0
	for y := 10; y < 110; y++ {
		for x := 10; x < 110; x++ {
			if ctx.ColorBuffer[y*w+x] == 0 {
				holes++
			}
		}
	}
	if holes > 0 {
		t.Errorf("trou de rastérisation sur l'arête partagée (%d pixels non couverts)", holes)
	}
}

func TestRaster3DSerialVsParallelParity(t *testing.T) {
	w, h := 256, 256
	ctxSerial := NewRasterContext3D(w, h, w)
	ctxParallel := NewRasterContext3D(w, h, w)

	ctxSerial.Clear(0, 1.0)
	ctxParallel.Clear(0, 1.0)

	rng := rand.New(rand.NewSource(1337))
	tris := make([]Triangle3D, 50)
	for i := range tris {
		cx := float32(rng.Intn(200) + 20)
		cy := float32(rng.Intn(200) + 20)
		tris[i] = Triangle3D{
			V0: Vertex3D{X: cx - 15, Y: cy - 15, Z: rng.Float32()*0.8 + 0.1, InvW: 1.0, Color: 0xFFFF0000},
			V1: Vertex3D{X: cx + 20, Y: cy - 5, Z: rng.Float32()*0.8 + 0.1, InvW: 0.8, Color: 0xFF00FF00},
			V2: Vertex3D{X: cx, Y: cy + 25, Z: rng.Float32()*0.8 + 0.1, InvW: 0.6, Color: 0xFF0000FF},
		}
	}

	for i := range tris {
		ctxSerial.RasterizeTriangle(&tris[i])
	}

	ctxParallel.RasterizeTrianglesParallel(tris)

	for i := range ctxSerial.ColorBuffer {
		if ctxSerial.ColorBuffer[i] != ctxParallel.ColorBuffer[i] {
			t.Fatalf("divergence série vs parallèle sur le pixel %d : Série=%08X, Parallèle=%08X",
				i, ctxSerial.ColorBuffer[i], ctxParallel.ColorBuffer[i])
		}
		if ctxSerial.DepthBuffer[i] != ctxParallel.DepthBuffer[i] {
			t.Fatalf("divergence profondeur série vs parallèle sur le pixel %d", i)
		}
	}
}

func TestRaster3DZeroAlloc(t *testing.T) {
	ctx := NewRasterContext3D(640, 480, 640)
	tri := Triangle3D{
		V0: Vertex3D{X: 100, Y: 100, Z: 0.1, InvW: 1.0, Color: 0xFFFFFFFF},
		V1: Vertex3D{X: 300, Y: 150, Z: 0.5, InvW: 0.5, Color: 0xFF00FF00},
		V2: Vertex3D{X: 200, Y: 400, Z: 0.8, InvW: 0.2, Color: 0xFF0000FF},
	}

	allocs := testing.AllocsPerRun(100, func() {
		ctx.Clear(0, 1.0)
		ctx.RasterizeTriangle(&tri)
	})

	if allocs != 0 {
		t.Errorf("échec invariant ARCHTIME : attendu 0 allocs/op, obtenu %f", allocs)
	}
}
