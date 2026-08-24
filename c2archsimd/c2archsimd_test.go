package c2archsimd

import (
	"bytes"
	"encoding/binary"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

type xorshift struct{ s uint64 }

func (r *xorshift) next() uint64 {
	r.s ^= r.s >> 12
	r.s ^= r.s << 25
	r.s ^= r.s >> 27
	return r.s * 0x2545F4914F6CDD1D
}

func foldMix(acc, val uint64) uint64 {
	return acc ^ (val + 0x9E3779B97F4A7C15 + (acc << 6) + (acc >> 2))
}

func TestKAT_HexEncodeDecode(t *testing.T) {
	uuidBin := [16]byte{
		0x01, 0x91, 0x7F, 0x8B, 0xC9, 0xA0, 0x71, 0x11,
		0x80, 0x00, 0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC,
	}
	var hexOut [32]byte
	C2archsimd_hex_encode16(uuidBin[:], hexOut[:])

	expectedHex := "01917f8bc9a071118000123456789abc"
	if string(hexOut[:]) != expectedHex {
		t.Fatalf("C2archsimd_hex_encode16 KAT failed: got %s, want %s", string(hexOut[:]), expectedHex)
	}

	var decoded [16]byte
	if C2archsimd_hex_decode32(hexOut[:], decoded[:]) != 0 {
		t.Fatalf("C2archsimd_hex_decode32 KAT returned error on valid hex")
	}
	if !bytes.Equal(decoded[:], uuidBin[:]) {
		t.Fatalf("C2archsimd_hex_decode32 roundtrip failed: got %x, want %x", decoded, uuidBin)
	}

	// Adversarial KAT: Invalid Hex Characters
	invalidHexInputs := []string{
		"01917f8bc9a071118000123456789abg",    // 'g' at end
		"Z1917f8bc9a071118000123456789abc",    // 'Z' at start
		"01917f8bc9a0711180!0123456789abc",    // '!' in middle
		"01917f8bc9a0711180\x000123456789abc", // NULL byte
	}
	for _, inv := range invalidHexInputs {
		var invIn [32]byte
		copy(invIn[:], inv)
		var invOut [16]byte
		for k := range invOut {
			invOut[k] = 0xAA
		}
		if C2archsimd_hex_decode32(invIn[:], invOut[:]) == 0 {
			t.Fatalf("C2archsimd_hex_decode32 succeeded on invalid hex input %q", inv)
		}
		var zero16 [16]byte
		if invOut != zero16 {
			t.Fatalf("C2archsimd_hex_decode32 failed to sanitize output on invalid hex %q: got %x", inv, invOut)
		}
	}
}

func TestKAT_VintLens(t *testing.T) {
	var in [32]byte
	in[0] = 0x25 // 1 octet
	in[1] = 0x7B // 2 octets
	in[2] = 0x9D // 4 octets
	in[3] = 0xC2 // 8 octets

	var out [32]byte
	C2archsimd_vint_lens32(in[:], out[:])

	if out[0] != 1 || out[1] != 2 || out[2] != 4 || out[3] != 8 {
		t.Fatalf("C2archsimd_vint_lens32 KAT mismatch: got [%d, %d, %d, %d], want [1, 2, 4, 8]", out[0], out[1], out[2], out[3])
	}
}

func TestZeroAlloc(t *testing.T) {
	var in16 [16]byte
	var out32 [32]byte
	var in32 [32]byte
	var table16 C2archsimd_table16_t
	var table256 C2archsimd_table256_t

	allocs := testing.AllocsPerRun(1000, func() {
		C2archsimd_hex_encode16(in16[:], out32[:])
		C2archsimd_hex_decode32(out32[:], in16[:])
		C2archsimd_lut16(in32[:], &table16, out32[:])
		C2archsimd_lut256(in32[:], &table256, out32[:])
		C2archsimd_vint_lens32(in32[:], out32[:])
	})

	if allocs != 0 {
		t.Fatalf("c2archsimd violated 0-allocation rule: allocs/op = %.2f", allocs)
	}
}

func TestC2ARCHSIMD_ParityVsCOracle(t *testing.T) {
	srcCandidates := []string{
		filepath.Join("..", "..", "sources", "c2archsimd"),
		filepath.Join("..", "..", "c2simd", "sources", "c2archsimd"),
		"/devhoros/c2simd/sources/c2archsimd",
	}
	var srcDir string
	for _, c := range srcCandidates {
		if _, err := os.Stat(filepath.Join(c, "test_c2archsimd_oracle.c")); err == nil {
			srcDir = c
			break
		}
	}
	if srcDir == "" {
		t.Fatalf("CRITICAL GATE FAILURE: Oracle C source absent dans les chemins candidats: %v", srcCandidates)
	}
	oracleSrc := filepath.Join(srcDir, "test_c2archsimd_oracle.c")
	cSrc := filepath.Join(srcDir, "c2archsimd.c")
	cSimdSrc := filepath.Join(srcDir, "c2archsimd_simd.c")

	tmpBin := filepath.Join(t.TempDir(), "test_c2archsimd_oracle")
	cmd := exec.Command("gcc", "-O2", "-mavx2", "-fsanitize=address,undefined", "-I", srcDir, oracleSrc, cSrc, cSimdSrc, "-lm", "-o", tmpBin)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("gcc compile failed with ASan/UBSan: %v\n%s", err, string(out))
	}

	cOut, err := exec.Command(tmpBin).CombinedOutput()
	if err != nil {
		t.Fatalf("C oracle execution failed under ASan/UBSan: %v\n%s", err, string(cOut))
	}

	wantLut16, wantLut256, wantHex, wantVint := parseFolds(t, string(cOut))

	rng := xorshift{s: 0x853c49e6748fea9b}
	var foldLut16, foldLut256, foldHex, foldVint uint64

	var table16 C2archsimd_table16_t
	var table256 C2archsimd_table256_t
	for k := 0; k < 16; k++ {
		table16.B[k] = byte(k * 17)
	}
	for k := 0; k < 256; k++ {
		table256.B[k] = byte(k*31 + 7)
	}

	var inBuf [32]byte
	var outBuf [32]byte
	var outAvx2 [32]byte
	var outHex [32]byte
	var outHexAvx2 [32]byte

	for iter := 0; iter < 1000000; iter++ {
		for b := 0; b < 32; b += 8 {
			r := rng.next()
			binary.LittleEndian.PutUint64(inBuf[b:b+8], r)
		}

		// 1. LUT16 (Scalar & AVX2)
		C2archsimd_lut16(inBuf[:], &table16, outBuf[:])
		if HasAVX2() {
			C2archsimd_lut16_avx2(inBuf[:], &table16, outAvx2[:])
			if !bytes.Equal(outBuf[:], outAvx2[:]) {
				t.Fatalf("LUT16 Go scalar vs Go AVX2 mismatch at iter %d", iter)
			}
		}
		for k := 0; k < 32; k += 8 {
			val := binary.LittleEndian.Uint64(outBuf[k : k+8])
			foldLut16 = foldMix(foldLut16, val)
		}

		// 2. LUT256 Directe
		C2archsimd_lut256(inBuf[:], &table256, outBuf[:])
		for k := 0; k < 32; k += 8 {
			val := binary.LittleEndian.Uint64(outBuf[k : k+8])
			foldLut256 = foldMix(foldLut256, val)
		}

		// 3. Hex Encode (Scalar & AVX2)
		var in16 [16]byte
		copy(in16[:], inBuf[:16])
		C2archsimd_hex_encode16(in16[:], outHex[:])
		if HasAVX2() {
			C2archsimd_hex_encode16_avx2(in16[:], outHexAvx2[:])
			if !bytes.Equal(outHex[:], outHexAvx2[:]) {
				t.Fatalf("HexEncode Go scalar vs Go AVX2 mismatch at iter %d", iter)
			}
		}
		for k := 0; k < 32; k += 8 {
			val := binary.LittleEndian.Uint64(outHex[k : k+8])
			foldHex = foldMix(foldHex, val)
		}

		// 4. Vint Lens (Scalar & AVX2)
		C2archsimd_vint_lens32(inBuf[:], outBuf[:])
		if HasAVX2() {
			C2archsimd_vint_lens32_avx2(inBuf[:], outAvx2[:])
			if !bytes.Equal(outBuf[:], outAvx2[:]) {
				t.Fatalf("VintLens Go scalar vs Go AVX2 mismatch at iter %d", iter)
			}
		}
		for k := 0; k < 32; k += 8 {
			val := binary.LittleEndian.Uint64(outBuf[k : k+8])
			foldVint = foldMix(foldVint, val)
		}
	}

	if foldLut16 != wantLut16 || foldLut256 != wantLut256 || foldHex != wantHex || foldVint != wantVint {
		t.Fatalf("PARITY MISMATCH:\n  Go Folds: lut16=0x%016X lut256=0x%016X hex=0x%016X vint=0x%016X\n  C  Folds: lut16=0x%016X lut256=0x%016X hex=0x%016X vint=0x%016X",
			foldLut16, foldLut256, foldHex, foldVint, wantLut16, wantLut256, wantHex, wantVint)
	}

	t.Logf("PARITÉ BIT-EXACTE 100%% VALIDÉE SCALAIRE + AVX2 VS GCC -O2 (ASAN/UBSAN OK) : Folds lut16=0x%016X lut256=0x%016X hex=0x%016X vint=0x%016X",
		foldLut16, foldLut256, foldHex, foldVint)
}

func parseFolds(t *testing.T, text string) (lut16, lut256, hexFold, vint uint64) {
	for _, line := range strings.Split(text, "\n") {
		fs := strings.Fields(line)
		if len(fs) != 3 || fs[0] != "FOLD" {
			continue
		}
		val, err := strconv.ParseUint(strings.TrimPrefix(fs[2], "0x"), 16, 64)
		if err != nil {
			t.Fatalf("Failed to parse fold %q: %v", line, err)
		}
		switch fs[1] {
		case "lut16":
			lut16 = val
		case "lut256":
			lut256 = val
		case "hex":
			hexFold = val
		case "vint":
			vint = val
		}
	}
	if lut16 == 0 || lut256 == 0 || hexFold == 0 || vint == 0 {
		t.Fatalf("Missing folds in C oracle output:\n%s", text)
	}
	return
}
