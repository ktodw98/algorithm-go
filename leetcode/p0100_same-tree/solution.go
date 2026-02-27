package p0100

var null = -10001

func solve(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q == nil || p == nil && q != nil {
		return false
	}

	return p.Val == q.Val && solve(p.Left, q.Left) && solve(p.Right, q.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SliceToTree(a []int) *TreeNode {
	/*Transform level traversal slice to TreeNode struct*/
	if len(a) == 0 {
		return nil
	}

	head := &TreeNode{Val: a[0]}

	q := []*TreeNode{}
	q = append(q, head)

	isLeft := true
	for _, v := range a {
		if v == null {
			if isLeft {
				isLeft = false
			} else {
				isLeft = true
				if len(q) > 1 {
					q = q[1:]
				} else {
					q = []*TreeNode{}
				}
			}
			continue
		}

		parentNode := q[0]

		nowNode := &TreeNode{Val: v}
		q = append(q, nowNode)

		if isLeft {
			parentNode.Left = nowNode
			isLeft = false
		} else {
			parentNode.Right = nowNode
			isLeft = true
			if len(q) > 1 {
				q = q[1:]
			} else {
				q = []*TreeNode{}
			}
		}
	}

	return head
}
