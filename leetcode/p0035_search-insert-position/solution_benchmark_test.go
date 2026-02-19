package p0035

import (
	"math/rand"
	"testing"
)

var length int = 1000000
var BenchNums []int = func() []int {
	a := make([]int, length)

	for i := 0; i < length; i++ {
		a[i] = i + 1
	}

	return a
}()

var BenchTarget int = rand.Int() % length

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(BenchNums, BenchTarget)
	}
}

func BenchmarkSolve2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve2(BenchNums, BenchTarget)
	}
}

// BenchmarkSolve-10       51737053                22.76 ns/op            0 B/op          0 allocs/op
// BenchmarkSolve2-10      78420049                15.55 ns/op            0 B/op          0 allocs/op
