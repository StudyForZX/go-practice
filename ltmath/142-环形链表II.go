package ltmath

func LT142_DetectCycleByDoublePoint(head *ListNode) *ListNode {

	slow, fast := head, head

	for fast != nil {

		slow = slow.Next

		if fast.Next == nil {
			return nil
		}

		fast = fast.Next

		if fast == slow {

			p := head

			for p != slow {
				p = p.Next
				slow = slow.Next
			}

			return p
		}
	}

	return nil
}

func LT142_DetectCycleByMap(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return nil
	}

	var res *ListNode
	listMap := map[*ListNode]bool{}

	for head != nil {

		if listMap[head] {
			res = head
			break
		}

		listMap[head] = true

		head = head.Next
	}

	return res
}
