package p0066

import "testing"

var BenchCarryDigits = func() []int {
	benchCarryDigit := make([]int, 1000000)
	for i := 0; i < len(benchCarryDigit); i++ {
		benchCarryDigit[i] = 9
	}
	return benchCarryDigit
}()

var BenchNonCarryDigits = func() []int {
	benchNonCarryDigit := make([]int, 1000000)
	for i := 0; i < len(benchNonCarryDigit); i++ {
		benchNonCarryDigit[i] = 8
	}
	return benchNonCarryDigit
}()

func BenchmarkCarrySolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(BenchCarryDigits)
	}
}
func BenchmarkNonCarrySolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(BenchNonCarryDigits)
	}
}

// BenchmarkCarrySolve-10               896           1274266 ns/op        16007175 B/op          2 allocs/op
// BenchmarkNonCarrySolve-10           1190           1055515 ns/op        16007170 B/op          2 allocs/op
