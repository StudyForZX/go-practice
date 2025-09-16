package ltmath

func LT128LongestConsecutive(nums []int) int {

	maxRes := 0

	numsMap := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		numsMap[num] = struct{}{}
	}

	for num := range numsMap {

		// 如果存在更小的数，那就不计算当前这个数
		// 因为计算最小的数的时候，已经包含了这个数
		if _, ok := numsMap[num-1]; !ok {

			count := 1

			for {
				if _, ok := numsMap[num+1]; ok {
					count++
					num++
				} else {
					break
				}
			}

			maxRes = max(maxRes, count)
		}

	}

	return maxRes
}
