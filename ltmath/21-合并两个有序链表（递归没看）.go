package ltmath

func LT21_MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	head := &ListNode{}
	prev := head

	for list1 != nil && list2 != nil {

		v1 := list1.Val
		v2 := list2.Val

		if v1 > v2 {
			prev.Next = list2
			list2 = list2.Next
		} else {
			prev.Next = list1
			list1 = list1.Next
		}

		prev = prev.Next
	}

	if list1 == nil {
		prev.Next = list2
	} else {
		prev.Next = list1
	}

	return head.Next
}
