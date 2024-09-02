package ltmath

// func LT128_LongestConsecutive(nums []int) int {

// 	max := 0
// 	numsMap := map[int]bool{}

// 	for _, num := range nums {
// 		numsMap[num] = true
// 	}

// 	for _, num := range nums {

// 		if numsMap[num-1] {
// 			continue
// 		}

// 		count := 1

// 		for nextNum := num + 1; numsMap[nextNum]; nextNum++ {
// 			count++
// 		}

// 		if count > max {
// 			max = count
// 		}
// 	}

// 	return max
// }

func LT128_LongestConsecutive(nums []int) int {

	numSet := map[int]bool{}

	for _, num := range nums {
		numSet[num] = true
	}

	longListNo := 0

	for num := range numSet {

		if !numSet[num-1] {

			currentNum := num
			currentlongListNo := 1

			for numSet[currentNum+1] {
				currentNum++
				currentlongListNo++
			}

			if longListNo < currentlongListNo {
				longListNo = currentlongListNo
			}
		}
	}

	return longListNo
}
