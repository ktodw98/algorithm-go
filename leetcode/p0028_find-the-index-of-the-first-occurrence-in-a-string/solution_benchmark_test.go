package p0028

import (
	"strings"
	"testing"
)

// worst case: 매 위치에서 needle 끝자리까지 비교 후 실패
var benchHaystack = strings.Repeat("a", 99999) + "b" // 100000자, 끝만 b
var benchNeedle = strings.Repeat("a", 999) + "b"     // 1000자, 끝만 b

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(benchHaystack, benchNeedle)
	}
}
func BenchmarkSolve2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve2(benchHaystack, benchNeedle)
	}
}

func BenchmarkSolve3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve3(benchHaystack, benchNeedle)
	}
}

func BenchmarkSolve4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve4(benchHaystack, benchNeedle)
	}
}

// BenchmarkSolve-10            294           4010977 ns/op         3168046 B/op     198002 allocs/op
// BenchmarkSolve2-10           897           1351273 ns/op               0 B/op          0 allocs/op
// BenchmarkSolve3-10            33          35039104 ns/op               0 B/op          0 allocs/op
// BenchmarkSolve4-10            33          34994452 ns/op               0 B/op          0 allocs/op
