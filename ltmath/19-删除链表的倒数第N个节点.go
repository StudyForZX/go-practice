package ltmath

func LT19_RemoveNthFromEndByStack(head *ListNode, n int) *ListNode {

	if head == nil || (head.Next == nil && n == 1) {
		return nil
	}

	tmp := head
	stack := []*ListNode{}

	for tmp != nil {

		stack = append(stack, tmp)

		tmp = tmp.Next
	}

	preIndex := len(stack) - n - 1

	if preIndex == -1 {
		head = head.Next
		return head
	}

	goalPre := stack[len(stack)-n-1]
	goalPre.Next = goalPre.Next.Next

	return head
}

func LT19_RemoveNthFromEndByFor(head *ListNode, n int) *ListNode {

	if head == nil || (head.Next == nil && n == 1) {
		return nil
	}

	count := 0
	length := 0
	tmp := head

	for tmp != nil {
		length++
		tmp = tmp.Next
	}

	goal := length - n
	tmp = head

	if goal == 0 {
		head = tmp.Next
		return head
	}

	for tmp != nil {

		count++

		if count == goal {
			tmp.Next = tmp.Next.Next
			break
		}

		tmp = tmp.Next
	}

	return head
}
