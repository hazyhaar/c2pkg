// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2oracle

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestOracleCompilationAndNegativeControl(t *testing.T) {
	cCode := `
#include <stdio.h>
#include <stdint.h>

int main(int argc, char **argv) {
    if (argc < 2) return 1;
    printf("ORACLE_OK:%s\n", argv[1]);
    return 0;
}
`
	tmpDir := t.TempDir()
	cFile := filepath.Join(tmpDir, "simple_oracle.c")
	if err := os.WriteFile(cFile, []byte(cCode), 0o644); err != nil {
		t.Fatal(err)
	}

	harness := NewHarness(cFile)
	bin := harness.Compile(t)

	out, err := RunOracleBinary(bin, "dogfood")
	if err != nil {
		t.Fatalf("run failed: %v", err)
	}

	if !bytes.Contains(out, []byte("ORACLE_OK:dogfood")) {
		t.Fatalf("unexpected oracle output: %s", string(out))
	}

	// Test du contrôle négatif
	validData := []byte("secret_oracle_valid_result")
	AssertNegativeControl(t, func(mutated []byte) bool {
		return bytes.Equal(mutated, validData)
	}, validData)
}
