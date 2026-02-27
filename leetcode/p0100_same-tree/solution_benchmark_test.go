package p0100

import "testing"

var BenchMarkP = func() *TreeNode {
	s := make([]int, 1_000_000)
	for i := range s {
		s[i] = i
	}
	return SliceToTree(s)
}()
var BenchMarkQ = func() *TreeNode {
	s := make([]int, 1_000_000)
	for i := range s {
		s[i] = i
	}
	return SliceToTree(s)
}()

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(BenchMarkP, BenchMarkQ)
	}
}

// BenchmarkSolve-10            255           4632084 ns/op               0 B/op          0 allocs/op
