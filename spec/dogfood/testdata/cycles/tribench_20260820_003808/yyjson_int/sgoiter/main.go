package main

import (
	"fmt"

	"trib/kernel"
)

func u64(v uint64) string { return fmt.Sprintf("%016x", v) }
func u32(v uint32) string { return fmt.Sprintf("%08x", v) }

func main() {
	// zero
	// single
	// double
	// medium
	// large
	// max
	{ buf := make([]byte, 32); n := kernel.Yyjson_write_u32(buf, 0); fmt.Printf("zero %s\n", string(buf[:n])) }
	{ buf := make([]byte, 32); n := kernel.Yyjson_write_u32(buf, 7); fmt.Printf("single %s\n", string(buf[:n])) }
	{ buf := make([]byte, 32); n := kernel.Yyjson_write_u32(buf, 42); fmt.Printf("double %s\n", string(buf[:n])) }
	{ buf := make([]byte, 32); n := kernel.Yyjson_write_u32(buf, 12345); fmt.Printf("medium %s\n", string(buf[:n])) }
	{ buf := make([]byte, 32); n := kernel.Yyjson_write_u32(buf, 123456789); fmt.Printf("large %s\n", string(buf[:n])) }
	{ buf := make([]byte, 32); n := kernel.Yyjson_write_u32(buf, 4294967295); fmt.Printf("max %s\n", string(buf[:n])) }
}
