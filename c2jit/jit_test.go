// SPDX-License-Identifier: Apache-2.0 OR MIT
package c2jit_test

import (
	"bytes"
	"runtime"
	"testing"
	"unsafe"

	c2jit "github.com/hazyhaar/c2pkg/c2jit"
	c2oracle "github.com/hazyhaar/c2pkg/c2oracle"
)

func TestAssembler_SanctuaryR14(t *testing.T) {
	asm := c2jit.NewAssembler(64)
	asm.MovRR(c2jit.R14, c2jit.RAX)
	if asm.Err() != c2jit.ErrSanctuaryRegisterViolation {
		t.Fatalf("Attendu ErrSanctuaryRegisterViolation lors de l'écriture sur R14, obtenu: %v", asm.Err())
	}
}

func TestAssembler_OpcodesBitExact(t *testing.T) {
	asm := c2jit.NewAssembler(128)

	// MOV RAX, RDI : 48 89 f8
	asm.MovRR(c2jit.RAX, c2jit.RDI)
	// ADD RAX, RSI : 48 01 f0
	asm.AddRR(c2jit.RAX, c2jit.RSI)
	// SUB RAX, RDX : 48 29 d0
	asm.SubRR(c2jit.RAX, c2jit.RDX)
	// XOR RAX, RAX : 48 31 c0
	asm.XorRR(c2jit.RAX, c2jit.RAX)
	// RET : c3
	asm.Ret()

	expected := []byte{
		0x48, 0x89, 0xF8,
		0x48, 0x01, 0xF0,
		0x48, 0x29, 0xD0,
		0x48, 0x31, 0xC0,
		0xC3,
	}

	if !bytes.Equal(asm.Bytes(), expected) {
		t.Fatalf("Opcode mismatch:\nGot:      %X\nExpected: %X", asm.Bytes(), expected)
	}

	// Contrôle négatif obligatoire (anti-tautologie)
	c2oracle.AssertNegativeControl(t, func(mutated []byte) bool {
		return bytes.Equal(asm.Bytes(), mutated)
	}, expected)
}

func TestAssembler_AVX2OpcodesBitExact(t *testing.T) {
	asm := c2jit.NewAssembler(128)

	// VZEROUPPER : c5 f8 77
	asm.Vzeroupper()
	// VMOVDQU YMM0, [RSI] : c4 e1 7e 6f 06
	asm.VmovdquLoad(c2jit.YMM0, c2jit.RSI, 0)
	// VPXOR YMM2, YMM0, YMM1 : c4 e1 7d ef d1
	asm.Vpxor(c2jit.YMM2, c2jit.YMM0, c2jit.YMM1)
	// VMOVDQU [RDI], YMM2 : c4 e1 7e 7f 17
	asm.VmovdquStore(c2jit.RDI, 0, c2jit.YMM2)
	// RET : c3
	asm.Ret()

	expected := []byte{
		0xC5, 0xF8, 0x77,
		0xC4, 0xE1, 0x7E, 0x6F, 0x06,
		0xC4, 0xE1, 0x7D, 0xEF, 0xD1,
		0xC4, 0xE1, 0x7E, 0x7F, 0x17,
		0xC3,
	}

	if !bytes.Equal(asm.Bytes(), expected) {
		t.Fatalf("AVX2 Opcode mismatch:\nGot:      %X\nExpected: %X", asm.Bytes(), expected)
	}
}

func TestJIT_Execution_ScalarAdd(t *testing.T) {
	if runtime.GOARCH != "amd64" {
		t.Skip("Test d'exécution JIT amd64 natif ignoré sur autre architecture")
	}

	asm := c2jit.NewAssembler(64)
	// Go internal ABI (amd64) : arg0=RAX, arg1=RBX, ret=RAX
	asm.AddRR(c2jit.RAX, c2jit.RBX)
	asm.Ret()

	if err := asm.Err(); err != nil {
		t.Fatalf("Erreur assemblage: %v", err)
	}

	buf, err := c2jit.AllocateBuffer(asm.Len())
	if err != nil {
		t.Fatalf("AllocateBuffer: %v", err)
	}
	defer buf.Free()

	if err := buf.Write(asm.Bytes()); err != nil {
		t.Fatalf("Write: %v", err)
	}

	if err := buf.SealExecutable(); err != nil {
		t.Fatalf("SealExecutable: %v", err)
	}

	fn := c2jit.AsFunc2(buf.Ptr())
	res := fn(40, 2)
	if res != 42 {
		t.Fatalf("Résultat JIT incorrect: attendu 42, obtenu %d", res)
	}

	// Test sur plusieurs valeurs
	testCases := [][3]int64{
		{100, 200, 300},
		{0, 0, 0},
		{-50, 100, 50},
		{123456789, 987654321, 1111111110},
	}

	for _, tc := range testCases {
		got := int64(fn(uintptr(tc[0]), uintptr(tc[1])))
		if got != tc[2] {
			t.Errorf("Add(%d, %d): attendu %d, obtenu %d", tc[0], tc[1], tc[2], got)
		}
	}
}

func TestJIT_Execution_VectorAVX2_XOR(t *testing.T) {
	if runtime.GOARCH != "amd64" {
		t.Skip("Test d'exécution JIT AVX2 amd64 ignoré sur autre architecture")
	}

	asm := c2jit.NewAssembler(128)
	// Go internal ABI (amd64) :
	// arg0 (RAX) = dst (*[8]uint32)
	// arg1 (RBX) = src1 (*[8]uint32)
	// arg2 (RCX) = src2 (*[8]uint32)
	asm.VmovdquLoad(c2jit.YMM0, c2jit.RBX, 0)
	asm.VmovdquLoad(c2jit.YMM1, c2jit.RCX, 0)
	asm.Vpxor(c2jit.YMM2, c2jit.YMM0, c2jit.YMM1)
	asm.VmovdquStore(c2jit.RAX, 0, c2jit.YMM2)
	asm.Vzeroupper()
	asm.Ret()

	if err := asm.Err(); err != nil {
		t.Fatalf("Erreur assemblage AVX2: %v", err)
	}

	buf, err := c2jit.AllocateBuffer(asm.Len())
	if err != nil {
		t.Fatalf("AllocateBuffer: %v", err)
	}
	defer buf.Free()

	if err := buf.Write(asm.Bytes()); err != nil {
		t.Fatalf("Write: %v", err)
	}

	if err := buf.SealExecutable(); err != nil {
		t.Fatalf("SealExecutable: %v", err)
	}

	fn := c2jit.AsFunc3(buf.Ptr())

	src1 := [8]uint32{1, 2, 3, 4, 5, 6, 7, 8}
	src2 := [8]uint32{0xFF, 0xEE, 0xDD, 0xCC, 0xBB, 0xAA, 0x99, 0x88}
	var dst [8]uint32

	fn(uintptr(unsafe.Pointer(&dst[0])), uintptr(unsafe.Pointer(&src1[0])), uintptr(unsafe.Pointer(&src2[0])))

	for i := 0; i < 8; i++ {
		expected := src1[i] ^ src2[i]
		if dst[i] != expected {
			t.Errorf("dst[%d] = 0x%X (attendu 0x%X)", i, dst[i], expected)
		}
	}
}

func BenchmarkJIT_ScalarAdd_Throughput(b *testing.B) {
	if runtime.GOARCH != "amd64" {
		b.Skip("amd64 only")
	}

	asm := c2jit.NewAssembler(64)
	asm.MovRR(c2jit.RAX, c2jit.RDI)
	asm.AddRR(c2jit.RAX, c2jit.RSI)
	asm.Ret()

	buf, err := c2jit.AllocateBuffer(asm.Len())
	if err != nil {
		b.Fatalf("AllocateBuffer: %v", err)
	}
	defer buf.Free()

	_ = buf.Write(asm.Bytes())
	_ = buf.SealExecutable()

	fn := c2jit.AsFunc2(buf.Ptr())

	b.ReportAllocs()
	b.ResetTimer()

	var acc uintptr
	for i := 0; i < b.N; i++ {
		acc = fn(acc, 1)
	}
	_ = acc
}
