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
	{ h := make([]uint32, 5); r := []uint32{0x3ffffff,0x3fffffe,0x3fffffd,0x3fffffc,0x3fffffb};
		block := make([]byte, 16); for j := range block { block[j] = 0x5a };
		kernel.Poly1305_donna32_block(h, r, block, 1 << 24);
		var s string; for _, v := range h { s += u32(v) }; fmt.Printf("zero %s\n", s) }
	{ h := make([]uint32, 5); r := []uint32{0x3ffffff,0x3fffffe,0x3fffffd,0x3fffffc,0x3fffffb};
		block := make([]byte, 16); for j := range block { block[j] = 0x5a };
		kernel.Poly1305_donna32_block(h, r, block, 1 << 24);
		var s string; for _, v := range h { s += u32(v) }; fmt.Printf("pattern %s\n", s) }
}
