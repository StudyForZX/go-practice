package ltmath

// 使用map逐一对比
func LT160_GetIntersectionNode_Hash(headA, headB *ListNode) *ListNode {

	headAMap := map[*ListNode]bool{}

	for headA != nil {
		headAMap[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := headAMap[headB]; ok {
			return headB
		}

		headB = headB.Next
	}

	return nil
}

func LT160_GetIntersectionNode_DoublePointer(headA, headB *ListNode) *ListNode {

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
