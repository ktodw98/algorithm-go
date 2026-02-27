package p0169

import "testing"

var n = 50000
var BenchmarkNums = func() []int {
	a := make([]int, n)
	half := n / 2

	for i := 0; i < half-1; i++ {
		a[i] = 1
	}

	for i := half - 1; i < n; i++ {
		a[i] = 2
	}

	return a
}()

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(BenchmarkNums)
	}
}

// BenchmarkSolve-10           1736            650041 ns/op               0 B/op          0 allocs/op
