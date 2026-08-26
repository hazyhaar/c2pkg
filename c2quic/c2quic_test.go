// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2quic

import (
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

func foldMix(acc, x uint64) uint64 {
	return acc ^ (x + 0x9E3779B97F4A7C15 + (acc << 6) + (acc >> 2))
}

func TestKAT_RFC9000_Vint(t *testing.T) {
	var r C2quic_vint_t
	if st := C2quic_vint_decode([]byte{0x25}, 1, 0, &r); st != OK || r.Value != 37 || r.Nbytes != 1 {
		t.Fatalf("37: st=%d v=%d n=%d", st, r.Value, r.Nbytes)
	}
	if st := C2quic_vint_decode([]byte{0x7b, 0xbd}, 2, 0, &r); st != OK || r.Value != 15293 || r.Nbytes != 2 {
		t.Fatalf("15293: st=%d v=%d n=%d", st, r.Value, r.Nbytes)
	}
	if st := C2quic_vint_decode([]byte{0x9d, 0x7f, 0x3e, 0x7d}, 4, 0, &r); st != OK || r.Value != 494878333 || r.Nbytes != 4 {
		t.Fatalf("494878333: st=%d v=%d n=%d", st, r.Value, r.Nbytes)
	}
	if st := C2quic_vint_decode([]byte{0xc2, 0x19, 0x7c, 0x5e, 0xff, 0x14, 0xe8, 0x8c}, 8, 0, &r); st != OK || r.Value != 151288809941952652 || r.Nbytes != 8 {
		t.Fatalf("8B: st=%d v=%d n=%d", st, r.Value, r.Nbytes)
	}
	enc := make([]byte, 8)
	if st := C2quic_vint_encode(37, enc, 8, 0, &r); st != OK || enc[0] != 0x25 || r.Nbytes != 1 {
		t.Fatalf("encode 37: st=%d %x n=%d", st, enc[0], r.Nbytes)
	}
	if st := C2quic_vint_encode(15293, enc, 8, 0, &r); st != OK || enc[0] != 0x7b || enc[1] != 0xbd {
		t.Fatalf("encode 15293: st=%d %x", st, enc[:2])
	}
}

func TestKAT_Frames(t *testing.T) {
	var r C2quic_vint_t
	body := make([]byte, 128)
	o := uint32(0)
	body[o] = 0x00
	o++
	body[o] = 0x01
	o++
	body[o] = 0x0F
	o++
	C2quic_vint_encode(7, body, 128, o, &r)
	o += r.Nbytes
	C2quic_vint_encode(100, body, 128, o, &r)
	o += r.Nbytes
	C2quic_vint_encode(4, body, 128, o, &r)
	o += r.Nbytes
	body[o] = 'A'
	o++
	body[o] = 'B'
	o++
	body[o] = 'C'
	o++
	body[o] = 'D'
	o++
	body[o] = 0x04
	o++
	C2quic_vint_encode(7, body, 128, o, &r)
	o += r.Nbytes
	C2quic_vint_encode(0x42, body, 128, o, &r)
	o += r.Nbytes
	C2quic_vint_encode(4, body, 128, o, &r)
	o += r.Nbytes
	body[o] = 0x02
	o++
	C2quic_vint_encode(20, body, 128, o, &r)
	o += r.Nbytes
	C2quic_vint_encode(1, body, 128, o, &r)
	o += r.Nbytes
	C2quic_vint_encode(0, body, 128, o, &r)
	o += r.Nbytes
	C2quic_vint_encode(5, body, 128, o, &r)
	o += r.Nbytes
	body[o] = 0x31
	o++
	C2quic_vint_encode(3, body, 128, o, &r)
	o += r.Nbytes
	body[o] = 'x'
	o++
	body[o] = 'y'
	o++
	body[o] = 'z'
	o++
	var slab C2quic_frame_slab_t
	st := C2quic_unpack_frames(body, o, &slab)
	if st != OK || slab.N != 5 {
		t.Fatalf("unpack st=%d n=%d", st, slab.N)
	}
	if slab.Frames[1].Typ != 0x0F || slab.Frames[1].Id != 7 || slab.Frames[1].Offset != 100 || slab.Frames[1].Payload_len != 4 {
		t.Fatalf("stream %+v", slab.Frames[1])
	}
	if slab.Frames[2].Typ != 0x04 || slab.Frames[2].Extra != 0x42 {
		t.Fatalf("reset %+v", slab.Frames[2])
	}
	if slab.Frames[3].Typ != 0x02 || slab.Frames[3].Id != 20 || slab.Frames[3].Extra != 5 {
		t.Fatalf("ack %+v", slab.Frames[3])
	}
	if slab.Frames[4].Typ != 0x31 || slab.Frames[4].Payload_len != 3 {
		t.Fatalf("dgram %+v", slab.Frames[4])
	}
}

func goFolds() (vint, hp, frames uint64) {
	var r C2quic_vint_t
	buf := make([]byte, 64)
	rng := xorshift{s: 0x853c49e6748fea9b}
	for i := 0; i < 1000000; i++ {
		v := rng.next() >> 2
		st := C2quic_vint_encode(v, buf, 64, 0, &r)
		if st != OK {
			return
		}
		st = C2quic_vint_decode(buf, 64, 0, &r)
		if st != OK || r.Value != v {
			return
		}
		vint = foldMix(vint, r.Value)
		vint = foldMix(vint, uint64(r.Nbytes))
		var b C2quic_batch8_t
		for k := 0; k < 8; k++ {
			vk := rng.next() >> 2
			b.Offs[k] = uint32(k * 8)
			C2quic_vint_encode(vk, buf, 64, b.Offs[k], &r)
		}
		C2quic_vint_decode_batch8_scalar(buf, 64, &b)
		for k := 0; k < 8; k++ {
			vint = foldMix(vint, b.Vals[k])
			vint = foldMix(vint, uint64(b.Nbytes[k]))
			vint = foldMix(vint, uint64(b.Status[k]))
		}
	}

	rng.s = 0x853c49e6748fea9b
	pkt := make([]byte, 64)
	for i := 0; i < 1000000; i++ {
		pnLen := uint32((rng.next() & 3) + 1)
		pnOff := uint32(1)
		for j := range pkt {
			pkt[j] = 0
		}
		var io C2quic_hp_io_t
		if rng.next()&1 != 0 {
			pkt[0] = 0xC0 | byte(pnLen-1)
		} else {
			pkt[0] = 0x40 | byte(pnLen-1)
		}
		pn := rng.next() & 0xFFFFFFFF
		for k := uint32(0); k < pnLen; k++ {
			pkt[pnOff+k] = byte(pn >> (8 * (pnLen - 1 - k)))
		}
		for k := 0; k < 5; k++ {
			io.Mask[k] = byte(rng.next())
		}
		for k := 5; k < 37; k++ {
			pkt[k] = 0xA5
		}
		if C2quic_hp_sample(pkt, 64, pnOff, &io) != OK {
			return
		}
		if C2quic_hp_apply(pkt, 64, pnOff, pnLen, &io) != OK {
			return
		}
		if C2quic_hp_remove(pkt, 64, pnOff, &io) != OK || io.Pn_len != pnLen {
			return
		}
		pnMask := (uint64(1) << (8 * pnLen)) - 1
		if C2quic_hp_read_pn(pkt, 64, pnOff, pnLen, &io) != OK || io.Pn != (pn&pnMask) {
			return
		}
		hp = foldMix(hp, uint64(pkt[0]))
		hp = foldMix(hp, io.Pn)
		hp = foldMix(hp, uint64(io.Sample[0]))
		hp = foldMix(hp, uint64(io.Sample[15]))
	}

	var r2 C2quic_vint_t
	var slab C2quic_frame_slab_t
	body := make([]byte, 128)
	o := uint32(0)
	body[o] = 0x00
	o++
	body[o] = 0x01
	o++
	body[o] = 0x0F
	o++
	C2quic_vint_encode(7, body, 128, o, &r2)
	o += r2.Nbytes
	C2quic_vint_encode(100, body, 128, o, &r2)
	o += r2.Nbytes
	C2quic_vint_encode(4, body, 128, o, &r2)
	o += r2.Nbytes
	body[o] = 'A'
	o++
	body[o] = 'B'
	o++
	body[o] = 'C'
	o++
	body[o] = 'D'
	o++
	body[o] = 0x04
	o++
	C2quic_vint_encode(7, body, 128, o, &r2)
	o += r2.Nbytes
	C2quic_vint_encode(0x42, body, 128, o, &r2)
	o += r2.Nbytes
	C2quic_vint_encode(4, body, 128, o, &r2)
	o += r2.Nbytes
	body[o] = 0x02
	o++
	C2quic_vint_encode(20, body, 128, o, &r2)
	o += r2.Nbytes
	C2quic_vint_encode(1, body, 128, o, &r2)
	o += r2.Nbytes
	C2quic_vint_encode(0, body, 128, o, &r2)
	o += r2.Nbytes
	C2quic_vint_encode(5, body, 128, o, &r2)
	o += r2.Nbytes
	body[o] = 0x31
	o++
	C2quic_vint_encode(3, body, 128, o, &r2)
	o += r2.Nbytes
	body[o] = 'x'
	o++
	body[o] = 'y'
	o++
	body[o] = 'z'
	o++
	C2quic_unpack_frames(body, o, &slab)
	frames = foldMix(frames, uint64(slab.N))
	frames = foldMix(frames, slab.Frames[1].Id)
	frames = foldMix(frames, uint64(slab.Frames[4].Payload_len))

	rng.s = 0x123456789abcdef0
	b := make([]byte, 256)
	for i := 0; i < 1000000; i++ {
		oo := uint32(0)
		sid := rng.next() & 0xFFFF
		plen := (rng.next() & 15) + 1
		var sl C2quic_frame_slab_t
		b[oo] = 0x0A
		oo++
		C2quic_vint_encode(sid, b, 256, oo, &r2)
		oo += r2.Nbytes
		C2quic_vint_encode(plen, b, 256, oo, &r2)
		oo += r2.Nbytes
		for k := uint32(0); k < uint32(plen); k++ {
			b[oo] = byte(rng.next())
			oo++
		}
		st := C2quic_unpack_frames(b, oo, &sl)
		if st != OK || sl.N != 1 || sl.Frames[0].Id != sid || sl.Frames[0].Payload_len != uint32(plen) {
			return
		}
		frames = foldMix(frames, sl.Frames[0].Id)
		frames = foldMix(frames, uint64(sl.Frames[0].Payload_len))
		frames = foldMix(frames, uint64(sl.Frames[0].Payload_off))
	}
	return
}

func TestVintRandom1e6(t *testing.T) {
	var r C2quic_vint_t
	rng := xorshift{s: 0x853c49e6748fea9b}
	buf := make([]byte, 8)
	for i := 0; i < 1000000; i++ {
		v := rng.next() >> 2
		if st := C2quic_vint_encode(v, buf, 8, 0, &r); st != OK {
			t.Fatalf("encode i=%d", i)
		}
		if st := C2quic_vint_decode(buf, r.Nbytes, 0, &r); st != OK || r.Value != v {
			t.Fatalf("decode i=%d got=%d want=%d", i, r.Value, v)
		}
	}
}

func TestC2QUIC_ParityVsCOracle(t *testing.T) {
	srcCandidates := []string{
		filepath.Join("..", "..", "sources", "c2quic"),
		filepath.Join("..", "..", "c2simd", "sources", "c2quic"),
		filepath.Join("sources", "c2quic"),
	}
	var srcDir string
	for _, c := range srcCandidates {
		if _, err := os.Stat(filepath.Join(c, "test_c2quic_oracle.c")); err == nil {
			srcDir = c
			break
		}
	}
	if srcDir == "" {
		t.Fatalf("CRITICAL GATE FAILURE: oracle C absent dans les chemins candidats: %v", srcCandidates)
	}
	tmpBin := filepath.Join(t.TempDir(), "c2quic_oracle")
	cmd := exec.Command("gcc", "-O2", "-mavx2", "-I", srcDir,
		filepath.Join(srcDir, "test_c2quic_oracle.c"),
		filepath.Join(srcDir, "c2quic.c"),
		filepath.Join(srcDir, "c2quic_simd.c"),
		"-o", tmpBin)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("gcc oracle: %v\n%s", err, out)
	}
	asanBin := filepath.Join(t.TempDir(), "c2quic_asan")
	asan := exec.Command("gcc", "-O1", "-g", "-fsanitize=address,undefined", "-I", srcDir,
		filepath.Join(srcDir, "test_c2quic_oracle.c"),
		filepath.Join(srcDir, "c2quic.c"),
		filepath.Join(srcDir, "c2quic_simd.c"),
		"-o", asanBin)
	if out, err := asan.CombinedOutput(); err != nil {
		t.Fatalf("gcc asan oracle: %v\n%s", err, out)
	}
	if out, err := exec.Command(asanBin).CombinedOutput(); err != nil {
		t.Fatalf("asan oracle: %v\n%s", err, out)
	}
	boundsBin := filepath.Join(t.TempDir(), "c2quic_bounds")
	bnd := exec.Command("gcc", "-O1", "-g", "-fsanitize=address,undefined", "-I", srcDir,
		filepath.Join(srcDir, "test_c2quic_bounds.c"),
		filepath.Join(srcDir, "c2quic.c"),
		"-o", boundsBin)
	if out, err := bnd.CombinedOutput(); err != nil {
		t.Fatalf("gcc asan bounds: %v\n%s", err, out)
	}
	if out, err := exec.Command(boundsBin).CombinedOutput(); err != nil {
		t.Fatalf("asan bounds: %v\n%s", err, out)
	}
	out, err := exec.Command(tmpBin).CombinedOutput()
	if err != nil {
		t.Fatalf("oracle run: %v\n%s", err, out)
	}
	text := string(out)
	if !strings.Contains(text, "ALL C ORACLE TESTS PASSED") {
		t.Fatalf("oracle C: %s", text)
	}
	wantV, wantH, wantF := parseFolds(t, text)
	gotV, gotH, gotF := goFolds()
	if gotV != wantV || gotH != wantH || gotF != wantF {
		t.Fatalf("folds Go vint=0x%016X hp=0x%016X fr=0x%016X C vint=0x%016X hp=0x%016X fr=0x%016X",
			gotV, gotH, gotF, wantV, wantH, wantF)
	}
	t.Logf("parité bit-exacte 1e6 vecteurs gcc -O2: vint=0x%016X hp=0x%016X frames=0x%016X", gotV, gotH, gotF)
}

func parseFolds(t *testing.T, text string) (vint, hp, frames uint64) {
	t.Helper()
	for _, line := range strings.Split(text, "\n") {
		fs := strings.Fields(line)
		if len(fs) != 3 || fs[0] != "FOLD" {
			continue
		}
		n, err := strconv.ParseUint(strings.TrimPrefix(fs[2], "0x"), 16, 64)
		if err != nil {
			t.Fatalf("fold %q: %v", line, err)
		}
		switch fs[1] {
		case "vint":
			vint = n
		case "hp":
			hp = n
		case "frames":
			frames = n
		}
	}
	if vint == 0 || hp == 0 || frames == 0 {
		t.Fatalf("folds manquants dans l'oracle:\n%s", text)
	}
	return
}

func TestZeroAlloc(t *testing.T) {
	buf := make([]byte, 64)
	var r C2quic_vint_t
	var io C2quic_hp_io_t
	var slab C2quic_frame_slab_t
	var b C2quic_batch8_t
	C2quic_vint_encode(15293, buf, 64, 0, &r)
	pkt := make([]byte, 64)
	pkt[0] = 0xC0
	pkt[1] = 0x01
	for i := 5; i < 37; i++ {
		pkt[i] = 0xA5
	}
	io.Mask[0] = 0x3C
	io.Mask[1] = 0x11
	body := []byte{0x0A, 0x07, 0x04, 'A', 'B', 'C', 'D'}
	allocs := testing.AllocsPerRun(1000, func() {
		C2quic_vint_encode(494878333, buf, 64, 0, &r)
		C2quic_vint_decode(buf, 64, 0, &r)
		b.Offs[0] = 0
		C2quic_vint_decode_batch8_scalar(buf, 64, &b)
		C2quic_hp_sample(pkt, 64, 1, &io)
		C2quic_hp_apply(pkt, 64, 1, 1, &io)
		C2quic_hp_remove(pkt, 64, 1, &io)
		C2quic_unpack_frames(body, uint32(len(body)), &slab)
	})
	if allocs != 0 {
		t.Fatalf("allocs/op = %.2f, attendu 0", allocs)
	}
}

func TestHeaderProtectionRoundtrip(t *testing.T) {
	pkt := make([]byte, 64)
	pkt[0] = 0xC3
	binary.BigEndian.PutUint32(pkt[1:5], 0x01020304)
	for i := 5; i < 37; i++ {
		pkt[i] = 0xA5
	}
	var io C2quic_hp_io_t
	copy(io.Mask[:], []byte{0xFF, 0x11, 0x22, 0x33, 0x44})
	if st := C2quic_hp_sample(pkt, 64, 1, &io); st != OK {
		t.Fatal(st)
	}
	if io.Sample[0] != 0xA5 || io.Sample[15] != 0xA5 {
		t.Fatalf("sample %x", io.Sample[:])
	}
	if st := C2quic_hp_apply(pkt, 64, 1, 4, &io); st != OK {
		t.Fatal(st)
	}
	if pkt[0] == 0xC3 {
		t.Fatal("first byte inchangé")
	}
	if st := C2quic_hp_remove(pkt, 64, 1, &io); st != OK || io.Pn_len != 4 {
		t.Fatalf("remove st=%d pn=%d", st, io.Pn_len)
	}
	if st := C2quic_hp_read_pn(pkt, 64, 1, 4, &io); st != OK || pkt[0] != 0xC3 || io.Pn != 0x01020304 {
		t.Fatalf("restore first=%x pn=%x st=%d", pkt[0], io.Pn, st)
	}
}

func TestBounds_EncodeOffPastLen(t *testing.T) {
	var r C2quic_vint_t
	buf := make([]byte, 8)
	if st := C2quic_vint_encode(37, buf, 8, 100, &r); st != Trunc {
		t.Fatalf("attendu TRUNC, st=%d", st)
	}
}

func TestBounds_ReadPN(t *testing.T) {
	var io C2quic_hp_io_t
	buf := make([]byte, 4)
	if st := C2quic_hp_read_pn(buf, 4, 1000, 4, &io); st != Trunc {
		t.Fatalf("attendu TRUNC, st=%d", st)
	}
}

func TestHPRemove_NoMutateOnTrunc(t *testing.T) {
	pkt := []byte{0x40, 0x01, 0x02}
	var io C2quic_hp_io_t
	io.Mask[0] = 0x1F
	first := pkt[0]
	if st := C2quic_hp_remove(pkt, 3, 99, &io); st != Trunc {
		t.Fatalf("attendu TRUNC, st=%d", st)
	}
	if pkt[0] != first {
		t.Fatalf("pkt[0] muté: avant=%x après=%x", first, pkt[0])
	}
}

func TestUnpack_PINGVarint(t *testing.T) {
	var slab C2quic_frame_slab_t
	body := []byte{0x40, 0x01}
	if st := C2quic_unpack_frames(body, 2, &slab); st != OK || slab.N != 1 || slab.Frames[0].Typ != 1 {
		t.Fatalf("PING 0x4001 st=%d n=%d typ=%d", st, slab.N, slab.Frames[0].Typ)
	}
}

func TestOverflowVint(t *testing.T) {
	var r C2quic_vint_t
	enc := make([]byte, 8)
	if st := C2quic_vint_encode(1<<62, enc, 8, 0, &r); st != Overflow {
		t.Fatalf("attendu overflow, st=%d", st)
	}
	if st := C2quic_vint_decode([]byte{0x25}, 0, 0, &r); st != Trunc {
		t.Fatalf("attendu trunc, st=%d", st)
	}
}

func TestKAT_Batch8_SIMDGather_Parity(t *testing.T) {
	rng := xorshift{s: 0xDEADBEEFCAFE1234}
	const totalBuf = 16384
	buf := make([]byte, totalBuf)
	for i := 0; i < totalBuf; i++ {
		buf[i] = byte(rng.next())
	}

	for iter := 0; iter < 10000; iter++ {
		var bSIMD, bScalar C2quic_batch8_t
		for k := 0; k < 8; k++ {
			// Offset garantissant au moins 8 octets disponibles
			off := uint32(rng.next() % (totalBuf - 16))
			bSIMD.Offs[k] = off
			bScalar.Offs[k] = off
		}

		C2quic_vint_decode_batch8(buf, totalBuf, &bSIMD)
		C2quic_vint_decode_batch8_scalar(buf, totalBuf, &bScalar)

		for k := 0; k < 8; k++ {
			if bSIMD.Vals[k] != bScalar.Vals[k] || bSIMD.Nbytes[k] != bScalar.Nbytes[k] || bSIMD.Status[k] != bScalar.Status[k] {
				t.Fatalf("Divergence batch8 iter=%d voie=%d: SIMD(val=%d, n=%d, st=%d) vs Scalar(val=%d, n=%d, st=%d)",
					iter, k, bSIMD.Vals[k], bSIMD.Nbytes[k], bSIMD.Status[k], bScalar.Vals[k], bScalar.Nbytes[k], bScalar.Status[k])
			}
		}
	}
}

func TestZeroAlloc_Batch8SIMD(t *testing.T) {
	buf := make([]byte, 256)
	for i := range buf {
		buf[i] = byte(i)
	}
	var b C2quic_batch8_t
	for k := 0; k < 8; k++ {
		b.Offs[k] = uint32(k * 16)
	}
	allocs := testing.AllocsPerRun(1000, func() {
		C2quic_vint_decode_batch8(buf, 256, &b)
	})
	if allocs != 0 {
		t.Fatalf("C2quic_vint_decode_batch8 allocs/op = %.2f, attendu 0", allocs)
	}
}

