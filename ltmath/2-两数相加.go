package ltmath

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var head, next *ListNode
	carry := 0

	for l1 != nil || l2 != nil {

		n1, n2 := 0, 0

		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10

		if head == nil {
			head = &ListNode{Val: sum}
			next = head
		} else {
			next.Next = &ListNode{Val: sum}
			next = next.Next
		}
	}

	if carry > 0 {
		next.Next = &ListNode{Val: carry}
	}

	return head
}
