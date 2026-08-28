package main

import (
	"math/rand"
	"testing"
)

var sinkArray [2]uint64
var sinkStruct U128Struct
var sinkLo, sinkHi uint64

func Benchmark_Add128_Array_StackSpill(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	a := [2]uint64{r.Uint64(), r.Uint64()}
	val := [2]uint64{r.Uint64(), r.Uint64()}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a = Add128_Array(a, val)
	}
	sinkArray = a
}

func Benchmark_Add128_Struct_RegABI(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	a := U128Struct{Lo: r.Uint64(), Hi: r.Uint64()}
	val := U128Struct{Lo: r.Uint64(), Hi: r.Uint64()}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a = Add128_Struct(a, val)
	}
	sinkStruct = a
}

func Benchmark_Add128_Scalars_RegABI(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	aLo, aHi := r.Uint64(), r.Uint64()
	bLo, bHi := r.Uint64(), r.Uint64()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		aLo, aHi = Add128_Scalars(aLo, aHi, bLo, bHi)
	}
	sinkLo, sinkHi = aLo, aHi
}

func Benchmark_Xor128_Array_StackSpill(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	a := [2]uint64{r.Uint64(), r.Uint64()}
	val := [2]uint64{r.Uint64(), r.Uint64()}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a = Xor128_Array(a, val)
	}
	sinkArray = a
}

func Benchmark_Xor128_Struct_RegABI(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	a := U128Struct{Lo: r.Uint64(), Hi: r.Uint64()}
	val := U128Struct{Lo: r.Uint64(), Hi: r.Uint64()}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a = Xor128_Struct(a, val)
	}
	sinkStruct = a
}
