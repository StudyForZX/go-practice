package ltmath

func Test_LT206_ReversePart(head *ListNode, m, n int) *ListNode {

	var pre, leftHead, leftTail, rightHead *ListNode
	res := head
	count := 0

	for head != nil {

		if count == m-1 {
			leftTail = head
		}

		if count == n+1 {
			rightHead = head
		}

		if count < m || count > n {
			head = head.Next
			count++
			continue
		}

		if count == m {
			leftHead = head
		}

		next := head.Next
		head.Next = pre
		pre = head
		head = next

		count++
	}

	leftTail.Next = pre
	leftHead.Next = rightHead

	return res
}
