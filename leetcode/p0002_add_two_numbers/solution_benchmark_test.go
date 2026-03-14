package p0002

import "testing"

func makeBenchmarkList(size int, digit int) *ListNode {
	dummy := &ListNode{}
	head := dummy
	for i := 0; i < size; i++ {
		head.Next = &ListNode{Val: digit}
		head = head.Next
	}
	return dummy.Next
}

var BenchMarkList1 = makeBenchmarkList(10000, 9)
var BenchMarkList2 = makeBenchmarkList(10000, 9)

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(BenchMarkList1, BenchMarkList2)
	}
}

// BenchmarkSolve-10           8811            121951 ns/op          160017 B/op      10001 allocs/op
