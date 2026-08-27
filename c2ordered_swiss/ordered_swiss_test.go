// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2ordered_swiss

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestSwissInitPacked(t *testing.T) {
	var ht Swiss_table_t
	ctrl := make([]byte, 64)
	slots := make([]uint32, 64)
	entries := make([]Swiss_element_t, 32)

	Swiss_init(&ht, ctrl, slots, entries, 32, 32)
	if ht.N_elements != 0 || ht.N_used != 0 {
		t.Fatalf("table non vide apres init: elems=%d used=%d", ht.N_elements, ht.N_used)
	}

	// Insert packed elements
	for i := int64(0); i < 10; i++ {
		ok := Swiss_insert_packed_int(&ht, i, 1, i*100)
		if !ok {
			t.Fatalf("echec insert packed pour i=%d", i)
		}
	}

	if ht.N_elements != 10 || ht.N_used != 10 {
		t.Fatalf("apres inserts: elems=%d used=%d", ht.N_elements, ht.N_used)
	}

	// Find packed elements
	for i := int64(0); i < 10; i++ {
		idx := Swiss_find_int(&ht, i)
		if idx != int(i) {
			t.Fatalf("find_int(%d): got %d, want %d", i, idx, i)
		}
		if ht.Entries[idx].Val.I64 != i*100 {
			t.Fatalf("val.i64 pour %d: got %d, want %d", i, ht.Entries[idx].Val.I64, i*100)
		}
	}

	// Find non-existent
	if idx := Swiss_find_int(&ht, 999); idx != -1 {
		t.Fatalf("find_int(999) attendu -1, got %d", idx)
	}
}

func TestSwissPopAndCompaction(t *testing.T) {
	var ht Swiss_table_t
	ctrl := make([]byte, 64)
	slots := make([]uint32, 64)
	entries := make([]Swiss_element_t, 32)

	Swiss_init(&ht, ctrl, slots, entries, 32, 32)
	for i := int64(0); i < 5; i++ {
		Swiss_insert_packed_int(&ht, i, 1, i+1)
	}

	ok := Swiss_pop(&ht)
	if !ok || ht.N_elements != 4 || ht.N_used != 4 {
		t.Fatalf("pop failed: ok=%v elems=%d used=%d", ok, ht.N_elements, ht.N_used)
	}

	ht.N_tombstones = 1
	ht.Entries[1].Val.Type_ = 0 // Simule suppression
	Swiss_compact(&ht)
	if ht.N_tombstones != 0 {
		t.Fatalf("tombstones apres compact: %d", ht.N_tombstones)
	}
}

func TestSwissVsCOracle(t *testing.T) {
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("gcc non disponible")
	}

	cSrcPath, err := filepath.Abs(filepath.Join("..", "..", "sources", "c2ordered_swiss", "ordered_swiss.c"))
	if err != nil || cSrcPath == "" {
		t.Skip("source C ordered_swiss.c introuvable")
	}

	testHarness := `
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

typedef struct swiss_val_s {
    uint32_t type;
    uint32_t flags;
    int64_t  i64;
} swiss_val_t;

typedef struct swiss_element_s {
    swiss_val_t val;
    uint64_t    hash;
    char       *str_key;
    uint32_t    str_len;
    int64_t     int_key;
} swiss_element_t;

typedef struct swiss_table_s {
    uint8_t         *ctrl;
    uint32_t        *ctrl_to_dense;
    swiss_element_t *entries;
    uint32_t         capacity_hash;
    uint32_t         capacity_data;
    uint32_t         n_elements;
    uint32_t         n_used;
    uint32_t         n_tombstones;
    uint32_t         flags;
} swiss_table_t;

void swiss_init(swiss_table_t *ht, uint8_t *ctrl_buf, uint32_t *slots_buf, swiss_element_t *entries_buf, uint32_t cap_hash, uint32_t cap_data);
bool swiss_insert_packed_int(swiss_table_t *ht, int64_t key, uint32_t val_type, int64_t val_i64);
int32_t swiss_find_int(const swiss_table_t *ht, int64_t key);
bool swiss_pop(swiss_table_t *ht);

int main(int argc, char **argv) {
    swiss_table_t ht;
    uint8_t ctrl[128];
    uint32_t slots[128];
    swiss_element_t entries[64];

    swiss_init(&ht, ctrl, slots, entries, 64, 64);
    for (int64_t i = 0; i < 20; ++i) {
        swiss_insert_packed_int(&ht, i, 1, i * 7 + 3);
    }

    for (int64_t i = 0; i < 20; ++i) {
        int32_t idx = swiss_find_int(&ht, i);
        printf("IDX_%lld=%d VAL_%lld=%lld\n", (long long)i, idx, (long long)i, (long long)ht.entries[idx].val.i64);
    }
    return 0;
}
`
	tmpDir := t.TempDir()
	harnessFile := filepath.Join(tmpDir, "harness.c")
	binFile := filepath.Join(tmpDir, "oracle_swiss")
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

	// Go table verification
	var ht Swiss_table_t
	ctrl := make([]byte, 128)
	slots := make([]uint32, 128)
	entries := make([]Swiss_element_t, 64)
	Swiss_init(&ht, ctrl, slots, entries, 64, 64)

	for i := int64(0); i < 20; i++ {
		Swiss_insert_packed_int(&ht, i, 1, i*7+3)
	}

	cOutStr := string(out)
	for i := int64(0); i < 20; i++ {
		idx := Swiss_find_int(&ht, i)
		val := ht.Entries[idx].Val.I64
		expectedLine := fmt.Sprintf("IDX_%d=%d VAL_%d=%d", i, idx, i, val)
		if !containsExactLine(cOutStr, expectedLine) {
			t.Fatalf("Divergence C vs Go: ligne attendue %q absente de l'oracle C\n%s", expectedLine, cOutStr)
		}
	}
}

func containsExactLine(text, expectedLine string) bool {
	lines := strings.Split(text, "\n")
	for _, l := range lines {
		if strings.TrimSpace(l) == strings.TrimSpace(expectedLine) {
			return true
		}
	}
	return false
}

func TestSwissZeroAlloc(t *testing.T) {
	var ht Swiss_table_t
	ctrl := make([]byte, 128)
	slots := make([]uint32, 128)
	entries := make([]Swiss_element_t, 64)
	Swiss_init(&ht, ctrl, slots, entries, 64, 64)

	for i := int64(0); i < 20; i++ {
		Swiss_insert_packed_int(&ht, i, 1, i*10)
	}

	allocs := testing.AllocsPerRun(1000, func() {
		_ = Swiss_find_int(&ht, 5)
	})
	if allocs != 0 {
		t.Fatalf("Allocs/op = %.2f, attendu 0", allocs)
	}
}
