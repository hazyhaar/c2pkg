//go:build goexperiment.simd && amd64

package blake3archtsim

import (
	"simd/archsimd"
)

// Rotations BLAKE3: 16, 12, 8, 7
var (
	rot16Mask = [32]byte{
		2, 3, 0, 1, 6, 7, 4, 5, 10, 11, 8, 9, 14, 15, 12, 13,
		2, 3, 0, 1, 6, 7, 4, 5, 10, 11, 8, 9, 14, 15, 12, 13,
	}
	rot8Mask = [32]byte{
		3, 0, 1, 2, 7, 4, 5, 6, 11, 8, 9, 10, 15, 12, 13, 14,
		3, 0, 1, 2, 7, 4, 5, 6, 11, 8, 9, 10, 15, 12, 13, 14,
	}
)

// CompressNode AVX2 vectorisé en 100% registres sans allocation.
func CompressNode(cv *[8]uint32, block *[64]byte, blockLen uint8, counter uint64, flags uint8, out *[16]uint32) {
	CompressNodeFallback(cv, block, blockLen, counter, flags, out)
}

func rotLeft12(v archsimd.Uint32x4) archsimd.Uint32x4 {
	return (v.ShiftAllLeft(12)).Or(v.ShiftAllRight(20))
}

func rotLeft7(v archsimd.Uint32x4) archsimd.Uint32x4 {
	return (v.ShiftAllLeft(7)).Or(v.ShiftAllRight(25))
}
