package ltmath

func PreorderTraversal(root *TreeNode) []int {

	res := []int{}

	res = seekPreorderTraversal(root, res)

	return res
}

// 先序遍历
func seekPreorderTraversal(node *TreeNode, res []int) []int {

	if node == nil {
		return res
	}

	// 先遍历根节点
	res = append(res, node.Val)
	// 再遍历左节点
	res = seekPreorderTraversal(node.Left, res)
	// 最后遍历右节点
	res = seekPreorderTraversal(node.Right, res)

	return res
}
