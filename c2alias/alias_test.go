// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2alias

import (
	"testing"
)

func TestOverlapDetection(t *testing.T) {
	buf := make([]byte, 128)

	s1 := buf[0:32]
	s2 := buf[32:64]
	if AnyOverlap(s1, s2) {
		t.Fatal("Disjoint buffers must not overlap")
	}
	if InexactOverlap(s1, s2) {
		t.Fatal("Disjoint buffers must not inexact-overlap")
	}

	s3 := buf[0:32]
	if InexactOverlap(s1, s3) {
		t.Fatal("Exact same slices must not inexact-overlap")
	}

	s4 := buf[1:33]
	if !AnyOverlap(s1, s4) {
		t.Fatal("Overlapping slices must report AnyOverlap")
	}
	if !InexactOverlap(s1, s4) {
		t.Fatal("Shifted slices must report InexactOverlap")
	}
}

func TestMatrixRunner(t *testing.T) {
	TestAliasingMatrix(t, 64, func(dst, src []byte) error {
		copy(dst, src)
		return nil
	})
}
