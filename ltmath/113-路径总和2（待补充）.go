package ltmath

// DFS按照root从上到下返回
func LT113_PathSum_DFS(root *TreeNode, targetSum int) [][]int {

	var path []int
	var res [][]int

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, targetSum int) {

		if node == nil {
			return
		}

		path = append(path, node.Val)

		if node.Left == nil && node.Right == nil && targetSum == node.Val {
			res = append(res, append([]int{}, path...))
		} else {
			dfs(node.Left, targetSum-node.Val)
			dfs(node.Right, targetSum-node.Val)
		}

		path = path[:len(path)-1]
	}

	dfs(root, targetSum)

	return res
}

// 按照广度优先

// type pair struct {
// 	node *TreeNode
// 	sum int
// }

// func LT113_PathSum_BFS(root *TreeNode, targetSum int) [][]int {

// 	if root == nil {
// 		return nil
// 	}

// 	parent := map[*TreeNode]*TreeNode{}

// 	getPath := func(node *TreeNode) []int {

// 		path = []int{}

// 		for node!=nil {
// 			path = append(path, node.Val)
// 			node = parent[node]
// 		}

// 		i, j := 0, len(path)-1

// 		for i < j {

// 			path[i], path[j] = path[j], path[i]

// 			i++
// 			j--
// 		}

// 		return
// 	}

// 	sum := 0
// 	queneTN := []*TreeNode{}
// 	queneVal := []int{}

// 	queneTN = append(queneTN, root)
// 	queneVal = append(queneVal, root.Val)

// 	for len(queneTN) > 0 {

// 		node := queneTN[0]
// 		nodeVal := queneVal[0]
// 		queneTN = queneTN[1:]
// 		queneVal = queneVal[1:]

// 		sum += nodeVal
// 		path = append(path, node.Val)

// 		if node.Left == nil && node.Right == nil && sum == targetSum {
// 			res = append(res, append([]int{}, path...))
// 		}

// 		if node.Left != nil {
// 			queneTN = append(queneTN, node.Left)
// 			queneVal = append(queneVal, sum+node.Left.Val)
// 		}

// 		if node.Right != nil {
// 			queneTN = append(queneTN, node.Right)
// 			queneVal = append(queneVal, sum+node.Right.Val)
// 		}

// 	}

// 	retrun res
// }
