//go:build !(goexperiment.simd && amd64)

package c2chacha4

import "github.com/hazyhaar/c2pkg/c2chacha8"

func C2chacha4_xor_blocks(out []byte, in []byte, n uint64, key []byte, nonce []byte, counter uint32) uint64 {
	if n > 512 {
		n = 512
	}
	return c2chacha8.C2chacha8_xor_blocks(out, in, n, key, nonce, counter)
}
