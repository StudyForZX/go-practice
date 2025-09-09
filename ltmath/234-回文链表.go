package ltmath

func LT234IsPalindromeByMyself(head *ListNode) bool {

	if head == nil {
		return true
	}

	vals := []int{}

	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}

	left := 0
	right := len(vals) - 1

	for left < right {

		if vals[left] != vals[right] {
			return false
		}

		left++
		right--

		if left == right {
			return true
		}

	}

	return true

}

// 递归实现
// func LT234IsPalindromeByRecursion(head *ListNode) bool {

// }
