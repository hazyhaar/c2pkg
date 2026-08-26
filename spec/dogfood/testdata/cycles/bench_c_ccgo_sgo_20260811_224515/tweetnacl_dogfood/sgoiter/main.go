package main

import (
	"fmt"

	"trib/kernel"
)

func u64(v uint64) string { return fmt.Sprintf("%016x", v) }
func u32(v uint32) string { return fmt.Sprintf("%08x", v) }

func main() {
	// eq
	d0 := []byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	e0 := []byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	// neq
	d1 := []byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	e1 := []byte{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}
	fmt.Printf("eq %d\n", kernel.Crypto_verify_16(d0, e0))
	fmt.Printf("neq %d\n", kernel.Crypto_verify_16(d1, e1))
}
