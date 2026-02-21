package p0021

import "testing"

// list1: [-100, -96, -92, ..., 96] 50개
// list2: [-99, -95, -91, ..., 97]  50개
// 값이 완전히 교차하는 최대 길이 입력
var benchmarkList1 = SliceToListNode(func() []int {
	s := make([]int, 50)
	for i := range s {
		s[i] = -100 + i*4
	}
	return s
}())

var benchmarkList2 = SliceToListNode(func() []int {
	s := make([]int, 50)
	for i := range s {
		s[i] = -99 + i*4
	}
	return s
}())

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(benchmarkList1, benchmarkList2)
	}
}

// BenchmarkSolve-10        1211608              1008 ns/op            1600 B/op        100 allocs/op
