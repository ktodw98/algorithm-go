package p0009

import "testing"

var benchX = 12345678987654321

func BenchmarkSolve1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(benchX)
	}
}

func BenchmarkSolve2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve2(benchX)
	}
}

// BenchmarkSolve1-10      72000537                21.44 ns/op            0 B/op          0 allocs/op
// BenchmarkSolve2-10      25597519                44.43 ns/op           24 B/op          1 allocs/op
