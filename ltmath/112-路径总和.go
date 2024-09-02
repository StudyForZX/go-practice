package ltmath

// 深度优先
func LT112_PathSum_DFS(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}

	resLeft := LT112_PathSum_DFS(root.Left, targetSum-root.Val)
	resRight := LT112_PathSum_DFS(root.Right, targetSum-root.Val)

	return resLeft || resRight
}

// 广度优先
func LT112_PathSum_BFS(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false
	}

	queNode, queVal := []*TreeNode{}, []int{}
	queNode = append(queNode, root)
	queVal = append(queVal, root.Val)

	for len(queNode) != 0 {

		now, tmp := queNode[0], queVal[0]
		queNode, queVal = queNode[1:], queVal[1:]

		if now.Left == nil && now.Right == nil {
			if tmp == targetSum {
				return true
			}

			continue
		}

		if now.Left != nil {
			queNode = append(queNode, now.Left)
			queVal = append(queVal, now.Left.Val+tmp)
		}

		if now.Right != nil {
			queNode = append(queNode, now.Right)
			queVal = append(queVal, now.Right.Val+tmp)
		}
	}

	return false
}
