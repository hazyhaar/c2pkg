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
	_ = d0
	// pattern
	d1 := []byte{165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165}
	_ = d1
	{ h := make([]uint64, 8); block := make([]byte, 128); copy(block, d0);
		kernel.Blake2b_compress_block(h, block, 0, 0, 0xffffffffffffffff, 0xffffffffffffffff);
		var s string; for _, v := range h { s += u64(v) }; fmt.Printf("zero %s\n", s) }
	{ h := make([]uint64, 8); block := make([]byte, 128); copy(block, d1);
		kernel.Blake2b_compress_block(h, block, 0, 0, 0xffffffffffffffff, 0xffffffffffffffff);
		var s string; for _, v := range h { s += u64(v) }; fmt.Printf("pattern %s\n", s) }
}
