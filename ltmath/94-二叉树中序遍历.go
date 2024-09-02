package ltmath

func InorderTraversal(root *TreeNode) []int {

	res := []int{}

	res = seekInorderTraversal(root, res)

	return res
}

// 中序遍历
func seekInorderTraversal(node *TreeNode, res []int) []int {

	if node == nil {
		return res
	}
	// 先遍历左节点
	res = seekInorderTraversal(node.Left, res)
	// 再遍历根节点
	res = append(res, node.Val)
	// 最后遍历右节点
	res = seekInorderTraversal(node.Right, res)

	return res
}
