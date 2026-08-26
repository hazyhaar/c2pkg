// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2fused

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func BenchmarkSeal(b *testing.B) {
	for _, n := range []int{1350, 4096, 8192, 65536} {
		key := make([]byte, 32)
		nonce := make([]byte, 12)
		in := make([]byte, n)
		outF := make([]byte, n)
		outS := make([]byte, n)
		tagF := make([]byte, 16)
		tagS := make([]byte, 16)
		rand.Read(key)
		rand.Read(nonce)
		rand.Read(in)
		tag := fmt.Sprintf("%d", n)
		b.Run("fused/"+tag, func(b *testing.B) {
			b.SetBytes(int64(n))
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				aeadFused(outF, tagF, in, nil, key, nonce)
			}
		})
		b.Run("seq/"+tag, func(b *testing.B) {
			b.SetBytes(int64(n))
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				aeadSeq(outS, tagS, in, nil, key, nonce)
			}
		})
	}
}
