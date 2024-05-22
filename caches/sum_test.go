package main

import (
	"testing"
)

func BenchmarkCachesSum2(b *testing.B) {
	s := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	for i := 0; i < b.N; i++ {
		sum2(s)
	}
}

func BenchmarkCachesSum8(b *testing.B) {
	s := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	for i := 0; i < b.N; i++ {
		sum8(s)
	}
}
