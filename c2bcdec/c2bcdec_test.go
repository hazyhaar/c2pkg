// SPDX-License-Identifier: MIT
package c2bcdec

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestBCDecVsCOracle(t *testing.T) {
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("gcc non disponible")
	}

	srcC, err := filepath.Abs(filepath.Join("sources", "bcdec.c"))
	if err != nil {
		t.Fatal(err)
	}
	srcH, err := filepath.Abs(filepath.Join("sources", "bcdec.h"))
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
    int mode = atoi(argv[1]);
    const char *in_path = argv[2];
    const char *out_path = argv[3];

    FILE *fin = fopen(in_path, "rb");
    if (!fin) return 1;
    uint8_t block[16];
    size_t n = fread(block, 1, 16, fin);
    fclose(fin);

    uint8_t out[64];
    if (mode == 1) {
        bcdec_bc1(block, out, 16);
    } else if (mode == 3) {
        bcdec_bc3(block, out, 16);
    } else if (mode == 4) {
        bcdec_bc4(block, out, 4);
    } else if (mode == 5) {
        bcdec_bc5(block, out, 8);
    }

    FILE *fout = fopen(out_path, "wb");
    if (!fout) return 1;
    if (mode == 1 || mode == 3) {
        fwrite(out, 1, 64, fout);
    } else if (mode == 4) {
        fwrite(out, 1, 16, fout);
    } else if (mode == 5) {
        fwrite(out, 1, 32, fout);
    }
    fclose(fout);
    return 0;
}
`, srcH, srcC)

	mainPath := filepath.Join(dir, "bcdec_oracle_main.c")
	if err := os.WriteFile(mainPath, []byte(mainC), 0600); err != nil {
		t.Fatal(err)
	}

	binPath := filepath.Join(dir, "bcdec_oracle_bin")
	cmdBuild := exec.Command("gcc", "-O2", "-Wall", "-Wextra", mainPath, "-o", binPath)
	if out, err := cmdBuild.CombinedOutput(); err != nil {
		t.Fatalf("échec compilation oracle C: %v, sortie: %s", err, string(out))
	}

	rng := rand.New(rand.NewSource(12345))

	t.Run("BC1_BitExact_Vs_COracle", func(t *testing.T) {
		for iter := 0; iter < 50; iter++ {
			block := make([]byte, 8)
			rng.Read(block)

			inPath := filepath.Join(dir, "in_bc1.bin")
			outPath := filepath.Join(dir, "out_bc1.bin")
			_ = os.WriteFile(inPath, block, 0600)

			cmdRun := exec.Command(binPath, "1", inPath, outPath)
			if out, err := cmdRun.CombinedOutput(); err != nil {
				t.Fatalf("échec oracle BC1: %v, out: %s", err, string(out))
			}

			cOut, _ := os.ReadFile(outPath)
			goOut := make([]byte, 64)
			DecodeBC1Block(block, goOut, 16)

			if !bytes.Equal(cOut, goOut) {
				t.Fatalf("divergence bit-exacte BC1 à l'itération %d", iter)
			}
		}
	})

	t.Run("BC3_BitExact_Vs_COracle", func(t *testing.T) {
		for iter := 0; iter < 50; iter++ {
			block := make([]byte, 16)
			rng.Read(block)

			inPath := filepath.Join(dir, "in_bc3.bin")
			outPath := filepath.Join(dir, "out_bc3.bin")
			_ = os.WriteFile(inPath, block, 0600)

			cmdRun := exec.Command(binPath, "3", inPath, outPath)
			if out, err := cmdRun.CombinedOutput(); err != nil {
				t.Fatalf("échec oracle BC3: %v, out: %s", err, string(out))
			}

			cOut, _ := os.ReadFile(outPath)
			goOut := make([]byte, 64)
			DecodeBC3Block(block, goOut, 16)

			if !bytes.Equal(cOut, goOut) {
				t.Fatalf("divergence bit-exacte BC3 à l'itération %d", iter)
			}
		}
	})

	t.Run("BC4_BitExact_Vs_COracle", func(t *testing.T) {
		for iter := 0; iter < 50; iter++ {
			block := make([]byte, 8)
			rng.Read(block)

			inPath := filepath.Join(dir, "in_bc4.bin")
			outPath := filepath.Join(dir, "out_bc4.bin")
			_ = os.WriteFile(inPath, block, 0600)

			cmdRun := exec.Command(binPath, "4", inPath, outPath)
			if out, err := cmdRun.CombinedOutput(); err != nil {
				t.Fatalf("échec oracle BC4: %v, out: %s", err, string(out))
			}

			cOut, _ := os.ReadFile(outPath)
			goOut := make([]byte, 16)
			DecodeBC4Block(block, goOut, 4)

			if !bytes.Equal(cOut, goOut) {
				t.Fatalf("divergence bit-exacte BC4 à l'itération %d", iter)
			}
		}
	})

	t.Run("BC5_BitExact_Vs_COracle", func(t *testing.T) {
		for iter := 0; iter < 50; iter++ {
			block := make([]byte, 16)
			rng.Read(block)

			inPath := filepath.Join(dir, "in_bc5.bin")
			outPath := filepath.Join(dir, "out_bc5.bin")
			_ = os.WriteFile(inPath, block, 0600)

			cmdRun := exec.Command(binPath, "5", inPath, outPath)
			if out, err := cmdRun.CombinedOutput(); err != nil {
				t.Fatalf("échec oracle BC5: %v, out: %s", err, string(out))
			}

			cOut, _ := os.ReadFile(outPath)
			goOut := make([]byte, 32)
			DecodeBC5Block(block, goOut, 8)

			if !bytes.Equal(cOut, goOut) {
				t.Fatalf("divergence bit-exacte BC5 à l'itération %d", iter)
			}
		}
	})
}

func TestBCDecZeroAlloc(t *testing.T) {
	block := make([]byte, 16)
	dst := make([]byte, 64)

	allocs := testing.AllocsPerRun(100, func() {
		DecodeBC3Block(block, dst, 16)
	})

	if allocs != 0 {
		t.Errorf("échec invariant ARCHTIME : attendu 0 allocs/op, obtenu %f", allocs)
	}
}
