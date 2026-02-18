package p0026

import (
	"math/rand"
	"testing"
)

func MakeBenchNums() []int {
	n := 10000
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a = append(a, rand.Int()%200-100)
	}
	return a
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchNums := MakeBenchNums()
		solve(benchNums)
	}
}
