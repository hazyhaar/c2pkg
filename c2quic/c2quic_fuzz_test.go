// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2quic

import "testing"

func FuzzVintEncode(f *testing.F) {
	f.Add(uint64(37), uint32(0), []byte{0, 0, 0, 0, 0, 0, 0, 0})
	f.Add(uint64(1)<<62, uint32(0), []byte{0, 0, 0, 0})
	f.Add(uint64(15293), uint32(100), []byte{1, 2, 3})
	f.Fuzz(func(t *testing.T, v uint64, off uint32, buf []byte) {
		var r C2quic_vint_t
		C2quic_vint_encode(v, buf, uint32(len(buf)), off, &r)
	})
}

func FuzzHPRemove(f *testing.F) {
	f.Add([]byte{0xC3, 1, 2, 3, 4}, uint32(1), byte(0x0F))
	f.Add([]byte{0x40}, uint32(99), byte(0x1F))
	f.Fuzz(func(t *testing.T, pkt []byte, pnOff uint32, mask0 byte) {
		var io C2quic_hp_io_t
		io.Mask[0] = mask0
		C2quic_hp_remove(pkt, uint32(len(pkt)), pnOff, &io)
	})
}

func FuzzUnpackFrames(f *testing.F) {
	f.Add([]byte{0x0A, 0x07, 0x04, 'A', 'B', 'C', 'D'})
	f.Add([]byte{0x40, 0x01})
	f.Add([]byte{0x00, 0x01})
	f.Fuzz(func(t *testing.T, body []byte) {
		var slab C2quic_frame_slab_t
		C2quic_unpack_frames(body, uint32(len(body)), &slab)
	})
}
