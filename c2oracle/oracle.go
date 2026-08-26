// SPDX-License-Identifier: Apache-2.0 OR MIT

// Package c2oracle standardise les harnais d'oracles C (gcc -O2 ASan/UBSan)
// et formalise l'exigence de contrôles négatifs (mutation testing anti-tautologie).
package c2oracle

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// OracleHarness configure et pilote l'exécution d'un binaire oracle C.
type OracleHarness struct {
	SourceFiles []string
	IncludeDirs []string
	Compiler    string
	Flags       []string
}

// NewHarness instancie un harnais standard avec GCC -O2 et assainisseurs mémoire.
func NewHarness(sources ...string) *OracleHarness {
	return &OracleHarness{
		SourceFiles: sources,
		Compiler:    "gcc",
		Flags: []string{
			"-O2",
			"-mavx2",
			"-fsanitize=address,undefined",
			"-Wall",
			"-Wextra",
		},
	}
}

// Compile construit le binaire oracle dans un répertoire temporaire et renvoie son chemin.
func (h *OracleHarness) Compile(t *testing.T) string {
	t.Helper()
	if _, err := exec.LookPath(h.Compiler); err != nil {
		t.Skipf("%s non disponible sur le système hôte", h.Compiler)
	}

	for _, src := range h.SourceFiles {
		if _, err := os.Stat(src); err != nil {
			t.Fatalf("Fichier source C oracle introuvable : %s (%v)", src, err)
		}
	}

	tmpDir := t.TempDir()
	binPath := filepath.Join(tmpDir, "c_oracle.bin")

	var args []string
	args = append(args, h.Flags...)
	for _, inc := range h.IncludeDirs {
		args = append(args, "-I", inc)
	}
	args = append(args, h.SourceFiles...)
	args = append(args, "-lm", "-o", binPath)

	cmd := exec.Command(h.Compiler, args...)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("Échec compilation oracle C (%s):\n%s", err, string(out))
	}

	return binPath
}

// AssertNegativeControl prouve mécaniquement que le prédicat de test rejette une sortie mutée
// (contrôle anti-tautologie obligatoire).
func AssertNegativeControl(t *testing.T, checkFn func(mutated []byte) bool, valid []byte) {
	t.Helper()
	if len(valid) == 0 {
		return
	}

	mutated := make([]byte, len(valid))
	copy(mutated, valid)

	// Altération d'un octet clé
	mutated[0] ^= 0xFF

	if checkFn(mutated) == true {
		t.Fatalf("FAIL-CLOSED : Le contrôle négatif a échoué ! La vérification accepte une sortie corrompue (test tautologique)")
	}
}

// RunOracleBinary exécute le binaire compilé avec les arguments donnés.
func RunOracleBinary(binPath string, args ...string) ([]byte, error) {
	cmd := exec.Command(binPath, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("exécution oracle C en échec (%w) : %s", err, string(out))
	}
	return out, nil
}
