package ltmath

func LT104_MaxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return max(LT104_MaxDepth(root.Left), LT104_MaxDepth(root.Right)) + 1
}
