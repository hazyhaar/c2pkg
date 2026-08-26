// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2quic

import "testing"

func BenchmarkVintDecode_sgoiter(b *testing.B) {
	var r C2quic_vint_t
	enc := make([]byte, 8)
	C2quic_vint_encode(494878333, enc, 8, 0, &r)
	var acc uint64
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C2quic_vint_decode(enc, r.Nbytes, 0, &r)
		acc ^= r.Value
	}
	sink = acc
}

var sink uint64

func BenchmarkVintDecode_conventional(b *testing.B) {
	var r C2quic_vint_t
	enc := make([]byte, 8)
	C2quic_vint_encode(494878333, enc, 8, 0, &r)
	var acc uint64
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v, _, _ := conventionalVintDecode(enc, 0)
		acc ^= v
	}
	sink = acc
}

func BenchmarkVintDecodeBatch8(b *testing.B) {
	buf := make([]byte, 64)
	var r C2quic_vint_t
	var batch C2quic_batch8_t
	for k := 0; k < 8; k++ {
		batch.Offs[k] = uint32(k * 8)
		C2quic_vint_encode(uint64(1000+k), buf, 64, batch.Offs[k], &r)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C2quic_vint_decode_batch8_scalar(buf, 64, &batch)
	}
}

func BenchmarkHeaderProtectApplyRemove(b *testing.B) {
	pkt := make([]byte, 64)
	pkt[0] = 0xC3
	pkt[1] = 1
	for i := 5; i < 37; i++ {
		pkt[i] = 0xA5
	}
	var io C2quic_hp_io_t
	io.Mask[0] = 0x0F
	io.Mask[1] = 0x11
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C2quic_hp_apply(pkt, 64, 1, 1, &io)
		C2quic_hp_remove(pkt, 64, 1, &io)
	}
}

func BenchmarkUnpackStream(b *testing.B) {
	body := []byte{0x0A, 0x07, 0x04, 'A', 'B', 'C', 'D'}
	var slab C2quic_frame_slab_t
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C2quic_unpack_frames(body, uint32(len(body)), &slab)
	}
}
