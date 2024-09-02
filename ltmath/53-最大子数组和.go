package ltmath

// func LT53_MaxSubArray(nums []int) int {
// 	pre := 0
// 	maxRes := nums[0]

// 	for _, num := range nums {
// 		pre = max(pre+num, num)
// 		maxRes = max(maxRes, pre)
// 	}

// 	return maxRes
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

func LT53_MaxSubArray(nums []int) int {

	max := nums[0]
	pre := nums[0]

	for i := 1; i < len(nums); i++ {

		if pre > 0 {
			pre = pre + nums[i]
		} else {
			pre = nums[i]
		}

		if pre > max {
			max = pre
		}
	}

	return max
}
