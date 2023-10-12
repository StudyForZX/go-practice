package ltmath

import "github.com/studyforzx/lt/common"

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func AddTwoNumbers(l1, l2 *ListNode) *ListNode {

// }

func AddTwoNumbers(l1, l2 []int) []int {

	maxLen := common.Max(len(l1), len(l2))

	var res = make([]int, maxLen)

	for i := 0; i < maxLen; i++ {
		res[i] = l1[i] + l2[i]
	}

	for i := 0; i < maxLen; i++ {

		if res[i] < 10 {
			continue
		}

		res[i] = res[i] - 10

		if i == maxLen-1 {
			res = append(res, res[i]+1)
		} else {
			res[i+1] = res[i+1] + 1
		}

	}

	if res[maxLen-1] > 10 {
		res[maxLen-1] = res[maxLen-1] - 10
		res = append(res, 1)
	}

	return res
}
