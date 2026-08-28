package c2blake2b

import "math/bits"

type Crypto_blake2b_ctx struct {
	Hash         [8]uint64
	Input_offset [2]uint64
	Input        [16]uint64
	Input_idx    uint64
	Hash_size    uint64
}

var (
	Iv    = [8]uint64{0x6a09e667f3bcc908, 0xbb67ae8584caa73b, 0x3c6ef372fe94f82b, 0xa54ff53a5f1d36f1, 0x510e527fade682d1, 0x9b05688c2b3e6c1f, 0x1f83d9abfb41bd6b, 0x5be0cd19137e2179}
	Sigma = [192]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 14, 10, 4, 8, 9, 15, 13, 6, 1, 12, 0, 2, 11, 7, 5, 3, 11, 8, 12, 0, 5, 2, 15, 13, 10, 14, 3, 6, 7, 1, 9, 4, 7, 9, 3, 1, 13, 12, 11, 14, 2, 6, 5, 10, 4, 0, 15, 8, 9, 0, 5, 7, 2, 4, 10, 15, 14, 1, 11, 12, 6, 8, 3, 13, 2, 12, 6, 10, 0, 11, 8, 3, 4, 13, 7, 5, 15, 14, 1, 9, 12, 5, 1, 15, 14, 13, 4, 10, 0, 7, 6, 3, 9, 2, 8, 11, 13, 11, 7, 14, 12, 1, 3, 9, 5, 0, 15, 4, 8, 6, 2, 10, 6, 15, 14, 9, 11, 3, 0, 8, 12, 2, 13, 7, 1, 4, 10, 5, 10, 2, 8, 4, 7, 6, 1, 5, 15, 11, 9, 14, 3, 12, 13, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 14, 10, 4, 8, 9, 15, 13, 6, 1, 12, 0, 2, 11, 7, 5, 3}
)

func rotr64(x uint64, n int) uint64 {
	return bits.RotateLeft64(x, -n)
}

func load32_le(s []byte) uint32 {
	return uint32(((uint32(s[0]) | uint32(uint32(s[1])<<8)) | uint32(uint32(s[2])<<16)) | uint32(uint32(s[3])<<24))
}

func load64_le(s []byte) uint64 {
	return uint64(uint64(load32_le(s[:])) | uint64(uint64(load32_le(s[4:]))<<32))
}

func store32_le(out []byte, in uint32) {
	out[0] = byte(in)
	out[1] = byte(in >> 8)
	out[2] = byte(in >> 16)
	out[3] = byte(in >> 24)
}

func store64_le(out []byte, in uint64) {
	store32_le(out[:], uint32(in))
	store32_le(out[4:], uint32(in>>32))
}

func load64_le_buf(dst []uint64, src []byte, size uint64) {
	for i := uint64(0); i < size; i++ {
		dst[i] = load64_le(src[int(i*8):])
	}
}

func store64_le_buf(dst []byte, src []uint64, size uint64) {
	for i := uint64(0); i < size; i++ {
		store64_le(dst[int(i*8):], src[i])
	}
}

func gap(x uint64, pow_2 uint64) uint64 {
	return (^x + 1) & (pow_2 - 1)
}
