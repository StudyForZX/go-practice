package ltmath

func LT189_RotateByExchange(nums []int, k int) []int {

	newNums := make([]int, len(nums))

	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}

	copy(nums, newNums)

	return nums
}

func LT189_RotateBySlice(nums []int, k int) []int {

	if k > len(nums) {
		k = k % len(nums)
	}

	numsLen := len(nums)

	numsLeft := nums[:numsLen-k]
	numsRight := nums[numsLen-k:]

	res := append([]int{}, numsRight...)
	res = append(res, numsLeft...)

	return res
}
