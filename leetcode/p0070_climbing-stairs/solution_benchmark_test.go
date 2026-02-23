package p0070

import "testing"

var BenchNum = 10000000

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(BenchNum)
	}
}

// BenchmarkSolve-10             78          15224527 ns/op        80003079 B/op          1 allocs/op
