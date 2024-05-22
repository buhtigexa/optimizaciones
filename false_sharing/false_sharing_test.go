package main

import "testing"

const n = 1_000_000

var globalResult1 Result

func BenchmarkFalseSharingBad(b *testing.B) {
	var local Result
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		inputs := make([]Input, n)
		b.StartTimer()
		local = falseSharingBadSolution(inputs)
	}
	globalResult1 = local
}

var globalResult2 ResultPaddingSolution

func BenchmarkFalseSharingFixedSolution_Padding(b *testing.B) {
	var local ResultPaddingSolution
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		inputs := make([]Input, n)
		b.StartTimer()
		local = falseSharingPaddingSolution(inputs)
	}
	globalResult2 = local
}
