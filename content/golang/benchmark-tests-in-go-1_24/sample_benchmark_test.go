package sample_test

import (
	"math/rand/v2"
	"sample"
	"testing"
)

func intList(n int) []int {
	list := make([]int, n)
	for i := range list {
		list[i] = rand.Int()
	}
	return list
}

func BenchmarkSum(b *testing.B) {
	b.ReportAllocs()
	input := intList(128 << 10)

	for b.Loop() {
		sample.Sum(input)
	}
}

func BenchmarkAdd(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		_ = sample.Add(1, 2)
	}
}
