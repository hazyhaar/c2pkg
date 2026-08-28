//go:build !(goexperiment.simd && amd64)

package c2fused

import (
	"github.com/hazyhaar/c2pkg/c2chacha8"
	"github.com/hazyhaar/c2pkg/c2poly1305"
)

type C2fused_poly struct {
	ctx c2poly1305.Crypto_poly1305_ctx
}

func C2fused_poly_init(st *C2fused_poly, key []byte) {
	c2poly1305.Crypto_poly1305_init(&st.ctx, key)
}

func C2fused_poly_update(st *C2fused_poly, m []byte, bytes uint64) {
	if bytes > uint64(len(m)) {
		bytes = uint64(len(m))
	}
	c2poly1305.Crypto_poly1305_update(&st.ctx, m[:bytes], bytes)
}

func C2fused_poly_blocks(st *C2fused_poly, m []byte, blocks uint64) {
	bytes := blocks * 16
	if bytes > uint64(len(m)) {
		bytes = uint64(len(m))
	}
	c2poly1305.Crypto_poly1305_update(&st.ctx, m[:bytes], bytes)
}

func C2fused_poly_final(st *C2fused_poly, mac []byte) {
	c2poly1305.Crypto_poly1305_final(&st.ctx, mac[:16])
}

func C2fused_seal_blocks(out []byte, in []byte, n uint64, key []byte, nonce []byte, counter uint32, st *C2fused_poly, prev []byte) {
	if prev != nil {
		C2fused_poly_blocks(st, prev, uint64(len(prev)/16))
	}
	c2chacha8.C2chacha8_xor_blocks(out, in, n, key, nonce, counter)
}
