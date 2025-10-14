package ltmath

func LT55CanJump(nums []int) bool {

	maxIndex := 0

	for i, num := range nums {

		if i > maxIndex {
			return false
		}

		maxIndex = max(maxIndex, i+num)

		if maxIndex >= len(nums)-1 {
			return true
		}
	}

	return false
}
