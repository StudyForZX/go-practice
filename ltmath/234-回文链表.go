package ltmath

// 笨方法 转数组 看首尾是否相同
func LT234_IsPalindrome(head *ListNode) bool {

	arr := []int{}
	tmp := head

	for tmp != nil {
		arr = append(arr, tmp.Val)
		tmp = tmp.Next
	}

	i := 0
	j := len(arr) - 1
	res := true

	for i < j {

		if arr[i] == arr[j] {
			i++
			j--
		} else {
			res = false
			break
		}

	}

	return res
}
