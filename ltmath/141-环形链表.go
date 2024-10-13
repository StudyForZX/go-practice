package ltmath

func LT141_HasCycleByHash(head *ListNode) bool {

	m := map[*ListNode]bool{}

	for head != nil {

		if m[head] {
			return true
		}

		m[head] = true
		head = head.Next
	}

	return false
}

func LT141_HasCycleByStep(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != fast {

		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}
