package ltmath

func LT45Jump(nums []int) int {

	if len(nums) < 2 {
		return 0
	}

	end := 0
	times := 0
	maxIndex := 0

	for i, num := range nums {

		maxIndex = max(maxIndex, i+num)

		if i == end {
			times++
			end = maxIndex

			if end >= len(nums)-1 {
				break
			}
		}
	}

	return times

}
