package p0083

import "testing"

var BenchLength = 10000
var BenchMarkLstNode *ListNode = func() *ListNode {
	head := &ListNode{}
	tail := head

	for i := 0; i < BenchLength; i++ {
		Node1 := &ListNode{Val: i}
		tail.Next = Node1
		tail = tail.Next

		Node2 := &ListNode{Val: i}
		tail.Next = Node2
		tail = tail.Next
	}

	return head
}()

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(BenchMarkLstNode)
	}
}

// BenchmarkSolve-10          11728            103239 ns/op          159984 B/op       9999 allocs/op
