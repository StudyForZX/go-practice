package ltmath

func LT25ReverseKGroup(head *ListNode, k int) *ListNode {

	tmpHead := &ListNode{Next: head}
	pre := tmpHead

	for {

		tail := pre
		for range k {
			tail = tail.Next
			if tail == nil {
				return tmpHead.Next
			}
		}

		next := tail.Next
		tail.Next = nil

		newHead := reverseListNode(pre.Next)
		oldHead := pre.Next

		pre.Next = newHead
		oldHead.Next = next

		pre = oldHead

	}

}

func reverseListNode(head *ListNode) *ListNode {

	var pre *ListNode

	for head != nil {

		next := head.Next
		head.Next = pre
		pre = head
		head = next

	}

	return pre

}
