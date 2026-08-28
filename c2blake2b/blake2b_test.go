package c2blake2b

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestBlake2bKeyed(t *testing.T) {
	shortKey := []byte("cle-courte-16-oc")
	msg := []byte("message")
	var ctx Crypto_blake2b_ctx
	Crypto_blake2b_keyed_init(&ctx, 32, shortKey, uint64(len(shortKey)))
	Crypto_blake2b_update(&ctx, msg, uint64(len(msg)))
	var out [32]byte
	Crypto_blake2b_final(&ctx, out[:])

	want, _ := hex.DecodeString("75513d30293425730ed8e58ee474ead4f12c26f2e2755dbe0e9ebe74bdf29051")
	if !bytes.Equal(out[:], want) {
		t.Fatalf("condensat keyed divergent de la référence: got %x want %x", out, want)
	}
}

func TestBlake2bZeroAlloc(t *testing.T) {
	data := make([]byte, 1024)
	var out [64]byte
	allocs := testing.AllocsPerRun(50, func() {
		var ctx Crypto_blake2b_ctx
		Crypto_blake2b_init(&ctx, 64)
		Crypto_blake2b_update(&ctx, data, uint64(len(data)))
		Crypto_blake2b_final(&ctx, out[:])
	})
	if allocs > 0 {
		t.Fatalf("BLAKE2b viole l'invariant 0-allocation: allocs/op = %.2f", allocs)
	}
}
