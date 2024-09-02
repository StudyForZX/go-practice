package ltmath

func PostorderTraversal(root *TreeNode) []int {

	res := []int{}

	res = seek(root, res)

	return res
}

// 后序遍历
func seek(node *TreeNode, res []int) []int {

	if node == nil {
		return res
	}

	// 先遍历左节点
	res = seek(node.Left, res)
	// 再遍历右节点
	res = seek(node.Right, res)
	// 最后遍历根节点
	res = append(res, node.Val)

	return res
}
