package p0094

var null = -101

func solve(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	lList := solve(root.Left)
	nowVal := root.Val
	rList := solve(root.Right)

	a := []int{}
	for _, v := range lList {
		a = append(a, v)
	}

	a = append(a, nowVal)

	for _, v := range rList {
		a = append(a, v)
	}
	return a
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SliceToLevelTree(a []int) *TreeNode {
	/*Pop Node from queue and Link L/R from slice*/

	// Null Case
	if len(a) < 1 {
		return nil
	}

	// Root Node
	root := &TreeNode{Val: a[0]}

	q := make([]*TreeNode, 0)
	q = append(q, root)

	i := 1
	r := 0
	for q != nil && i < len(a) {
		now := q[0]
		// Skip Null Node
		if a[i] == null {
			if r == 1 {
				if len(q) > 1 {
					q = q[1:]
				} else {
					break
				}
			}
			i++
			r = (r + 1) % 2

			continue
		}
		// Link Child Node
		if r == 0 {
			Node := &TreeNode{Val: a[i]}
			now.Left = Node
			q = append(q, Node)
		} else {
			Node := &TreeNode{Val: a[i]}
			now.Right = Node
			q = append(q, Node)

			if len(q) > 1 {
				q = q[1:]
			} else {
				break
			}
		}
		i++
		r = (r + 1) % 2
	}
	return root
}
