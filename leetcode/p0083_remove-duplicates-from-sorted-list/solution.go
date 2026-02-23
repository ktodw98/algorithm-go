package p0083

func solve(head *ListNode) *ListNode {
	current := head

	dummy := &ListNode{Val: -101}
	tail := dummy

	for current != nil {
		if current.Val != tail.Val {
			tail.Next = &ListNode{Val: current.Val}
			tail = tail.Next
		}
		current = current.Next
	}
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListNodeToSlice(ln *ListNode) []int {
	a := make([]int, 0)
	current := ln
	for current != nil {
		a = append(a, current.Val)
		current = current.Next
	}
	return a
}

func SliceToListNode(lst []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for _, v := range lst {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}

	return dummy.Next
}
