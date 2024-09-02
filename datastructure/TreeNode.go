package datastructure

import "fmt"

/**
 * 二叉树结构
 */
type TreeNode[T comparable] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

/**
 * 将数组按层序遍历转为二叉树
 *
 * 二叉树性质（i代表第几层 i >= 1）：
 * 二叉树第i层上的结点数目最多为：2*(i-1)
 * 二叉树深度为i时，总节点数量最多为：2*i - 1
 */
func (node *TreeNode[T]) ArrayToTreeNode(arrayT []T, nilValue T) *TreeNode[T] {

	if len(arrayT) == 0 {
		return nil
	}

	// 创建根节点
	root := &TreeNode[T]{Val: arrayT[0]}

	index := 1
	quene := []*TreeNode[T]{root}

	for index < len(arrayT) {

		node := quene[0]
		quene = quene[1:]

		// 创建左子树
		if index < len(arrayT) && arrayT[index] != nilValue {
			node.Left = &TreeNode[T]{Val: arrayT[index]}
			quene = append(quene, node.Left)
		}

		index++

		// 创建右子数
		if index < len(arrayT) && arrayT[index] != nilValue {
			node.Right = &TreeNode[T]{Val: arrayT[index]}
			quene = append(quene, node.Right)
		}

		index++
	}

	return root
}

/**
 * 中序遍历
 *
 * 先遍历左节点，再遍历根节点，最后遍历右节点
 */
func (node *TreeNode[T]) InorderTraversal() {

	if node == nil {
		return
	}
	//
	node.Left.InorderTraversal()
	fmt.Print(node.Val, " ")
	node.Right.InorderTraversal()
}
