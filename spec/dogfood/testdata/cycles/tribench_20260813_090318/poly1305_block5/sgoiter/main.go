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
	{ h := make([]uint32, 5); r := []uint32{1,2,3,4,5}; m := []uint32{9,8,7,6};
		kernel.Poly1305_block5(h, r, m); var s string; for _, v := range h { s += u32(v) }; fmt.Printf("zero %s\n", s) }
	{ h := make([]uint32, 5); r := []uint32{1,2,3,4,5}; m := []uint32{9,8,7,6};
		kernel.Poly1305_block5(h, r, m); var s string; for _, v := range h { s += u32(v) }; fmt.Printf("pattern %s\n", s) }
}
