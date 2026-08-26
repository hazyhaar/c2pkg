package main

import (
	"fmt"

	"trib/kernel"
)

func u64(v uint64) string { return fmt.Sprintf("%016x", v) }
func u32(v uint32) string { return fmt.Sprintf("%08x", v) }

func main() {
	// zero
	d0 := []byte{}
	// pattern
	d1 := []byte{165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165}
	{ st := []uint32{0x67452301, 0xefcdab89, 0x98badcfe, 0x10325476}; block := make([]byte, 64); copy(block, d0);
		kernel.Md5_transform_block(st, block); var s string; for _, v := range st { s += u32(v) }; fmt.Printf("zero %s\n", s) }
	{ st := []uint32{0x67452301, 0xefcdab89, 0x98badcfe, 0x10325476}; block := make([]byte, 64); copy(block, d1);
		kernel.Md5_transform_block(st, block); var s string; for _, v := range st { s += u32(v) }; fmt.Printf("pattern %s\n", s) }
}
