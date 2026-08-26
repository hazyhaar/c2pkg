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
	{ a := []uint32{0x11111111}; b := []uint32{0x22222222}; c := []uint32{0x33333333}; d := []uint32{0x44444444};
		kernel.Chacha20_quarter_round(a, b, c, d); fmt.Printf("zero %s%s%s%s\n", u32(a[0]), u32(b[0]), u32(c[0]), u32(d[0])) }
	{ a := []uint32{0x11111111}; b := []uint32{0x22222222}; c := []uint32{0x33333333}; d := []uint32{0x44444444};
		kernel.Chacha20_quarter_round(a, b, c, d); fmt.Printf("pattern %s%s%s%s\n", u32(a[0]), u32(b[0]), u32(c[0]), u32(d[0])) }
}
