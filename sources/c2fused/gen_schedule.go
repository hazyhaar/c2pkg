//go:build ignore

// Génère c2fused_schedule.h depuis sgoiter/spec/fused/fused_schedule.cue.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type fusedSpec struct {
	NQuarterRounds int   `json:"n_quarter_rounds"`
	NPolyBlocks    int   `json:"n_poly_blocks"`
	QRAfter        []int `json:"qr_after"`
}

func main() {
	srcRoot, err := filepath.Abs(filepath.Join(filepath.Dir(os.Args[0])))
	if err != nil {
		fatal(err)
	}
	// go run sets Args[0] to the temp binary; locate via cwd or -dir.
	cwd, _ := os.Getwd()
	cueDir := filepath.Join(cwd, "..", "..", "sgoiter", "spec", "fused")
	if _, err := os.Stat(filepath.Join(cueDir, "fused_schedule.cue")); err != nil {
		cueDir = filepath.Join(srcRoot, "..", "..", "sgoiter", "spec", "fused")
	}
	outH := filepath.Join(cwd, "c2fused_schedule.h")
	if len(os.Args) > 1 {
		outH = os.Args[1]
	}

	cmd := exec.Command("cue", "export", "--out", "json", ".")
	cmd.Dir = cueDir
	raw, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			fatal(fmt.Errorf("cue export: %w: %s", err, bytes.TrimSpace(ee.Stderr)))
		}
		fatal(fmt.Errorf("cue export: %w", err))
	}
	var spec fusedSpec
	if err := json.Unmarshal(raw, &spec); err != nil {
		fatal(err)
	}
	if spec.NQuarterRounds != 80 || spec.NPolyBlocks != 32 {
		fatal(fmt.Errorf("constantes inattendues : qr=%d poly=%d", spec.NQuarterRounds, spec.NPolyBlocks))
	}
	if len(spec.QRAfter) != spec.NPolyBlocks {
		fatal(fmt.Errorf("qr_after : %d entrées, attendu %d", len(spec.QRAfter), spec.NPolyBlocks))
	}

	slot := make([]int, spec.NQuarterRounds)
	for i := range slot {
		slot[i] = -1
	}
	prev := -1
	for k, qr := range spec.QRAfter {
		if qr < 0 || qr >= spec.NQuarterRounds {
			fatal(fmt.Errorf("qr_after[%d]=%d hors [0,%d)", k, qr, spec.NQuarterRounds))
		}
		if qr <= prev {
			fatal(fmt.Errorf("qr_after non monotone : %d puis %d", prev, qr))
		}
		if slot[qr] != -1 {
			fatal(fmt.Errorf("collision au quart %d", qr))
		}
		slot[qr] = k
		prev = qr
	}

	var b strings.Builder
	b.WriteString("/* Code generated from sgoiter/spec/fused/fused_schedule.cue. DO NOT EDIT. */\n")
	b.WriteString("#ifndef C2FUSED_SCHEDULE_H\n#define C2FUSED_SCHEDULE_H\n\n")
	b.WriteString("#include <stdint.h>\n\n")
	fmt.Fprintf(&b, "#define C2FUSED_N_QR %d\n#define C2FUSED_N_POLY %d\n\n", spec.NQuarterRounds, spec.NPolyBlocks)
	b.WriteString("/* slot[q] = index du bloc Poly (0..31) absorbé après le quart q, ou -1. */\n")
	b.WriteString("static const int8_t c2fused_poly_slot[C2FUSED_N_QR] = {\n")
	for q := 0; q < spec.NQuarterRounds; q++ {
		if q%8 == 0 {
			b.WriteString("    ")
		}
		fmt.Fprintf(&b, "%2d", slot[q])
		if q+1 < spec.NQuarterRounds {
			b.WriteString(",")
		}
		if q%8 == 7 || q+1 == spec.NQuarterRounds {
			b.WriteString("\n")
		} else {
			b.WriteString(" ")
		}
	}
	b.WriteString("};\n\n#endif\n")

	if err := os.WriteFile(outH, []byte(b.String()), 0o644); err != nil {
		fatal(err)
	}
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "gen_schedule: %v\n", err)
	os.Exit(1)
}
