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
	{ a:=libc.Xmalloc(tls,4); b:=libc.Xmalloc(tls,4); c:=libc.Xmalloc(tls,4); d:=libc.Xmalloc(tls,4)
		*(*uint32)(unsafe.Pointer(a))=0x11111111; *(*uint32)(unsafe.Pointer(b))=0x22222222
		*(*uint32)(unsafe.Pointer(c))=0x33333333; *(*uint32)(unsafe.Pointer(d))=0x44444444
		chacha20_quarter_round(tls,a,b,c,d)
		fmt.Printf("zero %s%s%s%s\n", fmtU32(*(*uint32)(unsafe.Pointer(a))), fmtU32(*(*uint32)(unsafe.Pointer(b))),
			fmtU32(*(*uint32)(unsafe.Pointer(c))), fmtU32(*(*uint32)(unsafe.Pointer(d))))
		libc.Xfree(tls,a);libc.Xfree(tls,b);libc.Xfree(tls,c);libc.Xfree(tls,d) }
	{ a:=libc.Xmalloc(tls,4); b:=libc.Xmalloc(tls,4); c:=libc.Xmalloc(tls,4); d:=libc.Xmalloc(tls,4)
		*(*uint32)(unsafe.Pointer(a))=0x11111111; *(*uint32)(unsafe.Pointer(b))=0x22222222
		*(*uint32)(unsafe.Pointer(c))=0x33333333; *(*uint32)(unsafe.Pointer(d))=0x44444444
		chacha20_quarter_round(tls,a,b,c,d)
		fmt.Printf("pattern %s%s%s%s\n", fmtU32(*(*uint32)(unsafe.Pointer(a))), fmtU32(*(*uint32)(unsafe.Pointer(b))),
			fmtU32(*(*uint32)(unsafe.Pointer(c))), fmtU32(*(*uint32)(unsafe.Pointer(d))))
		libc.Xfree(tls,a);libc.Xfree(tls,b);libc.Xfree(tls,c);libc.Xfree(tls,d) }
}

func typesize(n int) uint64 {
	if n <= 0 { return 1 }
	return uint64(n)
}
