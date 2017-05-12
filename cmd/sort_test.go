package cmd

import (
	"testing"
)

func Benchmark_Qsort(b *testing.B) {
	b.StopTimer()
	l := genRandIntList(1000)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		qsort(l)
	}
}

func Benchmark_Qsort2(b *testing.B) {
	b.StopTimer()
	l := genRandIntList(1000)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		qsort2(l)
	}
}
