package common

// 取多数中的最大值
func Max(nums ...int) int {
	var max int
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

// 取多数中的最小值
func Min(nums ...int) int {
	var min int
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}
