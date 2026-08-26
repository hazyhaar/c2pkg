// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2grid

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"unsafe"
)

func TestC2GridVsCOracle(t *testing.T) {
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("gcc non disponible")
	}

	cSrc := `
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>
#include "c2_grid.h"

int main(int argc, char **argv) {
    if (argc < 3) return 1;
    FILE *f = fopen(argv[1], "rb");
    if (!f) return 2;

    int width = 80, height = 24, stride = 80;
    int ncells = width * height;
    c2_cell_t *cells = (c2_cell_t*)malloc(ncells * sizeof(c2_cell_t));
    if (fread(cells, sizeof(c2_cell_t), ncells, f) != (size_t)ncells) {
        free(cells);
        fclose(f);
        return 3;
    }
    fclose(f);

    // 1. Scroll up 3 lignes
    c2_grid_scroll_up(cells, width, height, stride, 3);

    // 2. Clear row 10
    c2_grid_clear_row(cells, 10, stride, width, 5, 2);

    // 3. Clear bottom 2 rows
    c2_grid_clear(&cells[22 * stride], 2 * stride, 1, 4);

    FILE *out = fopen(argv[2], "wb");
    if (!out) {
        free(cells);
        return 4;
    }
    fwrite(cells, sizeof(c2_cell_t), ncells, out);
    fclose(out);
    free(cells);
    return 0;
}
`
	tmpDir := t.TempDir()
	cFile := filepath.Join(tmpDir, "oracle_grid.c")
	binFile := filepath.Join(tmpDir, "oracle_grid")
	if err := os.WriteFile(cFile, []byte(cSrc), 0o644); err != nil {
		t.Fatal(err)
	}

	incDir, err := filepath.Abs(filepath.Join("..", "..", "sources"))
	if err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("gcc", "-O2", "-I", incDir, "-o", binFile, cFile)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("gcc compile failed: %v\n%s", err, string(out))
	}

	// Préparer un jeu de cellules pseudo-aléatoire
	const width = 80
	const height = 24
	const ncells = width * height

	rng := rand.New(rand.NewSource(0x1337BEEF))
	cellsGo := make([]C2_cell_t, ncells)
	cellsBytes := make([]byte, ncells*8)

	for i := 0; i < ncells; i++ {
		cellsGo[i] = C2_cell_t{
			Rune_: uint32('A' + (rng.Intn(26))),
			Fg:    uint8(rng.Intn(256)),
			Bg:    uint8(rng.Intn(256)),
			Flags: uint8(rng.Intn(16)),
			Width: 1,
		}
		offset := i * 8
		binary.LittleEndian.PutUint32(cellsBytes[offset:offset+4], cellsGo[i].Rune_)
		cellsBytes[offset+4] = cellsGo[i].Fg
		cellsBytes[offset+5] = cellsGo[i].Bg
		cellsBytes[offset+6] = cellsGo[i].Flags
		cellsBytes[offset+7] = cellsGo[i].Width
	}

	inPath := filepath.Join(tmpDir, "in.bin")
	outPath := filepath.Join(tmpDir, "out.bin")
	if err := os.WriteFile(inPath, cellsBytes, 0o644); err != nil {
		t.Fatal(err)
	}

	// Exécuter l'oracle C
	runCmd := exec.Command(binFile, inPath, outPath)
	if out, err := runCmd.CombinedOutput(); err != nil {
		t.Fatalf("oracle C execution failed: %v\n%s", err, string(out))
	}

	cResultBytes, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatal(err)
	}

	// Exécuter exactement les mêmes opérations en Go transpilé
	C2_grid_scroll_up(cellsGo, width, height, width, 3)
	C2_grid_clear_row(cellsGo, 10, width, width, 5, 2)
	C2_grid_clear(cellsGo[22*width:], 2*width, 1, 4)

	// Comparer bit à bit les deux résultats
	for i := 0; i < ncells; i++ {
		offset := i * 8
		cRune := binary.LittleEndian.Uint32(cResultBytes[offset : offset+4])
		cfg := cResultBytes[offset+4]
		cbg := cResultBytes[offset+5]
		cflags := cResultBytes[offset+6]
		cwidth := cResultBytes[offset+7]

		g := cellsGo[i]
		if g.Rune_ != cRune || g.Fg != cfg || g.Bg != cbg || g.Flags != cflags || g.Width != cwidth {
			t.Fatalf("Divergence à l'indice %d: Go(Rune=%c, Fg=%d, Bg=%d, Flags=%d, W=%d) vs C(Rune=%c, Fg=%d, Bg=%d, Flags=%d, W=%d)",
				i, g.Rune_, g.Fg, g.Bg, g.Flags, g.Width, cRune, cfg, cbg, cflags, cwidth)
		}
	}
	_ = unsafe.Sizeof(cellsGo[0])
	fmt.Printf("Parité bit-exacte c2grid validée : %d cellules conformes à l'oracle GCC -O2\n", ncells)
}
