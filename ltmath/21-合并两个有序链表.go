package ltmath

func LT21_MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	current := &ListNode{}
	tmp := current

	for list1 != nil && list2 != nil {

		v1 := list1.Val
		v2 := list2.Val

		if v1 > v2 {
			tmp.Next = list2
			list2 = list2.Next
		} else {
			tmp.Next = list1
			list1 = list1.Next
		}

		tmp = tmp.Next
	}

	if list1 == nil {
		tmp.Next = list2
	}

	if list2 == nil {
		tmp.Next = list1
	}

	return current.Next
}
