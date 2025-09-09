package ltmath

// 使用map逐一对比
func LT160GetIntersectionNodeByMyself(headA, headB *ListNode) *ListNode {

	mapHead := map[*ListNode]struct{}{}

	for headA != nil {

		mapHead[headA] = struct{}{}

		headA = headA.Next

	}

	for headB != nil {
		if _, ok := mapHead[headB]; ok {
			return headB
		}
		headB = headB.Next
	}

	return nil

}

// 双指针，如果不相交，则最后都为nil，退出
func LT160GetIntersectionNodeByDoublePointer(headA, headB *ListNode) *ListNode {

	lA, lB := headA, headB

	for lA != lB {

		if lA == nil {
			lA = headB
		} else {
			lA = lA.Next
		}

		if lB == nil {
			lB = headA
		} else {
			lB = lB.Next
		}

	}

	return lA
}
