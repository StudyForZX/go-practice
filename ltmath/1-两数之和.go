package ltmath

// func LT1_TwoSum(nums []int, target int) (int, int) {

//  var first, second int

// 	for i := 0; i < len(nums); i++ {

// 		diff := target - nums[i]

// 		for j := 1; j < len(nums); j++ {

// 			if diff == nums[j] {

// 				first, second = i, j

// 				break

// 			}
// 		}
// 	}

// 	return first, second
// }

func LT1_TwoSum(nums []int, target int) []int {

	numsMap := map[int]int{}

	for k, num := range nums {
		if _, ok := numsMap[target-num]; ok {
			return []int{numsMap[target-num], k}
		}

		numsMap[num] = k
	}

	return nil
}

// func LT1_TwoSum(nums []int, target int) []int {

// 	index1, index2 := 0, 0
// 	numsMap := map[int]int{}

// 	for k, num := range nums {
// 		numsMap[num] = k
// 	}

// 	for k, num := range nums {

// 		diff := target - num

// 		if v, ok := numsMap[diff]; ok {
// 			index1 = k
// 			index2 = v

// 			if index1 == index2 {
// 				continue
// 			}

// 			break
// 		}
// 	}

// 	return []int{index1, index2}
// }
