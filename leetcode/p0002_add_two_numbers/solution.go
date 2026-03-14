package p0002

func solve(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	ans := &ListNode{}
	head := ans

	for l1 != nil || l2 != nil || carry != 0 {
		a, b := 0, 0

		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}

		c := a + b + carry

		carry = c / 10
		c = c % 10

		head.Next = &ListNode{Val: c}
		head = head.Next
	}

	return ans.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func ArrayToListnode(a []int) *ListNode {
	dummy := &ListNode{}
	head := dummy
	for _, val := range a {
		head.Next = &ListNode{Val: val}
		head = head.Next
	}
	return dummy.Next
}

func ListnodeToArray(l *ListNode) []int {
	a := make([]int, 0)
	for l != nil {
		a = append(a, l.Val)
		l = l.Next
	}
	return a
}
