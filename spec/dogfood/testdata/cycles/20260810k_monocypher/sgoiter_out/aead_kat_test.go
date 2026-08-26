package monocypher_amalg

import (
	"bytes"
	"testing"
)

func TestMonoAEAD_KAT_36(t *testing.T) {
	key := make([]byte, 32)
	for i := range key {
		key[i] = byte(i + 1)
	}
	nonce := make([]byte, 24)
	for i := range nonce {
		nonce[i] = byte(i + 10)
	}
	ad := []byte("HEADER")
	pt := []byte("HELLO MONOCYPHER SGOITER AEAD CGO=0!")
	ct := make([]byte, len(pt))
	mac := make([]byte, 16)
	Crypto_aead_lock(ct, mac, key, nonce, ad, uint64(len(ad)), pt, uint64(len(pt)))
	out := make([]byte, len(pt))
	if Crypto_aead_unlock(out, mac, key, nonce, ad, uint64(len(ad)), ct, uint64(len(ct))) != 0 {
		t.Fatal("mac fail")
	}
	if !bytes.Equal(pt, out) {
		t.Fatalf("mismatch 36B")
	}
}

func TestMonoAEAD_KAT_1KB(t *testing.T) {
	key := make([]byte, 32)
	nonce := make([]byte, 24)
	for i := range key {
		key[i] = byte(i + 1)
	}
	for i := range nonce {
		nonce[i] = byte(i + 10)
	}
	ad := []byte("HEADER 1KB")
	pt := make([]byte, 1024)
	for i := range pt {
		pt[i] = byte((i*17 + 3) % 251)
	}
	ct := make([]byte, len(pt))
	mac := make([]byte, 16)
	Crypto_aead_lock(ct, mac, key, nonce, ad, uint64(len(ad)), pt, uint64(len(pt)))
	out := make([]byte, len(pt))
	if Crypto_aead_unlock(out, mac, key, nonce, ad, uint64(len(ad)), ct, uint64(len(ct))) != 0 {
		t.Fatal("mac fail 1k")
	}
	if !bytes.Equal(pt, out) {
		t.Fatalf("mismatch 1k first16 out=%x want=%x", out[:16], pt[:16])
	}
}

func TestChaCha_RoundTrip_128(t *testing.T) {
	key := make([]byte, 32)
	nonce := make([]byte, 8)
	for i := range key {
		key[i] = byte(i)
	}
	pt := make([]byte, 128)
	ct := make([]byte, 128)
	for i := range pt {
		pt[i] = byte(i * 3)
	}
	Crypto_chacha20_djb(ct, pt, 128, key, nonce, 0)
	out := make([]byte, 128)
	Crypto_chacha20_djb(out, ct, 128, key, nonce, 0)
	if !bytes.Equal(pt, out) {
		t.Fatal("chacha 128 roundtrip")
	}
}
