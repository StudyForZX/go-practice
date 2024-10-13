package ltmath

func LT102_LevelOrder(root *TreeNode) [][]int {

	res := [][]int{}

	if root == nil {
		return res
	}

	q := []*TreeNode{root}

	for len(q) > 0 {

		size := len(q)

		tmp := []int{}

		for size > 0 {

			node := q[0]
			q = q[1:]

			tmp = append(tmp, node.Val)

			if node.Left != nil {

				q = append(q, node.Left)

			}

			if node.Right != nil {

				q = append(q, node.Right)

			}

			size--
		}

		res = append(res, tmp)
	}

	return res
}
