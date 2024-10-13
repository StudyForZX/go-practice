package ltmath

func LT128_LongestConsecutive(nums []int) int {

	maxRes := 0

	numsMap := map[int]bool{}

	for _, num := range nums {
		numsMap[num] = true
	}

	for _, num := range nums {

		if !numsMap[num-1] {

			count := 1

			for numsMap[num+1] {
				count++
				num = num + 1
			}

			maxRes = max(maxRes, count)
		}

	}

	return maxRes
}
