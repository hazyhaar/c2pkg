// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2mem

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestArenaAllocAndReset(t *testing.T) {
	var arena C2_arena_t
	buf := make([]byte, 1024)
	cleanups := make([]C2_cleanup_entry_t, 16)

	C2_arena_init(&arena, buf, 1024, cleanups, 16)

	off1 := C2_arena_alloc(&arena, 10)
	if off1 != 0 {
		t.Fatalf("off1 attendu 0, got %d", off1)
	}

	off2 := C2_arena_alloc(&arena, 24)
	if off2 != 16 { // aligné à 16 (10 -> 16)
		t.Fatalf("off2 attendu 16, got %d", off2)
	}

	ok := C2_arena_register_cleanup(&arena, 42, 1, 999)
	if !ok || arena.N_cleanups != 1 {
		t.Fatalf("register cleanup fail: ok=%v n=%d", ok, arena.N_cleanups)
	}

	count := C2_arena_reset(&arena)
	if count != 1 || arena.Offset != 0 || arena.N_cleanups != 0 {
		t.Fatalf("reset fail: count=%d off=%d n=%d", count, arena.Offset, arena.N_cleanups)
	}
}

func TestArenaVsCOracle(t *testing.T) {
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("gcc non disponible")
	}

	cSrcPath, err := filepath.Abs(filepath.Join("..", "..", "sources", "c2mem", "mem.c"))
	if err != nil || cSrcPath == "" {
		t.Skip("source C mem.c introuvable")
	}

	testHarness := `
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

typedef struct c2_cleanup_entry_s {
    uint32_t resource_id;
    uint32_t flags;
    int64_t  param;
} c2_cleanup_entry_t;

typedef struct c2_arena_s {
    uint8_t            *buffer;
    uint32_t            capacity;
    uint32_t            offset;
    uint32_t            peak_offset;
    uint32_t            n_cleanups;
    uint32_t            max_cleanups;
    c2_cleanup_entry_t *cleanups;
} c2_arena_t;

void c2_arena_init(c2_arena_t *arena, uint8_t *buffer, uint32_t capacity, c2_cleanup_entry_t *cleanup_buf, uint32_t max_cleanups);
int32_t c2_arena_alloc(c2_arena_t *arena, uint32_t size);
bool c2_arena_register_cleanup(c2_arena_t *arena, uint32_t resource_id, uint32_t flags, int64_t param);
uint32_t c2_arena_reset(c2_arena_t *arena);

int main(void) {
    c2_arena_t arena;
    uint8_t buf[2048];
    c2_cleanup_entry_t cleanups[32];

    c2_arena_init(&arena, buf, 2048, cleanups, 32);

    int32_t o1 = c2_arena_alloc(&arena, 13);
    int32_t o2 = c2_arena_alloc(&arena, 7);
    int32_t o3 = c2_arena_alloc(&arena, 64);
    c2_arena_register_cleanup(&arena, 101, 0, 12345);
    c2_arena_register_cleanup(&arena, 102, 1, 67890);

    printf("OFF1=%d OFF2=%d OFF3=%d CLEANUPS=%u PEAK=%u\n",
        o1, o2, o3, arena.n_cleanups, arena.peak_offset);

    uint32_t reset_count = c2_arena_reset(&arena);
    printf("RESET_COUNT=%u OFF_AFTER=%u\n", reset_count, arena.offset);
    return 0;
}
`
	tmpDir := t.TempDir()
	harnessFile := filepath.Join(tmpDir, "harness.c")
	binFile := filepath.Join(tmpDir, "oracle_mem")
	if err := os.WriteFile(harnessFile, []byte(testHarness), 0o644); err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("gcc", "-O2", "-o", binFile, harnessFile, cSrcPath)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("gcc compile failed: %v\n%s", err, string(out))
	}

	runCmd := exec.Command(binFile)
	out, err := runCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("oracle execution failed: %v\n%s", err, string(out))
	}

	// Go arena simulation
	var arena C2_arena_t
	buf := make([]byte, 2048)
	cleanups := make([]C2_cleanup_entry_t, 32)
	C2_arena_init(&arena, buf, 2048, cleanups, 32)

	o1 := C2_arena_alloc(&arena, 13)
	o2 := C2_arena_alloc(&arena, 7)
	o3 := C2_arena_alloc(&arena, 64)
	C2_arena_register_cleanup(&arena, 101, 0, 12345)
	C2_arena_register_cleanup(&arena, 102, 1, 67890)

	expectedLine1 := fmt.Sprintf("OFF1=%d OFF2=%d OFF3=%d CLEANUPS=%d PEAK=%d",
		o1, o2, o3, arena.N_cleanups, arena.Peak_offset)

	resetCount := C2_arena_reset(&arena)
	expectedLine2 := fmt.Sprintf("RESET_COUNT=%d OFF_AFTER=%d", resetCount, arena.Offset)

	cOutput := string(out)
	if !strings.Contains(cOutput, expectedLine1) {
		t.Fatalf("Divergence C vs Go: ligne 1 attendue %q non trouvée\n%s", expectedLine1, cOutput)
	}
	if !strings.Contains(cOutput, expectedLine2) {
		t.Fatalf("Divergence C vs Go: ligne 2 attendue %q non trouvée\n%s", expectedLine2, cOutput)
	}
}

func TestArenaZeroAlloc(t *testing.T) {
	var arena C2_arena_t
	buf := make([]byte, 4096)
	cleanups := make([]C2_cleanup_entry_t, 64)
	C2_arena_init(&arena, buf, 4096, cleanups, 64)

	allocs := testing.AllocsPerRun(1000, func() {
		_ = C2_arena_alloc(&arena, 32)
		_ = C2_arena_register_cleanup(&arena, 1, 0, 100)
		_ = C2_arena_reset(&arena)
	})
	if allocs != 0 {
		t.Fatalf("Allocs/op = %.2f, attendu 0", allocs)
	}
}
