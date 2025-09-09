package ltmath

func LT21MergeTwoListsByMyself(list1 *ListNode, list2 *ListNode) *ListNode {

	newList := &ListNode{}

	tmp := newList

	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			tmp.Next = list2
			list2 = list2.Next
		} else {
			tmp.Next = list1
			list1 = list1.Next
		}

		tmp = tmp.Next
	}

	if list1 != nil {
		tmp.Next = list1
	}

	if list2 != nil {
		tmp.Next = list2
	}

	return newList.Next

}
