// SPDX-License-Identifier: MIT
package c2gltf

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestGLTFVsCOracle(t *testing.T) {
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("gcc non disponible")
	}

	srcC, err := filepath.Abs(filepath.Join("sources", "cgltf.c"))
	if err != nil {
		t.Fatal(err)
	}
	srcH, err := filepath.Abs(filepath.Join("sources", "cgltf.h"))
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
    if (argc < 7) return 1;
    int comp_type = atoi(argv[1]);
    int type = atoi(argv[2]);
    int count = atoi(argv[3]);
    int stride = atoi(argv[4]);
    const char *in_path = argv[5];
    const char *out_path = argv[6];

    FILE *fin = fopen(in_path, "rb");
    if (!fin) return 1;
    fseek(fin, 0, SEEK_END);
    long sz = ftell(fin);
    fseek(fin, 0, SEEK_SET);
    uint8_t *src = (uint8_t *)malloc(sz);
    fread(src, 1, sz, fin);
    fclose(fin);

    size_t num_floats = count * 4; // Max comps per elem
    float *out_floats = (float *)malloc(num_floats * sizeof(float));
    size_t written = cgltf_accessor_read_float((cgltf_component_type)comp_type, (cgltf_type)type, src, count, stride, out_floats);

    FILE *fout = fopen(out_path, "wb");
    if (!fout) return 1;
    fwrite(out_floats, sizeof(float), written, fout);
    fclose(fout);

    free(src);
    free(out_floats);
    return 0;
}
`, srcH, srcC)

	mainPath := filepath.Join(dir, "cgltf_oracle_main.c")
	if err := os.WriteFile(mainPath, []byte(mainC), 0600); err != nil {
		t.Fatal(err)
	}

	binPath := filepath.Join(dir, "cgltf_oracle_bin")
	cmdBuild := exec.Command("gcc", "-O2", "-Wall", "-Wextra", mainPath, "-o", binPath)
	if out, err := cmdBuild.CombinedOutput(); err != nil {
		t.Fatalf("échec compilation oracle C: %v, sortie: %s", err, string(out))
	}

	rng := rand.New(rand.NewSource(54321))

	t.Run("AccessorReadFloat_BitExact_Vs_COracle", func(t *testing.T) {
		tests := []struct {
			compType ComponentType
			tType    Type
			count    int
			stride   int
		}{
			{ComponentTypeR32F, TypeVec3, 50, 12},
			{ComponentTypeR32F, TypeVec3, 50, 16}, // Interleaved stride
			{ComponentTypeR16U, TypeVec2, 40, 4},
			{ComponentTypeR8U, TypeVec4, 30, 4},
		}

		for _, tc := range tests {
			t.Run(fmt.Sprintf("Comp%d_Type%d", tc.compType, tc.tType), func(t *testing.T) {
				bufSize := tc.count * tc.stride
				rawBuf := make([]byte, bufSize)
				rng.Read(rawBuf)

				inPath := filepath.Join(dir, "in_acc.bin")
				outPath := filepath.Join(dir, "out_acc.bin")
				_ = os.WriteFile(inPath, rawBuf, 0600)

				cmdRun := exec.Command(binPath, fmt.Sprintf("%d", tc.compType), fmt.Sprintf("%d", tc.tType), fmt.Sprintf("%d", tc.count), fmt.Sprintf("%d", tc.stride), inPath, outPath)
				if out, err := cmdRun.CombinedOutput(); err != nil {
					t.Fatalf("échec oracle cgltf: %v, out: %s", err, string(out))
				}

				cOutBytes, _ := os.ReadFile(outPath)
				numFloats := len(cOutBytes) / 4
				cFloats := make([]float32, numFloats)
				for i := 0; i < numFloats; i++ {
					u := binary.LittleEndian.Uint32(cOutBytes[i*4 : (i+1)*4])
					cFloats[i] = math.Float32frombits(u)
				}

				goFloats := make([]float32, numFloats)
				written := AccessorReadFloat(tc.compType, tc.tType, rawBuf, tc.count, tc.stride, goFloats)

				if written != numFloats {
					t.Fatalf("quantité de flottants divergente : attendu %d, obtenu %d", numFloats, written)
				}

				for i := 0; i < numFloats; i++ {
					if math.Float32bits(cFloats[i]) != math.Float32bits(goFloats[i]) {
						t.Fatalf("divergence bit-exacte à l'index %d : C=%08X, Go=%08X", i, math.Float32bits(cFloats[i]), math.Float32bits(goFloats[i]))
					}
				}
			})
		}
	})

	t.Run("ParseGLBHeader_Synthetic", func(t *testing.T) {
		jsonContent := []byte(`{"asset":{"version":"2.0"}}     `) // 32 octets exacts (multiple de 4)
		jsonLen := uint32(len(jsonContent))
		binContent := []byte{0x01, 0x02, 0x03, 0x04} // 4 octets exacts
		binLen := uint32(len(binContent))

		totalLen := uint32(12 + 8 + jsonLen + 8 + binLen)
		glbBuf := new(bytes.Buffer)
		_ = binary.Write(glbBuf, binary.LittleEndian, uint32(MagicGLB))
		_ = binary.Write(glbBuf, binary.LittleEndian, uint32(2))
		_ = binary.Write(glbBuf, binary.LittleEndian, totalLen)

		_ = binary.Write(glbBuf, binary.LittleEndian, jsonLen)
		_ = binary.Write(glbBuf, binary.LittleEndian, uint32(ChunkJSON))
		glbBuf.Write(jsonContent)

		_ = binary.Write(glbBuf, binary.LittleEndian, binLen)
		_ = binary.Write(glbBuf, binary.LittleEndian, uint32(ChunkBIN))
		glbBuf.Write(binContent)

		hdr, err := ParseGLBHeader(glbBuf.Bytes())
		if err != nil {
			t.Fatalf("erreur de parsing GLB valide: %v", err)
		}

		if string(hdr.JSONData) != string(jsonContent) {
			t.Errorf("divergence JSONData: %s", string(hdr.JSONData))
		}
		if !bytes.Equal(hdr.BINData, binContent) {
			t.Errorf("divergence BINData: %v", hdr.BINData)
		}
	})
}
