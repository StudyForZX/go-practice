package ltmath

func PathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := rootSum(root, targetSum)
	res += PathSum(root.Left, targetSum)
	res += PathSum(root.Right, targetSum)
	return res
}

func rootSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return
	}
	val := root.Val
	if val == targetSum {
		res++
	}
	res += rootSum(root.Left, targetSum-val)
	res += rootSum(root.Right, targetSum-val)
	return
}
