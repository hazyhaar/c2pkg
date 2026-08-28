//go:build !(goexperiment.simd && amd64)

package c2chacha8

import (
	"encoding/binary"
	"math/bits"
)

func C2chacha8_xor_blocks(out []byte, in []byte, n uint64, key []byte, nonce []byte, counter uint32) uint64 {
	_ = key[31]
	_ = nonce[11]
	if n > 512 {
		n = 512
	}
	var state [16]uint32
	state[0] = 0x61707865
	state[1] = 0x3320646e
	state[2] = 0x79622d32
	state[3] = 0x6b206574
	state[4] = binary.LittleEndian.Uint32(key[0:4])
	state[5] = binary.LittleEndian.Uint32(key[4:8])
	state[6] = binary.LittleEndian.Uint32(key[8:12])
	state[7] = binary.LittleEndian.Uint32(key[12:16])
	state[8] = binary.LittleEndian.Uint32(key[16:20])
	state[9] = binary.LittleEndian.Uint32(key[20:24])
	state[10] = binary.LittleEndian.Uint32(key[24:28])
	state[11] = binary.LittleEndian.Uint32(key[28:32])
	state[13] = binary.LittleEndian.Uint32(nonce[0:4])
	state[14] = binary.LittleEndian.Uint32(nonce[4:8])
	state[15] = binary.LittleEndian.Uint32(nonce[8:12])

	var processed uint64
	for processed < n {
		state[12] = counter
		var x = state
		for i := 0; i < 10; i++ {
			// Column round
			x[0] += x[4]; x[12] = bits.RotateLeft32(x[12]^x[0], 16)
			x[8] += x[12]; x[4] = bits.RotateLeft32(x[4]^x[8], 12)
			x[0] += x[4]; x[12] = bits.RotateLeft32(x[12]^x[0], 8)
			x[8] += x[12]; x[4] = bits.RotateLeft32(x[4]^x[8], 7)

			x[1] += x[5]; x[13] = bits.RotateLeft32(x[13]^x[1], 16)
			x[9] += x[13]; x[5] = bits.RotateLeft32(x[5]^x[9], 12)
			x[1] += x[5]; x[13] = bits.RotateLeft32(x[13]^x[1], 8)
			x[9] += x[13]; x[5] = bits.RotateLeft32(x[5]^x[9], 7)

			x[2] += x[6]; x[14] = bits.RotateLeft32(x[14]^x[2], 16)
			x[10] += x[14]; x[6] = bits.RotateLeft32(x[6]^x[10], 12)
			x[2] += x[6]; x[14] = bits.RotateLeft32(x[14]^x[2], 8)
			x[10] += x[14]; x[6] = bits.RotateLeft32(x[6]^x[10], 7)

			x[3] += x[7]; x[15] = bits.RotateLeft32(x[15]^x[3], 16)
			x[11] += x[15]; x[7] = bits.RotateLeft32(x[7]^x[11], 12)
			x[3] += x[7]; x[15] = bits.RotateLeft32(x[15]^x[3], 8)
			x[11] += x[15]; x[7] = bits.RotateLeft32(x[7]^x[11], 7)

			// Diagonal round
			x[0] += x[5]; x[15] = bits.RotateLeft32(x[15]^x[0], 16)
			x[10] += x[15]; x[5] = bits.RotateLeft32(x[5]^x[10], 12)
			x[0] += x[5]; x[15] = bits.RotateLeft32(x[15]^x[0], 8)
			x[10] += x[15]; x[5] = bits.RotateLeft32(x[5]^x[10], 7)

			x[1] += x[6]; x[12] = bits.RotateLeft32(x[12]^x[1], 16)
			x[11] += x[12]; x[6] = bits.RotateLeft32(x[6]^x[11], 12)
			x[1] += x[6]; x[12] = bits.RotateLeft32(x[12]^x[1], 8)
			x[11] += x[12]; x[6] = bits.RotateLeft32(x[6]^x[11], 7)

			x[2] += x[7]; x[13] = bits.RotateLeft32(x[13]^x[2], 16)
			x[8] += x[13]; x[7] = bits.RotateLeft32(x[7]^x[8], 12)
			x[2] += x[7]; x[13] = bits.RotateLeft32(x[13]^x[2], 8)
			x[8] += x[13]; x[7] = bits.RotateLeft32(x[7]^x[8], 7)

			x[3] += x[4]; x[14] = bits.RotateLeft32(x[14]^x[3], 16)
			x[9] += x[14]; x[4] = bits.RotateLeft32(x[4]^x[9], 12)
			x[3] += x[4]; x[14] = bits.RotateLeft32(x[14]^x[3], 8)
			x[9] += x[14]; x[4] = bits.RotateLeft32(x[4]^x[9], 7)
		}

		var block [64]byte
		for i := 0; i < 16; i++ {
			binary.LittleEndian.PutUint32(block[i*4:], x[i]+state[i])
		}

		chunk := n - processed
		if chunk > 64 {
			chunk = 64
		}
		for i := uint64(0); i < chunk; i++ {
			out[processed+i] = in[processed+i] ^ block[i]
		}
		processed += chunk
		counter++
	}
	return processed
}
