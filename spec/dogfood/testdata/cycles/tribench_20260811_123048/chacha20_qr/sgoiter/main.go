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
	{ a, b, c, d := uint32(0x11111111), uint32(0x22222222), uint32(0x33333333), uint32(0x44444444);
		kernel.Chacha20_quarter_round(&a, &b, &c, &d); fmt.Printf("zero %s%s%s%s\n", u32(a), u32(b), u32(c), u32(d)) }
	{ a, b, c, d := uint32(0x11111111), uint32(0x22222222), uint32(0x33333333), uint32(0x44444444);
		kernel.Chacha20_quarter_round(&a, &b, &c, &d); fmt.Printf("pattern %s%s%s%s\n", u32(a), u32(b), u32(c), u32(d)) }
}
