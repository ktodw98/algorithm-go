package p0067

import "testing"

var benchmarkA = func() string {
	a := ""
	for i := 0; i < 10000; i++ {
		a = a + "1"
	}
	return a
}()
var benchmarkB = func() string {
	return "1"
}()

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(benchmarkA, benchmarkB)
	}
}
