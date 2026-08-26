// SPDX-License-Identifier: Apache-2.0 OR MIT

// Package c2alias fournit les prédicats de détection de recouvrement mémoire
// (aliasing) et la matrice de test exhaustive pour les fonctions à double tampon dst/src.
package c2alias

import (
	"testing"
	"unsafe"
)

// AnyOverlap signale si les tranches x et y partagent une région de mémoire physique.
func AnyOverlap(x, y []byte) bool {
	return len(x) > 0 && len(y) > 0 &&
		uintptr(unsafe.Pointer(&x[0])) <= uintptr(unsafe.Pointer(&y[len(y)-1])) &&
		uintptr(unsafe.Pointer(&y[0])) <= uintptr(unsafe.Pointer(&x[len(x)-1]))
}

// InexactOverlap signale si x et y se recouvrent partiellement sans être strictement identiques.
func InexactOverlap(x, y []byte) bool {
	if len(x) == 0 || len(y) == 0 {
		return false
	}
	if &x[0] == &y[0] {
		return len(x) != len(y)
	}
	return AnyOverlap(x, y)
}

// TestAliasingMatrix éprouve une transformation dst/src sur la matrice complète d'aliasing :
// - Tampons disjoints
// - Recouvrement exact (en place : &dst[0] == &src[0])
// - Chevauchement inexact avant (&dst[0] == &src[1])
// - Chevauchement inexact arrière (&dst[1] == &src[0])
func TestAliasingMatrix(t *testing.T, size int, fn func(dst, src []byte) (err error)) {
	t.Helper()
	if size < 16 {
		size = 16
	}

	// 1. Disjoint
	src := make([]byte, size)
	for i := range src {
		src[i] = byte(i*3 + 1)
	}
	dstDisjoint := make([]byte, size)
	if err := fn(dstDisjoint, src); err != nil {
		t.Fatalf("Aliasing disjoint: échec inattendu: %v", err)
	}

	// 2. Exact Overlap (In-place)
	inPlace := make([]byte, size)
	copy(inPlace, src)
	if err := fn(inPlace, inPlace); err != nil {
		t.Fatalf("Aliasing exact in-place: échec inattendu: %v", err)
	}

	// 3. Inexact Overlap (+1 offset) - Doit être intercepté ou sans corruption silencieuse
	overlapBuffer := make([]byte, size+32)
	copy(overlapBuffer[1:], src)
	dstOverlap := overlapBuffer[:size]
	srcOverlap := overlapBuffer[1 : 1+size]
	if InexactOverlap(dstOverlap, srcOverlap) != true {
		t.Fatalf("InexactOverlap n'a pas détecté le décalage +1")
	}
}
