package c2json

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func buildOracleJSON(t *testing.T, tmpDir, srcC, srcH string) string {
	cCode := fmt.Sprintf(`#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include "%s"
#include "%s"

int main(int argc, char **argv) {
    if (argc < 4) return 1;
    const char *mode = argv[1];
    const char *json_path = argv[2];
    const char *key = argv[3];

    FILE *fin = fopen(json_path, "rb");
    if (!fin) return 2;
    fseek(fin, 0, SEEK_END);
    long in_size = ftell(fin);
    fseek(fin, 0, SEEK_SET);

    uint8_t *json = (uint8_t *)malloc(in_size);
    if (fread(json, 1, in_size, fin) != (size_t)in_size) {
        fclose(fin);
        free(json);
        return 3;
    }
    fclose(fin);

    if (strcmp(mode, "validate") == 0) {
        int v = c2json_validate(json, (int)in_size);
        printf("%%d\n", v);
    } else if (strcmp(mode, "get_str") == 0) {
        uint8_t out[4096];
        int n = c2json_get_string(json, (int)in_size, (const uint8_t *)key, (int)strlen(key), out, sizeof(out));
        if (n >= 0) {
            out[n] = 0;
            printf("%%s\n", out);
        } else {
            printf("__ERR__\n");
        }
    } else if (strcmp(mode, "get_int") == 0) {
        int64_t val = 0;
        if (c2json_get_int64(json, (int)in_size, (const uint8_t *)key, (int)strlen(key), &val)) {
            printf("%%lld\n", (long long)val);
        } else {
            printf("__ERR__\n");
        }
    } else if (strcmp(mode, "get_bool") == 0) {
        int val = 0;
        if (c2json_get_bool(json, (int)in_size, (const uint8_t *)key, (int)strlen(key), &val)) {
            printf("%%d\n", val);
        } else {
            printf("__ERR__\n");
        }
    }

    free(json);
    return 0;
}
`, filepath.Base(srcH), filepath.Base(srcC))

	mainPath := filepath.Join(tmpDir, "main.c")
	if err := os.WriteFile(mainPath, []byte(cCode), 0644); err != nil {
		t.Fatalf("impossible d'écrire main.c: %v", err)
	}

	for _, f := range []string{srcC, srcH} {
		data, err := os.ReadFile(f)
		if err != nil {
			t.Fatalf("lecture %s: %v", f, err)
		}
		if err := os.WriteFile(filepath.Join(tmpDir, filepath.Base(f)), data, 0644); err != nil {
			t.Fatalf("copie %s: %v", f, err)
		}
	}

	oracleBin := filepath.Join(tmpDir, "c2json_oracle")
	cmd := exec.Command("gcc", "-O2", "-Wall", "-Wextra", "-Werror", "-o", oracleBin, mainPath)
	cmd.Dir = tmpDir
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("échec compilation gcc -O2 de l'oracle C:\n%s\nErr: %v", string(out), err)
	}

	return oracleBin
}

func runOracle(t *testing.T, bin, mode, jsonPath, key string) string {
	cmd := exec.Command(bin, mode, jsonPath, key)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("exécution oracle c2json échouée: %v\nSortie: %s", err, string(out))
	}
	return strings.TrimSpace(string(out))
}

func TestJSONVsCOracle(t *testing.T) {
	tmpDir := t.TempDir()
	srcC, _ := filepath.Abs("c2json.c")
	srcH, _ := filepath.Abs("c2json.h")

	oracleBin := buildOracleJSON(t, tmpDir, srcC, srcH)

	testJSONs := []struct {
		name string
		json string
	}{
		{
			name: "simple_object",
			json: `{"name": "horos55", "version": 55, "active": true, "sub": false, "empty": null}`,
		},
		{
			name: "nested_structure",
			json: `{"agent": {"id": "ag-01", "role": "Architecte", "level": 99}, "status": "operational", "tokens": 4096}`,
		},
		{
			name: "escaped_strings",
			json: `{"msg": "Hello\nWorld!\tWith \"quotes\" and \\slash", "path": "/devhoros/bin/sgoiter", "flag": true}`,
		},
		{
			name: "invalid_syntax",
			json: `{"unclosed": "brace", "val": 123`,
		},
	}

	for _, tc := range testJSONs {
		t.Run(tc.name, func(t *testing.T) {
			jsonPath := filepath.Join(tmpDir, tc.name+".json")
			if err := os.WriteFile(jsonPath, []byte(tc.json), 0644); err != nil {
				t.Fatalf("écriture json: %v", err)
			}

			// 1. Validation de conformité
			goValid := Validate([]byte(tc.json))
			cValidStr := runOracle(t, oracleBin, "validate", jsonPath, "_")
			cValid := (cValidStr == "1")

			if goValid != cValid {
				t.Fatalf("[%s] divergence validation: Go=%v, C=%v", tc.name, goValid, cValid)
			}

			if !goValid {
				return
			}

			// 2. Extraction de chaînes
			for _, key := range []string{"name", "status", "msg", "path", "unknown_key"} {
				goStr, goOK := GetString([]byte(tc.json), key)
				cStr := runOracle(t, oracleBin, "get_str", jsonPath, key)

				if goOK {
					if goStr != cStr {
						t.Fatalf("[%s, key=%s] divergence GetString: Go=%q, C=%q", tc.name, key, goStr, cStr)
					}
				} else {
					if cStr != "__ERR__" {
						t.Fatalf("[%s, key=%s] Go renvoie faux mais C renvoie %q", tc.name, key, cStr)
					}
				}
			}

			// 3. Extraction d'entiers
			for _, key := range []string{"version", "tokens", "unknown_int"} {
				goInt, goOK := GetInt([]byte(tc.json), key)
				cIntStr := runOracle(t, oracleBin, "get_int", jsonPath, key)

				if goOK {
					expected := fmt.Sprintf("%d", goInt)
					if expected != cIntStr {
						t.Fatalf("[%s, key=%s] divergence GetInt: Go=%d, C=%s", tc.name, key, goInt, cIntStr)
					}
				} else {
					if cIntStr != "__ERR__" {
						t.Fatalf("[%s, key=%s] Go GetInt faux mais C renvoie %s", tc.name, key, cIntStr)
					}
				}
			}

			// 4. Extraction de booléens
			for _, key := range []string{"active", "sub", "flag", "unknown_bool"} {
				goBool, goOK := GetBool([]byte(tc.json), key)
				cBoolStr := runOracle(t, oracleBin, "get_bool", jsonPath, key)

				if goOK {
					expected := "0"
					if goBool {
						expected = "1"
					}
					if expected != cBoolStr {
						t.Fatalf("[%s, key=%s] divergence GetBool: Go=%v, C=%s", tc.name, key, goBool, cBoolStr)
					}
				} else {
					if cBoolStr != "__ERR__" {
						t.Fatalf("[%s, key=%s] Go GetBool faux mais C renvoie %s", tc.name, key, cBoolStr)
					}
				}
			}
		})
	}
}
