package p0005

import (
	"strings"
	"testing"
)

var benchWorst = strings.Repeat("a", 5000)
var benchMid = strings.Repeat("x", 2000) + "racecar" + strings.Repeat("x", 2000)
var benchEasy = "abcdefghijklmnopqrstuvwxyz" + strings.Repeat("0123456789", 300)

func BenchmarkWorstSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(benchWorst)
	}
}

func BenchmarkMidSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(benchMid)
	}
}

func BenchmarkEasySolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(benchEasy)
	}
}

//BenchmarkWorstSolve-4             181473              6461 ns/op               0 B/op          0 allocs/op
//BenchmarkMidSolve-4               229765              5189 ns/op               0 B/op          0 allocs/op
//BenchmarkEasySolve-4                 100          10620222 ns/op               0 B/op          0 allocs/op
