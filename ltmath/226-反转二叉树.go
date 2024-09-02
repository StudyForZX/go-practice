package ltmath

func LT226_InvertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	left := LT226_InvertTree(root.Left)
	right := LT226_InvertTree(root.Right)

	root.Right = left
	root.Left = right

	return root
}
