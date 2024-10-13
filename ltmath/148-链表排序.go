package ltmath

import (
	"sort"
)

// 转数组排序方式
func LT148_SortListByTransToArray(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	nums := []int{}

	for head != nil {

		nums = append(nums, head.Val)

		head = head.Next
	}

	sort.Ints(nums)

	current := &ListNode{Val: nums[0]}
	tmp := current

	for i := 1; i < len(nums); i++ {
		tmp.Next = &ListNode{Val: nums[i]}
		tmp = tmp.Next
	}

	return current
}
