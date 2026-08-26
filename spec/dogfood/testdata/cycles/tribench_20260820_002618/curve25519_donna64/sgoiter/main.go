package main

import (
	"fmt"

	"trib/kernel"
)

func u64(v uint64) string { return fmt.Sprintf("%016x", v) }
func u32(v uint32) string { return fmt.Sprintf("%08x", v) }

func main() {
	// zero
	// pattern
	{ in := []uint64{0x123456789, 0x23456789a, 0x3456789ab, 0x456789abc, 0x56789abcd};
		out := make([]uint64, 5);
		kernel.Curve25519_f51_mul121666(out, in);
		var s string; for _, v := range out { s += u64(v) }; fmt.Printf("zero %s\n", s) }
	{ in := []uint64{0x123456789, 0x23456789a, 0x3456789ab, 0x456789abc, 0x56789abcd};
		out := make([]uint64, 5);
		kernel.Curve25519_f51_mul121666(out, in);
		var s string; for _, v := range out { s += u64(v) }; fmt.Printf("pattern %s\n", s) }
}
