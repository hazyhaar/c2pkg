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
	// pattern
	{ h:=libc.Xmalloc(tls,20); r:=libc.Xmalloc(tls,20); m:=libc.Xmalloc(tls,16)
		rv:=[]uint32{1,2,3,4,5}; mv:=[]uint32{9,8,7,6}
		for j:=0;j<5;j++ { *(*uint32)(unsafe.Pointer(h+uintptr(j*4)))=0; *(*uint32)(unsafe.Pointer(r+uintptr(j*4)))=rv[j] }
		for j:=0;j<4;j++ { *(*uint32)(unsafe.Pointer(m+uintptr(j*4)))=mv[j] }
		poly1305_block5(tls,h,r,m)
		var s string; for j:=0;j<5;j++ { s += fmtU32(*(*uint32)(unsafe.Pointer(h+uintptr(j*4)))) }
		fmt.Printf("zero %s\n", s); libc.Xfree(tls,h);libc.Xfree(tls,r);libc.Xfree(tls,m) }
	{ h:=libc.Xmalloc(tls,20); r:=libc.Xmalloc(tls,20); m:=libc.Xmalloc(tls,16)
		rv:=[]uint32{1,2,3,4,5}; mv:=[]uint32{9,8,7,6}
		for j:=0;j<5;j++ { *(*uint32)(unsafe.Pointer(h+uintptr(j*4)))=0; *(*uint32)(unsafe.Pointer(r+uintptr(j*4)))=rv[j] }
		for j:=0;j<4;j++ { *(*uint32)(unsafe.Pointer(m+uintptr(j*4)))=mv[j] }
		poly1305_block5(tls,h,r,m)
		var s string; for j:=0;j<5;j++ { s += fmtU32(*(*uint32)(unsafe.Pointer(h+uintptr(j*4)))) }
		fmt.Printf("pattern %s\n", s); libc.Xfree(tls,h);libc.Xfree(tls,r);libc.Xfree(tls,m) }
}

func typesize(n int) uint64 {
	if n <= 0 { return 1 }
	return uint64(n)
}
