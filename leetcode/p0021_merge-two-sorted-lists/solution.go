package p0021

func solve(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	tail1, tail2 := list1, list2
	for {
		if tail1 == nil && tail2 == nil {
			break
		} else if tail1 != nil && tail2 != nil {
			if tail1.Val <= tail2.Val {
				tail.Next = &ListNode{Val: tail1.Val}
				tail1 = tail1.Next
			} else {
				tail.Next = &ListNode{Val: tail2.Val}
				tail2 = tail2.Next
			}
			tail = tail.Next
		} else if tail1 != nil {
			tail.Next = &ListNode{Val: tail1.Val}
			tail = tail.Next
			tail1 = tail1.Next
		} else {
			tail.Next = &ListNode{Val: tail2.Val}
			tail = tail.Next
			tail2 = tail2.Next
		}
	}

	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
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

func ListNodeToSlice(ln *ListNode) []int {
	lst := make([]int, 0)
	for ln != nil {
		lst = append(lst, ln.Val)
		ln = ln.Next
	}
	return lst
}
