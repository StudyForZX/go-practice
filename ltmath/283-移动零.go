package ltmath

func LT283MoveZeroesByMyself(nums []int) []int {

	if len(nums) < 2 {
		return nums
	}

	left := 0
	right := 0

	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}

		right++
	}

	return nums

}

func LT283MoveZeroes(nums []int) []int {
	// 记录非零元素应该存放的位置
	// 其实就是第一个0的位置
	mark := 0

	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[mark] = nums[mark], nums[i]
			mark++
		}
	}

	return nums
}
