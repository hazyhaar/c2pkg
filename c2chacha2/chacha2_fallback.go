//go:build !(goexperiment.simd && amd64)

package c2chacha2

import "github.com/hazyhaar/c2pkg/c2chacha8"

func C2chacha2_xor_blocks(out []byte, in []byte, n uint64, key []byte, nonce []byte, counter uint32) uint64 {
	if n > 128 {
		n = 128
	}
	return c2chacha8.C2chacha8_xor_blocks(out, in, n, key, nonce, counter)
}
