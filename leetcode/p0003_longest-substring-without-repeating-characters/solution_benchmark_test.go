package p0003

import (
	"strings"
	"testing"
)

// 전부 같은 문자 10K — 매번 중복, inner loop 최대 빈도
var benchAllSame = strings.Repeat("a", 10000)

// 26자 반복 10K — set이 계속 차고 비워짐
var benchRepeatingAlpha = strings.Repeat("abcdefghijklmnopqrstuvwxyz", 385)

// 전부 유니크 26자 — 중복 없이 쭉 진행
var benchAllUnique = "abcdefghijklmnopqrstuvwxyz"

func BenchmarkAllSame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(benchAllSame)
	}
}

func BenchmarkRepeatingAlpha(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(benchRepeatingAlpha)
	}
}

func BenchmarkAllUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(benchAllUnique)
	}
}

// BenchmarkAllSame-10                12675             99393 ns/op               0 B/op          0 allocs/op
// BenchmarkRepeatingAlpha-10          6085            182051 ns/op             936 B/op          5 allocs/op
// BenchmarkAllUnique-10            1861281               611.2 ns/op           936 B/op          5 allocs/op
