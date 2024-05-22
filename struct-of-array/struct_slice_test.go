package main

import (
	"testing"
)

func generateBarData(size int) StructOfSlice {
	a := make([]int64, size)
	b := make([]int64, size)
	for i := 0; i < size; i++ {
		a[i] = int64(i)
		b[i] = int64(i)
	}
	return StructOfSlice{a: a, b: b}
}

func generateFooData(size int) []Struct {
	foos := make([]Struct, size)
	for i := 0; i < size; i++ {
		foos[i] = Struct{a: int64(i), b: int64(i)}
	}
	return foos
}

func BenchmarkStructOfSlices(b *testing.B) {
	bar := generateBarData(16)
	for i := 0; i < b.N; i++ {
		sumStructOfSlice(bar)
	}
}

func BenchmarkSliceOfStruct(b *testing.B) {
	foos := generateFooData(16)
	for i := 0; i < b.N; i++ {
		sumSliceOfStruct(foos)
	}
}
