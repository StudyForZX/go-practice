package practice

func PracticeAppendScliceAndRetrunIndex(n int) int {

	if n == 1 {
		return 1
	}

	k := 0
	nums := []int{1}

	for len(nums) < n {
		nums = append(nums, 2*nums[k]+1, 3*nums[k]+1)
		k++
	}

	return nums[n-1]
}
