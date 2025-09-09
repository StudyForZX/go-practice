package ltmath

func LT141HasCycleByMyself(head *ListNode) bool {

	listPointMap := map[*ListNode]struct{}{}

	for head != nil {

		if _, ok := listPointMap[head]; ok {
			return true
		}

		listPointMap[head] = struct{}{}

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
