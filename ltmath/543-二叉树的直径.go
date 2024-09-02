package ltmath

func LT543_DiameterOfBinaryTree(root *TreeNode) int {

	var res = 0

	var depth func(*TreeNode) int

	depth = func(node *TreeNode) int {

		if node == nil {
			return 0
		}

		left := depth(node.Left)
		right := depth(node.Right)
		res = max(res, left+right)
		return max(left, right) + 1
	}

	depth(root)

	return res
}
