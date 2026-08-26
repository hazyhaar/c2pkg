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
	{ st:=libc.Xmalloc(tls,16); block:=libc.Xmalloc(tls,64)
		iv:=[]uint32{0x67452301,0xefcdab89,0x98badcfe,0x10325476}
		for j:=0;j<4;j++ { *(*uint32)(unsafe.Pointer(st+uintptr(j*4)))=iv[j] }
		for j:=0;j<64;j++ { *(*byte)(unsafe.Pointer(block+uintptr(j)))=0 }
		put(tls, block, d0)
		md5_transform_full_block(tls, st, block)
		var s string; for j:=0;j<4;j++ { s += fmtU32(*(*uint32)(unsafe.Pointer(st+uintptr(j*4)))) }
		fmt.Printf("zero %s\n", s); libc.Xfree(tls,st); libc.Xfree(tls,block) }
	{ st:=libc.Xmalloc(tls,16); block:=libc.Xmalloc(tls,64)
		iv:=[]uint32{0x67452301,0xefcdab89,0x98badcfe,0x10325476}
		for j:=0;j<4;j++ { *(*uint32)(unsafe.Pointer(st+uintptr(j*4)))=iv[j] }
		for j:=0;j<64;j++ { *(*byte)(unsafe.Pointer(block+uintptr(j)))=0 }
		put(tls, block, d1)
		md5_transform_full_block(tls, st, block)
		var s string; for j:=0;j<4;j++ { s += fmtU32(*(*uint32)(unsafe.Pointer(st+uintptr(j*4)))) }
		fmt.Printf("pattern %s\n", s); libc.Xfree(tls,st); libc.Xfree(tls,block) }
}

func typesize(n int) uint64 {
	if n <= 0 { return 1 }
	return uint64(n)
}
