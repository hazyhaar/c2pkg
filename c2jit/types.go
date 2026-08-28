// SPDX-License-Identifier: Apache-2.0 OR MIT
// Package c2jit implémente un micro-assembleur machine dynamique AVX2 / x86-64 en pur Go
// sous cycle de vie de sécurité mémoire W^X strict (zéro CGO, zéro allocation).
package c2jit

import (
	"errors"
	"fmt"
)

var (
	// ErrSanctuaryRegisterViolation est retourné si une instruction tente d'altérer R14 (goroutine pointer 'g').
	ErrSanctuaryRegisterViolation = errors.New("c2jit: tentative d'écriture interdite sur le registre sanctuarisé R14 (g pointer)")
	// ErrBufferTooSmall est retourné si le tampon exécutable alloué est insuffisant.
	ErrBufferTooSmall = errors.New("c2jit: tampon d'émission machine insuffisant")
)

// Reg identifie un registre général 64-bit x86-64.
type Reg uint8

const (
	RAX Reg = 0
	RCX Reg = 1
	RDX Reg = 2
	RBX Reg = 3
	RSP Reg = 4
	RBP Reg = 5
	RSI Reg = 6
	RDI Reg = 7
	R8  Reg = 8
	R9  Reg = 9
	R10 Reg = 10
	R11 Reg = 11
	R12 Reg = 12
	R13 Reg = 13
	R14 Reg = 14 // SANCTUARISÉ : g pointer runtime Go
	R15 Reg = 15
)

// YMMReg identifie un registre vectoriel AVX2 256-bit (YMM0..YMM15).
type YMMReg uint8

const (
	YMM0 YMMReg = iota
	YMM1
	YMM2
	YMM3
	YMM4
	YMM5
	YMM6
	YMM7
	YMM8
	YMM9
	YMM10
	YMM11
	YMM12
	YMM13
	YMM14
	YMM15
)

// Condition identifie une condition de saut conditionnel x86-64 (Jcc).
type Condition uint8

const (
	CondOverflow Condition = 0x0 // JO
	CondNoOverflow Condition = 0x1 // JNO
	CondBelow Condition = 0x2 // JB / JC / JNAE
	CondAboveEqual Condition = 0x3 // JAE / JNC / JNB
	CondEqual Condition = 0x4 // JE / JZ
	CondNotEqual Condition = 0x5 // JNE / JNZ
	CondBelowEqual Condition = 0x6 // JBE / JNA
	CondAbove Condition = 0x7 // JA / JNBE
	CondSign Condition = 0x8 // JS
	CondNoSign Condition = 0x9 // JNS
	CondParity Condition = 0xA // JP / JPE
	CondNoParity Condition = 0xB // JNP / JPO
	CondLess Condition = 0xC // JL / JNGE
	CondGreaterEqual Condition = 0xD // JGE / JNL
	CondLessEqual Condition = 0xE // JLE / JNG
	CondGreater Condition = 0xF // JG / JNLE
)

// EncodeModRM calcule l'octet ModR/M x86-64 : (mod << 6) | (reg << 3) | rm
func EncodeModRM(mod uint8, reg uint8, rm uint8) byte {
	return ((mod & 0x3) << 6) | ((reg & 0x7) << 3) | (rm & 0x7)
}

// EncodeSIB calcule l'octet SIB x86-64 : (scale << 6) | (index << 3) | base
func EncodeSIB(scale uint8, index uint8, base uint8) byte {
	return ((scale & 0x3) << 6) | ((index & 0x7) << 3) | (base & 0x7)
}

// EncodeREX calcule le préfixe REX x86-64 : 0100 W R X B
func EncodeREX(w, r, x, b bool) byte {
	var rex byte = 0x40
	if w {
		rex |= 0x08
	}
	if r {
		rex |= 0x04
	}
	if x {
		rex |= 0x02
	}
	if b {
		rex |= 0x01
	}
	return rex
}

// VEXPrefix2B produit le préfixe VEX 2 octets (0xC5) pour les opcodes dans la table 0F avec W=0 et pas d'index X/B étendu.
// Format : 0xC5 | (~R & 1) << 7 | (~vvvv & 0xF) << 3 | (L & 1) << 2 | (pp & 3)
func VEXPrefix2B(r bool, vvvv YMMReg, l256 bool, pp uint8) [2]byte {
	var rBit byte = 1
	if r {
		rBit = 0
	}
	var lBit byte = 0
	if l256 {
		lBit = 1
	}
	vField := (15 - (uint8(vvvv) & 0xF)) & 0xF
	b2 := (rBit << 7) | (vField << 3) | (lBit << 2) | (pp & 0x3)
	return [2]byte{0xC5, b2}
}

// VEXPrefix3B produit le préfixe VEX 3 octets (0xC4) complet.
// Format : 0xC4 | (~R << 7) | (~X << 6) | (~B << 5) | (map & 0x1F)
//          | (W << 7) | (~vvvv << 3) | (L << 2) | (pp & 3)
func VEXPrefix3B(r, x, b bool, mmmm uint8, w bool, vvvv YMMReg, l256 bool, pp uint8) [3]byte {
	var rBit, xBit, bBit byte = 1, 1, 1
	if r {
		rBit = 0
	}
	if x {
		xBit = 0
	}
	if b {
		bBit = 0
	}
	b2 := (rBit << 7) | (xBit << 6) | (bBit << 5) | (mmmm & 0x1F)

	var wBit, lBit byte = 0, 0
	if w {
		wBit = 1
	}
	if l256 {
		lBit = 1
	}
	vField := (15 - (uint8(vvvv) & 0xF)) & 0xF
	b3 := (wBit << 7) | (vField << 3) | (lBit << 2) | (pp & 0x3)

	return [3]byte{0xC4, b2, b3}
}

func (r Reg) String() string {
	names := [...]string{"RAX", "RCX", "RDX", "RBX", "RSP", "RBP", "RSI", "RDI", "R8", "R9", "R10", "R11", "R12", "R13", "R14(g)", "R15"}
	if int(r) < len(names) {
		return names[r]
	}
	return fmt.Sprintf("Reg(%d)", r)
}
