// SPDX-License-Identifier: Apache-2.0 OR MIT
package c2jit

import (
	"encoding/binary"
	"fmt"
)

// Assembler est le générateur linéaire d'instructions machine x86-64 / AVX2.
type Assembler struct {
	buf []byte
	err error
}

// NewAssembler instancie un nouvel assembleur avec une capacité initiale réservée.
func NewAssembler(capHint int) *Assembler {
	if capHint <= 0 {
		capHint = 256
	}
	return &Assembler{
		buf: make([]byte, 0, capHint),
	}
}

// Err retourne la première erreur survenue durant l'assemblage (ex. altération de R14).
func (a *Assembler) Err() error {
	return a.err
}

// Bytes retourne les octets machines générés.
func (a *Assembler) Bytes() []byte {
	return a.buf
}

// Len retourne la taille en octets du code généré.
func (a *Assembler) Len() int {
	return len(a.buf)
}

// Reset réinitialise le tampon d'assemblage sans désallouer la mémoire sous-jacente.
func (a *Assembler) Reset() {
	a.buf = a.buf[:0]
	a.err = nil
}

// EmitBytes écrit une séquence brute d'octets.
func (a *Assembler) EmitBytes(b ...byte) {
	a.buf = append(a.buf, b...)
}

func (a *Assembler) checkSanctuary(r Reg) bool {
	if r == R14 {
		if a.err == nil {
			a.err = ErrSanctuaryRegisterViolation
		}
		return false
	}
	return true
}

// EmitREX écrit un préfixe REX 64-bit adapté aux registres étendus (R8..R15).
func (a *Assembler) emitREXW(r, b Reg) {
	rex := EncodeREX(true, r >= 8, false, b >= 8)
	a.buf = append(a.buf, rex)
}

// MovRR émet : MOV dst, src (64-bit register-to-register)
func (a *Assembler) MovRR(dst, src Reg) {
	if !a.checkSanctuary(dst) {
		return
	}
	a.emitREXW(src, dst)
	a.buf = append(a.buf, 0x89, EncodeModRM(3, uint8(src)&7, uint8(dst)&7))
}

// MovRI émet : MOV dst, imm64 (chargement d'immédiat 64-bit)
func (a *Assembler) MovRI(dst Reg, imm int64) {
	if !a.checkSanctuary(dst) {
		return
	}
	rex := EncodeREX(true, false, false, dst >= 8)
	opcode := 0xB8 + (uint8(dst) & 7)
	a.buf = append(a.buf, rex, opcode)
	var immBuf [8]byte
	binary.LittleEndian.PutUint64(immBuf[:], uint64(imm))
	a.buf = append(a.buf, immBuf[:]...)
}

// AddRR émet : ADD dst, src (64-bit)
func (a *Assembler) AddRR(dst, src Reg) {
	if !a.checkSanctuary(dst) {
		return
	}
	a.emitREXW(src, dst)
	a.buf = append(a.buf, 0x01, EncodeModRM(3, uint8(src)&7, uint8(dst)&7))
}

// SubRR émet : SUB dst, src (64-bit)
func (a *Assembler) SubRR(dst, src Reg) {
	if !a.checkSanctuary(dst) {
		return
	}
	a.emitREXW(src, dst)
	a.buf = append(a.buf, 0x29, EncodeModRM(3, uint8(src)&7, uint8(dst)&7))
}

// AndRR émet : AND dst, src (64-bit)
func (a *Assembler) AndRR(dst, src Reg) {
	if !a.checkSanctuary(dst) {
		return
	}
	a.emitREXW(src, dst)
	a.buf = append(a.buf, 0x21, EncodeModRM(3, uint8(src)&7, uint8(dst)&7))
}

// OrRR émet : OR dst, src (64-bit)
func (a *Assembler) OrRR(dst, src Reg) {
	if !a.checkSanctuary(dst) {
		return
	}
	a.emitREXW(src, dst)
	a.buf = append(a.buf, 0x09, EncodeModRM(3, uint8(src)&7, uint8(dst)&7))
}

// XorRR émet : XOR dst, src (64-bit)
func (a *Assembler) XorRR(dst, src Reg) {
	if !a.checkSanctuary(dst) {
		return
	}
	a.emitREXW(src, dst)
	a.buf = append(a.buf, 0x31, EncodeModRM(3, uint8(src)&7, uint8(dst)&7))
}

// ImulRR émet : IMUL dst, src (64-bit signed multiply)
func (a *Assembler) ImulRR(dst, src Reg) {
	if !a.checkSanctuary(dst) {
		return
	}
	a.emitREXW(dst, src)
	a.buf = append(a.buf, 0x0F, 0xAF, EncodeModRM(3, uint8(dst)&7, uint8(src)&7))
}

// CmpRR émet : CMP a, b (64-bit)
func (a *Assembler) CmpRR(aReg, bReg Reg) {
	a.emitREXW(bReg, aReg)
	a.buf = append(a.buf, 0x39, EncodeModRM(3, uint8(bReg)&7, uint8(aReg)&7))
}

// TestRR émet : TEST a, b (64-bit)
func (a *Assembler) TestRR(aReg, bReg Reg) {
	a.emitREXW(bReg, aReg)
	a.buf = append(a.buf, 0x85, EncodeModRM(3, uint8(bReg)&7, uint8(aReg)&7))
}

// Ret émet : RET (0xC3)
func (a *Assembler) Ret() {
	a.buf = append(a.buf, 0xC3)
}

// Vzeroupper émet : VZEROUPPER (0xC5 0xF8 0x77)
func (a *Assembler) Vzeroupper() {
	a.buf = append(a.buf, 0xC5, 0xF8, 0x77)
}

func (a *Assembler) emitMem(regCode uint8, base Reg, disp int32) {
	mod := uint8(2) // disp32
	if disp == 0 && (base&7) != 5 {
		mod = 0 // pas de déplacement (sauf RBP/R13 qui exigent disp8 0)
	} else if disp >= -128 && disp <= 127 {
		mod = 1 // disp8
	}

	rm := uint8(base) & 7
	// RSP / R12 exigent un octet SIB
	needsSIB := (rm == 4)

	if needsSIB {
		a.buf = append(a.buf, EncodeModRM(mod, regCode, 4))
		a.buf = append(a.buf, EncodeSIB(0, 4, rm))
	} else {
		a.buf = append(a.buf, EncodeModRM(mod, regCode, rm))
	}

	if mod == 1 {
		a.buf = append(a.buf, byte(int8(disp)))
	} else if mod == 2 || (mod == 0 && (base&7) == 5) {
		var d [4]byte
		binary.LittleEndian.PutUint32(d[:], uint32(disp))
		a.buf = append(a.buf, d[:]...)
	}
}

// VmovdquLoad émet : VMOVDQU ymm_dst, [base + disp] (chargement non aligné 256-bit AVX2)
func (a *Assembler) VmovdquLoad(dst YMMReg, base Reg, disp int32) {
	// Prefix : VEX.256.F3.0F.WIG 0x6F
	vex := VEXPrefix3B(dst >= 8, false, base >= 8, 1, false, 0, true, 2)
	a.buf = append(a.buf, vex[0], vex[1], vex[2], 0x6F)
	a.emitMem(uint8(dst)&7, base, disp)
}

// VmovdquStore émet : VMOVDQU [base + disp], ymm_src (écriture non alignée 256-bit AVX2)
func (a *Assembler) VmovdquStore(base Reg, disp int32, src YMMReg) {
	// Prefix : VEX.256.F3.0F.WIG 0x7F
	vex := VEXPrefix3B(src >= 8, false, base >= 8, 1, false, 0, true, 2)
	a.buf = append(a.buf, vex[0], vex[1], vex[2], 0x7F)
	a.emitMem(uint8(src)&7, base, disp)
}

// Vpaddd émet : VPADDD dst, src1, src2 (addition vectorielle 8x int32 AVX2)
func (a *Assembler) Vpaddd(dst, src1, src2 YMMReg) {
	// Prefix : VEX.256.66.0F.WIG 0xFE /r
	vex := VEXPrefix3B(dst >= 8, false, src2 >= 8, 1, false, src1, true, 1)
	a.buf = append(a.buf, vex[0], vex[1], vex[2], 0xFE, EncodeModRM(3, uint8(dst)&7, uint8(src2)&7))
}

// Vpsubd émet : VPSUBD dst, src1, src2 (soustraction vectorielle 8x int32 AVX2)
func (a *Assembler) Vpsubd(dst, src1, src2 YMMReg) {
	// Prefix : VEX.256.66.0F.WIG 0xFA /r
	vex := VEXPrefix3B(dst >= 8, false, src2 >= 8, 1, false, src1, true, 1)
	a.buf = append(a.buf, vex[0], vex[1], vex[2], 0xFA, EncodeModRM(3, uint8(dst)&7, uint8(src2)&7))
}

// Vpxor émet : VPXOR dst, src1, src2 (XOR vectoriel 256-bit AVX2)
func (a *Assembler) Vpxor(dst, src1, src2 YMMReg) {
	// Prefix : VEX.256.66.0F.WIG 0xEF /r
	vex := VEXPrefix3B(dst >= 8, false, src2 >= 8, 1, false, src1, true, 1)
	a.buf = append(a.buf, vex[0], vex[1], vex[2], 0xEF, EncodeModRM(3, uint8(dst)&7, uint8(src2)&7))
}

// Vpand émet : VPAND dst, src1, src2 (ET vectoriel 256-bit AVX2)
func (a *Assembler) Vpand(dst, src1, src2 YMMReg) {
	// Prefix : VEX.256.66.0F.WIG 0xDB /r
	vex := VEXPrefix3B(dst >= 8, false, src2 >= 8, 1, false, src1, true, 1)
	a.buf = append(a.buf, vex[0], vex[1], vex[2], 0xDB, EncodeModRM(3, uint8(dst)&7, uint8(src2)&7))
}

// Vpor émet : VPOR dst, src1, src2 (OU vectoriel 256-bit AVX2)
func (a *Assembler) Vpor(dst, src1, src2 YMMReg) {
	// Prefix : VEX.256.66.0F.WIG 0xEB /r
	vex := VEXPrefix3B(dst >= 8, false, src2 >= 8, 1, false, src1, true, 1)
	a.buf = append(a.buf, vex[0], vex[1], vex[2], 0xEB, EncodeModRM(3, uint8(dst)&7, uint8(src2)&7))
}

// Vpaddq émet : VPADDQ dst, src1, src2 (addition vectorielle 4x int64 AVX2)
func (a *Assembler) Vpaddq(dst, src1, src2 YMMReg) {
	// Prefix : VEX.256.66.0F.WIG 0xD4 /r
	vex := VEXPrefix3B(dst >= 8, false, src2 >= 8, 1, false, src1, true, 1)
	a.buf = append(a.buf, vex[0], vex[1], vex[2], 0xD4, EncodeModRM(3, uint8(dst)&7, uint8(src2)&7))
}

// JmpRel8 émet : JMP disp8 (saut relatif court)
func (a *Assembler) JmpRel8(disp int8) {
	a.buf = append(a.buf, 0xEB, byte(disp))
}

// JccRel8 émet : Jcc disp8 (saut conditionnel court)
func (a *Assembler) JccRel8(cond Condition, disp int8) {
	a.buf = append(a.buf, 0x70|byte(cond), byte(disp))
}

// DisasmSummary fournit un résumé textuel de la taille et de la signature du bloc émis.
func (a *Assembler) DisasmSummary() string {
	return fmt.Sprintf("JIT Block [%d bytes, err=%v]", len(a.buf), a.err)
}
