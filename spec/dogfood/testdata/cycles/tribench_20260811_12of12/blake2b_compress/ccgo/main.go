package main

import (
	"fmt"
	"unsafe"

	"modernc.org/libc"
)

func fmtU64(v uint64) string { return fmt.Sprintf("%016x", v) }
func fmtU32(v uint32) string { return fmt.Sprintf("%08x", v) }

func put(tls *libc.TLS, p uintptr, b []byte) {
	for i := 0; i < len(b); i++ {
		*(*byte)(unsafe.Pointer(p + uintptr(i))) = b[i]
	}
}
func get(tls *libc.TLS, p uintptr, n int) []byte {
	out := make([]byte, n)
	for i := 0; i < n; i++ {
		out[i] = *(*byte)(unsafe.Pointer(p + uintptr(i)))
	}
	return out
}

func main() {
	tls := libc.NewTLS()
	defer tls.Close()
	// zero
	d0 := []byte{}
	// pattern
	d1 := []byte{165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165,165}
	{ h := libc.Xmalloc(tls, 8*8); block := libc.Xmalloc(tls, 128)
		for j:=0;j<8;j++ { *(*uint64)(unsafe.Pointer(h+uintptr(j*8))) = 0 }
		for j:=0;j<128;j++ { *(*byte)(unsafe.Pointer(block+uintptr(j))) = 0 }
		put(tls, block, d0)
		blake2b_compress_block(tls, h, block, 0, 0, 0xffffffffffffffff, 0xffffffffffffffff)
		var s string; for j:=0;j<8;j++ { s += fmtU64(*(*uint64)(unsafe.Pointer(h+uintptr(j*8)))) }
		fmt.Printf("zero %s\n", s); libc.Xfree(tls,h); libc.Xfree(tls,block) }
	{ h := libc.Xmalloc(tls, 8*8); block := libc.Xmalloc(tls, 128)
		for j:=0;j<8;j++ { *(*uint64)(unsafe.Pointer(h+uintptr(j*8))) = 0 }
		for j:=0;j<128;j++ { *(*byte)(unsafe.Pointer(block+uintptr(j))) = 0 }
		put(tls, block, d1)
		blake2b_compress_block(tls, h, block, 0, 0, 0xffffffffffffffff, 0xffffffffffffffff)
		var s string; for j:=0;j<8;j++ { s += fmtU64(*(*uint64)(unsafe.Pointer(h+uintptr(j*8)))) }
		fmt.Printf("pattern %s\n", s); libc.Xfree(tls,h); libc.Xfree(tls,block) }
}

func typesize(n int) uint64 {
	if n <= 0 { return 1 }
	return uint64(n)
}
