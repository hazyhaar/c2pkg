// SPDX-License-Identifier: MIT
package c2meshopt

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"
)

func TestMeshOptVsCOracle(t *testing.T) {
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("gcc non disponible")
	}

	srcC, err := filepath.Abs(filepath.Join("sources", "meshopt.c"))
	if err != nil {
		t.Fatal(err)
	}
	srcH, err := filepath.Abs(filepath.Join("sources", "meshopt.h"))
	if err != nil {
		t.Fatal(err)
	}

	dir := t.TempDir()
	mainC := fmt.Sprintf(`#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include "%s"
#include "%s"

int main(int argc, char **argv) {
    if (argc < 4) return 1;
    size_t index_count = (size_t)atoi(argv[1]);
    size_t vertex_count = (size_t)atoi(argv[2]);
    const char *in_path = argv[3];
    const char *out_path = argv[4];

    FILE *fin = fopen(in_path, "rb");
    if (!fin) return 1;
    uint32_t *indices = (uint32_t *)malloc(index_count * sizeof(uint32_t));
    fread(indices, sizeof(uint32_t), index_count, fin);
    fclose(fin);

    uint32_t *dest = (uint32_t *)malloc(index_count * sizeof(uint32_t));
    meshopt_optimize_vertex_cache(dest, indices, index_count, vertex_count);

    FILE *fout = fopen(out_path, "wb");
    if (!fout) return 1;
    fwrite(dest, sizeof(uint32_t), index_count, fout);
    fclose(fout);

    free(indices);
    free(dest);
    return 0;
}
`, srcH, srcC)

	mainPath := filepath.Join(dir, "meshopt_oracle_main.c")
	if err := os.WriteFile(mainPath, []byte(mainC), 0600); err != nil {
		t.Fatal(err)
	}

	binPath := filepath.Join(dir, "meshopt_oracle_bin")
	cmdBuild := exec.Command("gcc", "-O2", "-Wall", "-Wextra", mainPath, "-o", binPath, "-lm")
	if out, err := cmdBuild.CombinedOutput(); err != nil {
		t.Fatalf("échec compilation oracle C: %v, sortie: %s", err, string(out))
	}

	rng := rand.New(rand.NewSource(9999))

	t.Run("VertexCache_BitExact_Vs_COracle", func(t *testing.T) {
		for iter := 0; iter < 10; iter++ {
			vertexCount := 60
			triCount := 150
			indexCount := triCount * 3

			indices := make([]uint32, indexCount)
			for i := 0; i < triCount; i++ {
				indices[i*3+0] = uint32(rng.Intn(vertexCount))
				indices[i*3+1] = uint32(rng.Intn(vertexCount))
				indices[i*3+2] = uint32(rng.Intn(vertexCount))
			}

			inPath := filepath.Join(dir, fmt.Sprintf("in_%d.bin", iter))
			outPath := filepath.Join(dir, fmt.Sprintf("out_%d.bin", iter))

			fIn, _ := os.Create(inPath)
			for _, v := range indices {
				_ = binary.Write(fIn, binary.LittleEndian, v)
			}
			fIn.Close()

			cmdRun := exec.Command(binPath, fmt.Sprintf("%d", indexCount), fmt.Sprintf("%d", vertexCount), inPath, outPath)
			if out, err := cmdRun.CombinedOutput(); err != nil {
				t.Fatalf("échec oracle meshopt: %v, out: %s", err, string(out))
			}

			cOutBytes, _ := os.ReadFile(outPath)
			cOutIndices := make([]uint32, indexCount)
			for i := 0; i < indexCount; i++ {
				cOutIndices[i] = binary.LittleEndian.Uint32(cOutBytes[i*4 : (i+1)*4])
			}

			goOut := make([]uint32, indexCount)
			OptimizeVertexCache(goOut, indices, vertexCount)

			if !reflect.DeepEqual(cOutIndices, goOut) {
				t.Fatalf("divergence bit-exacte à l'itération %d", iter)
			}

			acmrBefore := CalcVertexCacheStats(indices, vertexCount, 16)
			acmrAfter := CalcVertexCacheStats(goOut, vertexCount, 16)
			if acmrAfter > acmrBefore {
				t.Errorf("l'optimisation du cache doit réduire l'ACMR : avant %f, après %f", acmrBefore, acmrAfter)
			}
		}
	})
}
