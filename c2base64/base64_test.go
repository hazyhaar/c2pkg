// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2base64

import (
	"bytes"
	"encoding/base64"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestBase64StdLibParity(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	for length := 0; length < 1000; length++ {
		data := make([]byte, length)
		rng.Read(data)

		dst := make([]byte, ((length+2)/3)*4)
		n := Base64_encode_stream(data, uint64(length), dst)

		expected := base64.StdEncoding.EncodeToString(data)
		if string(dst[:n]) != expected {
			t.Fatalf("Longueur %d: Got %q, want %q", length, string(dst[:n]), expected)
		}
	}
}

func TestBase64VsCOracle(t *testing.T) {
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("gcc non disponible")
	}

	candidatePaths := []string{
		filepath.Join("..", "sources", "base64_simd.c"),
		filepath.Join("..", "..", "sources", "base64_simd.c"),
		filepath.Join("..", "..", "spec", "c_sources", "testdata", "c_sources", "base64_simd.c"),
	}
	var cSrcPath string
	for _, p := range candidatePaths {
		if abs, err := filepath.Abs(p); err == nil {
			if _, err := os.Stat(abs); err == nil {
				cSrcPath = abs
				break
			}
		}
	}
	if cSrcPath == "" {
		t.Skip("source C base64_simd.c introuvable pour l'oracle")
	}

	testHarness := `
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>

size_t base64_encode_stream(const uint8_t *src, size_t len, char *dst);

int main(int argc, char **argv) {
    if (argc < 3) return 1;
    FILE *f = fopen(argv[1], "rb");
    if (!f) return 2;
    fseek(f, 0, SEEK_END);
    long sz = ftell(f);
    fseek(f, 0, SEEK_SET);

    uint8_t *buf = (uint8_t*)malloc(sz);
    if (fread(buf, 1, sz, f) != (size_t)sz) {
        free(buf);
        fclose(f);
        return 3;
    }
    fclose(f);

    size_t outCap = ((sz + 2) / 3) * 4;
    char *outBuf = (char*)malloc(outCap);
    size_t outLen = base64_encode_stream(buf, sz, outBuf);

    FILE *out = fopen(argv[2], "wb");
    if (!out) {
        free(buf);
        free(outBuf);
        return 4;
    }
    fwrite(outBuf, 1, outLen, out);
    fclose(out);
    free(buf);
    free(outBuf);
    return 0;
}
`
	tmpDir := t.TempDir()
	harnessFile := filepath.Join(tmpDir, "harness.c")
	binFile := filepath.Join(tmpDir, "oracle_b64")
	if err := os.WriteFile(harnessFile, []byte(testHarness), 0o644); err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("gcc", "-O2", "-o", binFile, harnessFile, cSrcPath)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("gcc compile failed: %v\n%s", err, string(out))
	}

	rng := rand.New(rand.NewSource(0xABCD1234))
	for iter := 0; iter < 100; iter++ {
		len_ := rng.Intn(4096)
		raw := make([]byte, len_)
		rng.Read(raw)

		inPath := filepath.Join(tmpDir, "in.bin")
		outPath := filepath.Join(tmpDir, "out.bin")
		if err := os.WriteFile(inPath, raw, 0o644); err != nil {
			t.Fatal(err)
		}

		runCmd := exec.Command(binFile, inPath, outPath)
		if out, err := runCmd.CombinedOutput(); err != nil {
			t.Fatalf("oracle execution failed: %v\n%s", err, string(out))
		}

		cResult, err := os.ReadFile(outPath)
		if err != nil {
			t.Fatal(err)
		}

		goDst := make([]byte, ((len_+2)/3)*4)
		n := Base64_encode_stream(raw, uint64(len_), goDst)

		if !bytes.Equal(goDst[:n], cResult) {
			t.Fatalf("Divergence iter %d (len %d): Go=%q vs C=%q", iter, len_, string(goDst[:n]), string(cResult))
		}
	}
}

func TestBase64ZeroAlloc(t *testing.T) {
	data := []byte("Hello World! Dogfooding sgoiter bit-exact.")
	dst := make([]byte, 128)

	allocs := testing.AllocsPerRun(1000, func() {
		_ = Base64_encode_stream(data, uint64(len(data)), dst)
	})
	if allocs != 0 {
		t.Fatalf("Allocs/op = %.2f, attendu 0", allocs)
	}
}
