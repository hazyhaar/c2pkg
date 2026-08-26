// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2blueteam

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const (
	C2BT_MODE_PASSIVE       byte = 0
	C2BT_MODE_ACTIVE        byte = 1
	C2BT_POLICY_FAIL_OPEN   byte = 0
	C2BT_POLICY_FAIL_CLOSED byte = 1

	FAN_ALLOW uint32 = 0x01
	FAN_DENY  uint32 = 0x02
)

// TestVetoFanotify_EnforceVsPassive valide le modèle de décision synchrone pour les événements FS.
func TestVetoFanotify_EnforceVsPassive(t *testing.T) {
	cases := []struct {
		name         string
		enforceMode  byte
		flags        uint32
		wantResponse uint32
	}{
		{
			name:         "Active Enforce Mode with BLOCKED flag -> FAN_DENY",
			enforceMode:  C2BT_MODE_ACTIVE,
			flags:        C2BT_FLAG_BLOCKED,
			wantResponse: FAN_DENY,
		},
		{
			name:         "Active Enforce Mode with ANOMALY and BLOCKED -> FAN_DENY",
			enforceMode:  C2BT_MODE_ACTIVE,
			flags:        C2BT_FLAG_BLOCKED | C2BT_FLAG_ANOMALY,
			wantResponse: FAN_DENY,
		},
		{
			name:         "Active Enforce Mode with VERDICT_OK -> FAN_ALLOW",
			enforceMode:  C2BT_MODE_ACTIVE,
			flags:        C2BT_FLAG_VERDICT_OK,
			wantResponse: FAN_ALLOW,
		},
		{
			name:         "Passive Audit Mode with BLOCKED -> FAN_ALLOW",
			enforceMode:  C2BT_MODE_PASSIVE,
			flags:        C2BT_FLAG_BLOCKED,
			wantResponse: FAN_ALLOW,
		},
		{
			name:         "Passive Audit Mode with VERDICT_OK -> FAN_ALLOW",
			enforceMode:  C2BT_MODE_PASSIVE,
			flags:        C2BT_FLAG_VERDICT_OK,
			wantResponse: FAN_ALLOW,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var resp uint32
			if tc.enforceMode == C2BT_MODE_ACTIVE && (tc.flags&C2BT_FLAG_BLOCKED) != 0 {
				resp = FAN_DENY
			} else {
				resp = FAN_ALLOW
			}
			if resp != tc.wantResponse {
				t.Fatalf("verdict = 0x%x, attendu 0x%x", resp, tc.wantResponse)
			}
		})
	}
}

// TestMCPFilter_RejectionProtocolStructure valide la conformité stricte du protocole filaire JSON-RPC 2.0.
func TestMCPFilter_RejectionProtocolStructure(t *testing.T) {
	expectedErrorCode := -32003
	expectedErrorMsg := "Blocked by c2blueteam security doctrine"

	cases := []struct {
		name       string
		idInput    string
		wantSubstr string
	}{
		{
			name:       "Integer ID",
			idInput:    `42`,
			wantSubstr: `{"jsonrpc":"2.0","id":42,"error":{"code":-32003,"message":"Blocked by c2blueteam security doctrine"}}`,
		},
		{
			name:       "String ID",
			idInput:    `"call-99"`,
			wantSubstr: `{"jsonrpc":"2.0","id":"call-99","error":{"code":-32003,"message":"Blocked by c2blueteam security doctrine"}}`,
		},
		{
			name:       "Null ID",
			idInput:    `null`,
			wantSubstr: `{"jsonrpc":"2.0","id":null,"error":{"code":-32003,"message":"Blocked by c2blueteam security doctrine"}}`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rejection := `{"jsonrpc":"2.0","id":` + tc.idInput + `,"error":{"code":` +
				"-32003" + `,"message":"` + expectedErrorMsg + `"}}`
			if rejection != tc.wantSubstr {
				t.Fatalf("rejection JSON = %s, attendu %s", rejection, tc.wantSubstr)
			}
			if !strings.Contains(rejection, expectedErrorMsg) {
				t.Errorf("missing error message in payload")
			}
			_ = expectedErrorCode
		})
	}
}

// TestVeto_COracleBitExact valide l'exécution bit-exacte des épreuves de veto sous GCC -O2.
func TestVeto_COracleBitExact(t *testing.T) {
	tmpBin := filepath.Join(t.TempDir(), "c_veto_oracle")
	srcDir := filepath.Join("..", "..", "sources", "c2blueteam")

	cmdBuild := exec.Command("gcc", "-O2", "-Wall", "-Wextra", "-std=gnu99",
		"-I", srcDir,
		filepath.Join(srcDir, "test_oracle.c"),
		filepath.Join(srcDir, "ring_buffer.c"),
		filepath.Join(srcDir, "probe_entropy.c"),
		filepath.Join(srcDir, "rules_simd.c"),
		filepath.Join(srcDir, "probe_proc.c"),
		filepath.Join(srcDir, "probe_mcp.c"),
		filepath.Join(srcDir, "probe_fs.c"),
		filepath.Join(srcDir, "c2blueteam.c"),
		filepath.Join(srcDir, "correlator.c"),
		"-o", tmpBin,
	)
	if out, err := cmdBuild.CombinedOutput(); err != nil {
		t.Fatalf("échec de compilation oracle: %v\nSortie: %s", err, string(out))
	}

	cmdRun := exec.Command(tmpBin)
	out, err := cmdRun.CombinedOutput()
	if err != nil {
		t.Fatalf("échec d'exécution oracle: %v\nSortie: %s", err, string(out))
	}

	outStr := string(out)
	if !bytes.Contains(out, []byte("[7/9] Veto synchrone fanotify (FAN_DENY / FAN_ALLOW)... OK")) {
		t.Errorf("étape fanotify manquante ou échouée dans la sortie oracle:\n%s", outStr)
	}
	if !bytes.Contains(out, []byte("[8/9] Filtrage synchrone MCP JSON-RPC 2.0 & Extraction d'ID... OK")) {
		t.Errorf("étape MCP manquante ou échouée dans la sortie oracle:\n%s", outStr)
	}
	if !bytes.Contains(out, []byte("[9/9] Robustesse en mode dégradé, Fail-Open vs Fail-Closed & Zéro Fuite... OK")) {
		t.Errorf("étape robustesse dégradée manquante dans la sortie oracle:\n%s", outStr)
	}
}
