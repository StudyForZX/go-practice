package ltmath

// 每个窗口都比较一次
func LT239MaxSlidingWindowByMyself(nums []int, k int) []int {

	left := 0
	right := k - 1
	resMax := []int{}
	numsLen := len(nums)

	for right < numsLen {

		tmpMax := nums[left]

		for i := left; i <= right; i++ {
			tmpMax = max(nums[i], tmpMax)
		}

		resMax = append(resMax, tmpMax)

		left++
		right++

	}

	return resMax

}

// 如果max还在下一个窗口 就跳过循环 直接对比新加入的数字跟最大值
func LT239MaxSlidingWindowByMyself2(nums []int, k int) []int {

	left := 0
	right := k - 1
	resMax := []int{}
	numsLen := len(nums)
	needFor := true

	for right < numsLen {

		tmpMax := nums[left]
		tmpMaxIndex := left

		if tmpMaxIndex > left {
			needFor = false
		} else {
			needFor = true
		}

		if needFor {
			for i := left; i <= right; i++ {
				if nums[i] > tmpMax {
					tmpMax = nums[i]
					tmpMaxIndex = i
				}
			}
		} else {
			tmpMax = max(nums[tmpMaxIndex], nums[right])
		}

		resMax = append(resMax, tmpMax)

		left++
		right++

	}

	return resMax

}

// 官方题解
// TODO: 看大顶堆
// func LT239MaxSlidingWindow(nums []int, k int) []int {

// }
