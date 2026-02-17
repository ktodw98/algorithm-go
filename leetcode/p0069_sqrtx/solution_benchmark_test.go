package p0069

import "testing"

var testInt = 10000000000000

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(testInt)
	}
}
