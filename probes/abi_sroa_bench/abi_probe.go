package main

import (
	"math/bits"
)

type U128Struct struct {
	Lo uint64
	Hi uint64
}

// 1. Array parameter: forces stack spill in Go 1.27 due to types/size.go:619 returning MaxUint8 for [N>1]T.
// Generates 8 memory access instructions (MOVQ / MOVUPS).
//go:noinline
func Add128_Array(a, b [2]uint64) [2]uint64 {
	lo, c := bits.Add64(a[0], b[0], 0)
	hi, _ := bits.Add64(a[1], b[1], c)
	return [2]uint64{lo, hi}
}

// 2. Struct parameter: passes 100% in register ABI (RAX, RBX, RCX, RDX) since fields <= MaxStruct.
// Generates 0 memory access instructions (pure ADDQ + ADCQ).
//go:noinline
func Add128_Struct(a, b U128Struct) U128Struct {
	lo, c := bits.Add64(a.Lo, b.Lo, 0)
	hi, _ := bits.Add64(a.Hi, b.Hi, c)
	return U128Struct{Lo: lo, Hi: hi}
}

// 3. Discrete scalars parameter: passes 100% in registers.
// Generates 0 memory access instructions.
//go:noinline
func Add128_Scalars(aLo, aHi, bLo, bHi uint64) (uint64, uint64) {
	lo, c := bits.Add64(aLo, bLo, 0)
	hi, _ := bits.Add64(aHi, bHi, c)
	return lo, hi
}

// 4. Bitwise XOR on array: stack spill.
//go:noinline
func Xor128_Array(a, b [2]uint64) [2]uint64 {
	return [2]uint64{a[0] ^ b[0], a[1] ^ b[1]}
}

// 5. Bitwise XOR on struct: 100% register ABI.
//go:noinline
func Xor128_Struct(a, b U128Struct) U128Struct {
	return U128Struct{Lo: a.Lo ^ b.Lo, Hi: a.Hi ^ b.Hi}
}

func main() {}


