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
	// eq
	d0 := []byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	e0 := []byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	// neq
	d1 := []byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	e1 := []byte{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}
	{ x:=libc.Xmalloc(tls,16); y:=libc.Xmalloc(tls,16); put(tls,x,d0); put(tls,y,e0)
		r:=crypto_verify_16(tls,x,y); fmt.Printf("eq %d\n", int(r)); libc.Xfree(tls,x); libc.Xfree(tls,y) }
	{ x:=libc.Xmalloc(tls,16); y:=libc.Xmalloc(tls,16); put(tls,x,d1); put(tls,y,e1)
		r:=crypto_verify_16(tls,x,y); fmt.Printf("neq %d\n", int(r)); libc.Xfree(tls,x); libc.Xfree(tls,y) }
}

func typesize(n int) uint64 {
	if n <= 0 { return 1 }
	return uint64(n)
}
