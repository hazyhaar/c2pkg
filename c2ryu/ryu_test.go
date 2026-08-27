// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2ryu

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestU64ToStringParity(t *testing.T) {
	nums := []uint64{0, 1, 9, 10, 42, 100, 999, 123456789, math.MaxUint32, math.MaxUint64}
	buf := make([]byte, 32)
	for _, n := range nums {
		l := C2_u64_to_str(n, buf)
		got := string(buf[:l])
		want := strconv.FormatUint(n, 10)
		if got != want {
			t.Fatalf("u64 %d: got %q, want %q", n, got, want)
		}
	}
}

func TestI64ToStringParity(t *testing.T) {
	nums := []int64{0, 1, -1, 42, -42, math.MaxInt32, math.MinInt32, math.MaxInt64, math.MinInt64}
	buf := make([]byte, 32)
	for _, n := range nums {
		l := C2_i64_to_str(n, buf)
		got := string(buf[:l])
		want := strconv.FormatInt(n, 10)
		if got != want {
			t.Fatalf("i64 %d: got %q, want %q", n, got, want)
		}
	}
}

func TestRyuVsCOracle(t *testing.T) {
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("gcc non disponible")
	}

	cSrcPath, err := filepath.Abs(filepath.Join("..", "..", "sources", "c2ryu", "ryu.c"))
	if err != nil || cSrcPath == "" {
		t.Skip("source C ryu.c introuvable")
	}

	testHarness := `
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>

uint32_t c2_u64_to_str(uint64_t v, char *dst);
uint32_t c2_i64_to_str(int64_t v, char *dst);
uint32_t c2_d2s_buffered(double f, char *dst);

int main(int argc, char **argv) {
    if (argc < 2) return 1;
    double val = atof(argv[1]);
    char buf[64];
    uint32_t len = c2_d2s_buffered(val, buf);
    printf("%s\n", buf);
    return 0;
}
`
	tmpDir := t.TempDir()
	harnessFile := filepath.Join(tmpDir, "harness.c")
	binFile := filepath.Join(tmpDir, "oracle_ryu")
	if err := os.WriteFile(harnessFile, []byte(testHarness), 0o644); err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("gcc", "-O2", "-o", binFile, harnessFile, cSrcPath)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("gcc compile failed: %v\n%s", err, string(out))
	}

	rng := rand.New(rand.NewSource(0xCAFEBABE))
	testFloats := []float64{0.0, 1.0, -1.0, 3.14159265, 42.0, -123.456, 1000000.5}
	for i := 0; i < 50; i++ {
		testFloats = append(testFloats, (rng.Float64()-0.5)*100000.0)
	}

	buf := make([]byte, 64)
	for _, f := range testFloats {
		argStr := fmt.Sprintf("%.8f", f)
		runCmd := exec.Command(binFile, argStr)
		out, err := runCmd.CombinedOutput()
		if err != nil {
			t.Fatalf("oracle execution failed: %v\n%s", err, string(out))
		}

		cOut := strings.TrimSpace(string(out))
		l := C2_d2s_buffered(f, buf)
		goOut := string(buf[:l])

		// Double conversion comparison within float64 tolerance
		cVal, err1 := strconv.ParseFloat(cOut, 64)
		goVal, err2 := strconv.ParseFloat(goOut, 64)
		if err1 == nil && err2 == nil {
			diff := math.Abs(cVal - goVal)
			if diff > 1e-4 {
				t.Fatalf("float divergence for %f: go=%q (%f) c=%q (%f)", f, goOut, goVal, cOut, cVal)
			}
		}
	}
}

func TestRyuZeroAlloc(t *testing.T) {
	buf := make([]byte, 64)
	allocs := testing.AllocsPerRun(1000, func() {
		_ = C2_u64_to_str(123456789, buf)
		_ = C2_i64_to_str(-987654321, buf)
		_ = C2_d2s_buffered(3.1415926535, buf)
	})
	if allocs != 0 {
		t.Fatalf("Allocs/op = %.2f, attendu 0", allocs)
	}
}
