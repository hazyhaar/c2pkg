// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2pcre

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestRegexSimpleMatches(t *testing.T) {
	tests := []struct {
		pattern string
		text    string
		want    bool
	}{
		{"hello", "hello world", true},
		{"^hello", "hello world", true},
		{"^world", "hello world", false},
		{"world$", "hello world", true},
		{"world$", "hello world!", false},
		{"a.*b", "axxxb", true},
		{"a.+b", "ab", false},
		{"a.+b", "axb", true},
		{"a?b", "b", true},
		{"a?b", "ab", true},
		{"\\d+", "id_12345_token", true},
		{"^\\d+$", "12345", true},
		{"^\\d+$", "123a45", false},
	}

	for _, tc := range tests {
		var re C2_regex_t
		nOps := C2_pcre_compile([]byte(tc.pattern), uint32(len(tc.pattern)), &re)
		if nOps <= 0 {
			t.Fatalf("compile failed for %q", tc.pattern)
		}

		var m C2_match_t
		matched := C2_pcre_exec(&re, []byte(tc.text), uint32(len(tc.text)), &m) == 1
		if matched != tc.want {
			t.Fatalf("pattern %q in text %q: got %v, want %v", tc.pattern, tc.text, matched, tc.want)
		}
	}
}

func TestRegexVsCOracle(t *testing.T) {
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("gcc non disponible")
	}

	cSrcPath, err := filepath.Abs(filepath.Join("..", "..", "sources", "c2pcre", "pcre.c"))
	if err != nil || cSrcPath == "" {
		t.Skip("source C pcre.c introuvable")
	}

	testHarness := `
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>
#include <string.h>

#define MAX_OPCODES 128

typedef struct c2_regex_op_s {
    uint8_t opcode;
    uint8_t ch;
    uint8_t target_op;
} c2_regex_op_t;

typedef struct c2_regex_s {
    c2_regex_op_t ops[MAX_OPCODES];
    uint32_t      n_ops;
} c2_regex_t;

typedef struct c2_match_s {
    int32_t start;
    int32_t end;
} c2_match_t;

int32_t c2_pcre_compile(const char *pattern, uint32_t pat_len, c2_regex_t *re);
int32_t c2_pcre_exec(const c2_regex_t *re, const char *text, uint32_t text_len, c2_match_t *match);

int main(int argc, char **argv) {
    if (argc < 3) return 1;
    const char *pat = argv[1];
    const char *txt = argv[2];

    c2_regex_t re;
    c2_pcre_compile(pat, strlen(pat), &re);

    c2_match_t m = {-1, -1};
    int32_t res = c2_pcre_exec(&re, txt, strlen(txt), &m);

    printf("RES=%d START=%d END=%d\n", res, m.start, m.end);
    return 0;
}
`
	tmpDir := t.TempDir()
	harnessFile := filepath.Join(tmpDir, "harness.c")
	binFile := filepath.Join(tmpDir, "oracle_pcre")
	if err := os.WriteFile(harnessFile, []byte(testHarness), 0o644); err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("gcc", "-O2", "-o", binFile, harnessFile, cSrcPath)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("gcc compile failed: %v\n%s", err, string(out))
	}

	testCases := [][2]string{
		{"foo", "foobar"},
		{"^bar", "foobar"},
		{"bar$", "foobar"},
		{".*test", "this is a test"},
		{"\\d+", "number 42 in text"},
		{"PHP_127", "PHP_127 Sovereign Stack"},
		{"a+b", "aaaaab"},
	}

	for _, tc := range testCases {
		pat, txt := tc[0], tc[1]
		runCmd := exec.Command(binFile, pat, txt)
		out, err := runCmd.CombinedOutput()
		if err != nil {
			t.Fatalf("oracle execution failed: %v\n%s", err, string(out))
		}

		var re C2_regex_t
		C2_pcre_compile([]byte(pat), uint32(len(pat)), &re)
		var m C2_match_t
		res := C2_pcre_exec(&re, []byte(txt), uint32(len(txt)), &m)

		expectedLine := fmt.Sprintf("RES=%d START=%d END=%d", res, m.Start, m.End)
		cOutStr := strings.TrimSpace(string(out))
		if cOutStr != expectedLine {
			t.Fatalf("Divergence C vs Go pour pattern %q dans %q: Oracle_C=%q, Go=%q",
				pat, txt, cOutStr, expectedLine)
		}
	}
}

func TestRegexZeroAlloc(t *testing.T) {
	var re C2_regex_t
	pat := []byte("^\\d+_token$")
	txt := []byte("123456_token")
	C2_pcre_compile(pat, uint32(len(pat)), &re)

	var m C2_match_t
	allocs := testing.AllocsPerRun(1000, func() {
		_ = C2_pcre_exec(&re, txt, uint32(len(txt)), &m)
	})
	if allocs != 0 {
		t.Fatalf("Allocs/op = %.2f, attendu 0", allocs)
	}
}
