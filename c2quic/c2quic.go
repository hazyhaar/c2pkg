package c2quic

const (
	OK       = uint32(0)
	Trunc    = uint32(1)
	Overflow = uint32(2)
	Invalid  = uint32(3)
	Full     = uint32(4)
)

func conventionalVintDecode(buf []byte, off int) (uint64, int, bool) {
	if off < 0 || off >= len(buf) {
		return 0, 0, false
	}
	n := 1 << (buf[off] >> 6)
	if off+n > len(buf) {
		return 0, 0, false
	}
	v := uint64(buf[off] & 0x3f)
	for i := 1; i < n; i++ {
		v = (v << 8) | uint64(buf[off+i])
	}
	return v, n, true
}
