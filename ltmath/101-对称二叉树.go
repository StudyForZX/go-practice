package ltmath

func LT101_IsSymmetric(root *TreeNode) bool {
	return check(root.Left, root.Right)
}

func check(left, right *TreeNode) bool {

	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && check(left.Left, right.Right) && check(left.Right, right.Left)
}
