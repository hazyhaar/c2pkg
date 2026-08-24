package c2archsimd

import (
	"encoding/hex"
	"testing"
)

var BenchSinkByte byte
var BenchSinkUint32 uint32

// 1. Hex Encode : sgoiter vs encoding/hex
func BenchmarkHexEncode_Sgoiter(b *testing.B) {
	in := [16]byte{0x01, 0x91, 0x7F, 0x8B, 0xC9, 0xA0, 0x71, 0x11, 0x80, 0x00, 0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC}
	var out [32]byte
	b.SetBytes(16)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C2archsimd_hex_encode16(in[:], out[:])
		BenchSinkByte ^= out[0]
	}
}

func BenchmarkHexEncode_Stdlib(b *testing.B) {
	in := [16]byte{0x01, 0x91, 0x7F, 0x8B, 0xC9, 0xA0, 0x71, 0x11, 0x80, 0x00, 0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC}
	var out [32]byte
	b.SetBytes(16)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hex.Encode(out[:], in[:])
		BenchSinkByte ^= out[0]
	}
}

// 2. Hex Decode : sgoiter vs encoding/hex
func BenchmarkHexDecode_Sgoiter(b *testing.B) {
	in := [32]byte{}
	copy(in[:], "01917f8bc9a071118000123456789abc")
	var out [16]byte
	b.SetBytes(32)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		st := C2archsimd_hex_decode32(in[:], out[:])
		BenchSinkUint32 ^= st
		BenchSinkByte ^= out[0]
	}
}

func BenchmarkHexDecode_Stdlib(b *testing.B) {
	in := [32]byte{}
	copy(in[:], "01917f8bc9a071118000123456789abc")
	var out [16]byte
	b.SetBytes(32)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = hex.Decode(out[:], in[:])
		BenchSinkByte ^= out[0]
	}
}

// 3. LUT16 : sgoiter vs boucle naïve
func BenchmarkLUT16_Sgoiter(b *testing.B) {
	in := [32]byte{}
	for i := range in {
		in[i] = byte(i)
	}
	var table16 C2archsimd_table16_t
	for i := range table16.B {
		table16.B[i] = byte(i * 17)
	}
	var out [32]byte
	b.SetBytes(32)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C2archsimd_lut16(in[:], &table16, out[:])
		BenchSinkByte ^= out[0]
	}
}

func BenchmarkLUT16_NaiveLoop(b *testing.B) {
	in := [32]byte{}
	for i := range in {
		in[i] = byte(i)
	}
	var table16 [16]byte
	for i := range table16 {
		table16[i] = byte(i * 17)
	}
	var out [32]byte
	b.SetBytes(32)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < 32; k++ {
			out[k] = table16[in[k]&0x0F]
		}
		BenchSinkByte ^= out[0]
	}
}

// 4. LUT256 Directe : sgoiter vs boucle naïve
func BenchmarkLUT256_Sgoiter(b *testing.B) {
	in := [32]byte{}
	for i := range in {
		in[i] = byte(i)
	}
	var table256 C2archsimd_table256_t
	for i := range table256.B {
		table256.B[i] = byte(i*31 + 7)
	}
	var out [32]byte
	b.SetBytes(32)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C2archsimd_lut256(in[:], &table256, out[:])
		BenchSinkByte ^= out[0]
	}
}

func BenchmarkLUT256_NaiveLoop(b *testing.B) {
	in := [32]byte{}
	for i := range in {
		in[i] = byte(i)
	}
	var table256 [256]byte
	for i := range table256 {
		table256[i] = byte(i*31 + 7)
	}
	var out [32]byte
	b.SetBytes(32)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < 32; k++ {
			out[k] = table256[in[k]]
		}
		BenchSinkByte ^= out[0]
	}
}

// 5. VintLens32 : sgoiter vs switch conditionnel naïf
func BenchmarkVintLens32_Sgoiter(b *testing.B) {
	in := [32]byte{0x25, 0x7B, 0x9D, 0xC2, 0x25, 0x7B, 0x9D, 0xC2, 0x25, 0x7B, 0x9D, 0xC2, 0x25, 0x7B, 0x9D, 0xC2}
	var out [32]byte
	b.SetBytes(32)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C2archsimd_vint_lens32(in[:], out[:])
		BenchSinkByte ^= out[0]
	}
}

func BenchmarkVintLens32_NaiveSwitch(b *testing.B) {
	in := [32]byte{0x25, 0x7B, 0x9D, 0xC2, 0x25, 0x7B, 0x9D, 0xC2, 0x25, 0x7B, 0x9D, 0xC2, 0x25, 0x7B, 0x9D, 0xC2}
	var out [32]byte
	b.SetBytes(32)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < 32; k++ {
			switch in[k] >> 6 {
			case 0:
				out[k] = 1
			case 1:
				out[k] = 2
			case 2:
				out[k] = 4
			default:
				out[k] = 8
			}
		}
		BenchSinkByte ^= out[0]
	}
}

// 6. AVX2 Vectorized Benchmarks (sgoiter AST emitted)
func BenchmarkVintLens32_AVX2_Sgoiter(b *testing.B) {
	in := [32]byte{0x25, 0x7B, 0x9D, 0xC2, 0x25, 0x7B, 0x9D, 0xC2, 0x25, 0x7B, 0x9D, 0xC2, 0x25, 0x7B, 0x9D, 0xC2}
	var out [32]byte
	b.SetBytes(32)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C2archsimd_vint_lens32_avx2(in[:], out[:])
		BenchSinkByte ^= out[0]
	}
}

func BenchmarkLUT16_AVX2_Sgoiter(b *testing.B) {
	in := [32]byte{}
	for i := range in {
		in[i] = byte(i)
	}
	var table16 C2archsimd_table16_t
	for i := range table16.B {
		table16.B[i] = byte(i * 17)
	}
	var out [32]byte
	b.SetBytes(32)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C2archsimd_lut16_avx2(in[:], &table16, out[:])
		BenchSinkByte ^= out[0]
	}
}

func BenchmarkLUT256_AVX2_Sgoiter(b *testing.B) {
	in := [32]byte{}
	for i := range in {
		in[i] = byte(i)
	}
	var table16Lo C2archsimd_table16_t
	var table16Hi C2archsimd_table16_t
	for i := range table16Lo.B {
		table16Lo.B[i] = byte(i * 17)
		table16Hi.B[i] = byte(i*31 + 7)
	}
	var out [32]byte
	b.SetBytes(32)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C2archsimd_lut256_avx2(in[:], &table16Lo, &table16Hi, out[:])
		BenchSinkByte ^= out[0]
	}
}

func BenchmarkHexEncode_AVX2_Sgoiter(b *testing.B) {
	in := [16]byte{0x01, 0x91, 0x7F, 0x8B, 0xC9, 0xA0, 0x71, 0x11, 0x80, 0x00, 0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC}
	var out [32]byte
	b.SetBytes(16)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C2archsimd_hex_encode16_avx2(in[:], out[:])
		BenchSinkByte ^= out[0]
	}
}

// 7. Photometric Solid Blend 8-voies (32 octets) : sgoiter vs AVX2
func BenchmarkBlendSolidPhotometric8_Sgoiter(b *testing.B) {
	in := [8]uint32{0xFF112233, 0xFF445566, 0xFF778899, 0xFFAABBCC, 0xFFDDEEFF, 0xFF001122, 0xFF334455, 0xFF667788}
	var lutLin [256]uint16
	var lutSrgb [4096]byte
	for i := range lutLin {
		lutLin[i] = uint16(i * 257)
	}
	for i := range lutSrgb {
		lutSrgb[i] = byte(i >> 4)
	}
	ctx := C2archsimd_solid_blend_ctx_t{
		Sa:              128,
		Inv_a:           127,
		Sr_scaled:       128 * 128,
		Sg_scaled:       128 * 128,
		Sb_scaled:       128 * 128,
		Lut_srgb_to_lin: lutLin[:],
		Lut_lin_to_srgb: lutSrgb[:],
	}
	var out [8]uint32
	b.SetBytes(32)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C2archsimd_blend_solid_photometric8(in[:], &ctx, out[:])
		BenchSinkByte ^= byte(out[0])
	}
}

func BenchmarkBlendSolidPhotometric8_AVX2(b *testing.B) {
	in := [8]uint32{0xFF112233, 0xFF445566, 0xFF778899, 0xFFAABBCC, 0xFFDDEEFF, 0xFF001122, 0xFF334455, 0xFF667788}
	var lutLin [256]uint16
	var lutSrgb [4096]byte
	for i := range lutLin {
		lutLin[i] = uint16(i * 257)
	}
	for i := range lutSrgb {
		lutSrgb[i] = byte(i >> 4)
	}
	ctx := C2archsimd_solid_blend_ctx_t{
		Sa:              128,
		Inv_a:           127,
		Sr_scaled:       128 * 128,
		Sg_scaled:       128 * 128,
		Sb_scaled:       128 * 128,
		Lut_srgb_to_lin: lutLin[:],
		Lut_lin_to_srgb: lutSrgb[:],
	}
	var out [8]uint32
	b.SetBytes(32)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C2archsimd_blend_solid_photometric8_avx2(in[:], &ctx, out[:])
		BenchSinkByte ^= byte(out[0])
	}
}

