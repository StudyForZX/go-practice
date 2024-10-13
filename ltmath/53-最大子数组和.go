package ltmath

func LT53_MaxSubArray(nums []int) int {

	if len(nums) < 2 {
		return nums[0]
	}

	sum := 0
	maxRes := nums[0]

	for _, num := range nums {

		if sum > 0 {
			sum = sum + num
		} else {
			sum = num
		}

		maxRes = max(sum, maxRes)
	}

	return maxRes
}
