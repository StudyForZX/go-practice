package ltmath

func LT2AddTwoNumbersByMyself(l1 *ListNode, l2 *ListNode) *ListNode {

	tmp := []int{}
	carray := 0
	last := 0

	for l1 != nil || l2 != nil {

		v1 := 0
		v2 := 0

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		remainder := (v1 + v2) % 10
		if carray > 0 {
			remainder = remainder + carray
			if remainder == 10 {
				remainder = 0
				last = 1
			}
		}

		tmp = append(tmp, remainder)

		carray = (v1 + v2) / 10
		if last != 0 {
			carray = carray + last
		}
	}

	if last != 0 {
		tmp = append(tmp, last)
	}

	resListHead := &ListNode{
		Val:  tmp[0],
		Next: nil,
	}

	tmpResListHead := resListHead

	for i := 1; i < len(tmp); i++ {
		tmpResListHead.Next = &ListNode{
			Val:  tmp[i],
			Next: nil,
		}

		tmpResListHead = tmpResListHead.Next
	}

	return resListHead

}

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
