package ltmath

// 迭代反转链表
// func LT206_ReverseListByFor(head *ListNode) *ListNode {

// 	var pre *ListNode = nil
// 	current := head

// 	for current != nil {

// 		next := current.Next

// 		current.Next = pre

// 		pre = current

// 		current = next
// 	}

// 	return head

// }

// 迭代反转链表
func LT206_ReverseListByFor(head *ListNode) *ListNode {

	var current *ListNode

	for head != nil {
		next := head.Next
		head.Next = current
		current = head
		head = next
	}

	return current
}

// 递归写法
func LT206_ReverseListByRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 这一步很精妙，每一次newHead都是指向空指针（链表为空）或保留在原链表中的最后一个节点（链表不空），作用就是返回新的头结点
	newHead := LT206_ReverseListByRecursion(head.Next)
	// 最后一个节点指向倒数第二个节点
	head.Next.Next = head
	// 倒数第二个节点的下一节点置空。此时倒数第三个节点仍指向倒数第二个节点，下一次递归中将倒数第二个节点下一节点指向倒数第三个节点，不断重复这一过程
	head.Next = nil

	return newHead
}
