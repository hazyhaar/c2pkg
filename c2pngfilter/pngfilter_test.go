// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2pngfilter

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestPngFilterVsCOracle(t *testing.T) {
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("gcc non disponible")
	}

	cSrcPath, err := filepath.Abs(filepath.Join("..", "..", "sources", "stbi_png_filter.c"))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(cSrcPath); err != nil {
		cSrcPath, err = filepath.Abs(filepath.Join("..", "..", "spec", "c_sources", "testdata", "c_sources", "stbi_png_filter.c"))
		if err != nil {
			t.Fatal(err)
		}
	}

	testHarness := `
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>

void stbi_unfilter_row(uint8_t *recon, const uint8_t *scanline, const uint8_t *prev, size_t len, int bpp, int filter_type);

int main(int argc, char **argv) {
    if (argc < 7) return 1;
    size_t len = (size_t)atoi(argv[1]);
    int bpp = atoi(argv[2]);
    int filter_type = atoi(argv[3]);

    FILE *fScan = fopen(argv[4], "rb");
    FILE *fPrev = fopen(argv[5], "rb");
    if (!fScan || !fPrev) return 2;

    uint8_t *scanline = (uint8_t*)malloc(len);
    uint8_t *prev = (uint8_t*)malloc(len);
    uint8_t *recon = (uint8_t*)malloc(len);

    if (fread(scanline, 1, len, fScan) != len || fread(prev, 1, len, fPrev) != len) {
        free(scanline); free(prev); free(recon);
        fclose(fScan); fclose(fPrev);
        return 3;
    }
    fclose(fScan); fclose(fPrev);

    stbi_unfilter_row(recon, scanline, prev, len, bpp, filter_type);

    FILE *out = fopen(argv[6], "wb");
    if (!out) return 4;
    fwrite(recon, 1, len, out);
    fclose(out);

    free(scanline); free(prev); free(recon);
    return 0;
}
`
	tmpDir := t.TempDir()
	harnessFile := filepath.Join(tmpDir, "harness.c")
	binFile := filepath.Join(tmpDir, "oracle_png")
	if err := os.WriteFile(harnessFile, []byte(testHarness), 0o644); err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("gcc", "-O2", "-o", binFile, harnessFile, cSrcPath)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("gcc compile failed: %v\n%s", err, string(out))
	}

	rng := rand.New(rand.NewSource(0x54621987))
	bppList := []int{1, 2, 3, 4}

	for filterType := 0; filterType <= 4; filterType++ {
		for _, bpp := range bppList {
			for iter := 0; iter < 10; iter++ {
				len_ := 16 + rng.Intn(1024)
				scanline := make([]byte, len_)
				prev := make([]byte, len_)
				rng.Read(scanline)
				rng.Read(prev)

				scanPath := filepath.Join(tmpDir, "scan.bin")
				prevPath := filepath.Join(tmpDir, "prev.bin")
				outPath := filepath.Join(tmpDir, "out.bin")

				if err := os.WriteFile(scanPath, scanline, 0o644); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(prevPath, prev, 0o644); err != nil {
					t.Fatal(err)
				}

				runCmd := exec.Command(binFile,
					fmt.Sprintf("%d", len_),
					fmt.Sprintf("%d", bpp),
					fmt.Sprintf("%d", filterType),
					scanPath, prevPath, outPath,
				)
				if out, err := runCmd.CombinedOutput(); err != nil {
					t.Fatalf("oracle failed filter=%d bpp=%d: %v\n%s", filterType, bpp, err, string(out))
				}

				cResult, err := os.ReadFile(outPath)
				if err != nil {
					t.Fatal(err)
				}

				goRecon := make([]byte, len_)
				Stbi_unfilter_row(goRecon, scanline, prev, uint64(len_), bpp, filterType)

				if !bytes.Equal(goRecon, cResult) {
					t.Fatalf("Divergence filter=%d bpp=%d len=%d iter=%d", filterType, bpp, len_, iter)
				}
			}
		}
	}
}

func TestPngFilterZeroAlloc(t *testing.T) {
	recon := make([]byte, 512)
	scan := make([]byte, 512)
	prev := make([]byte, 512)

	allocs := testing.AllocsPerRun(1000, func() {
		Stbi_unfilter_row(recon, scan, prev, 512, 4, 4)
	})
	if allocs != 0 {
		t.Fatalf("Allocs/op = %.2f, attendu 0", allocs)
	}
}
