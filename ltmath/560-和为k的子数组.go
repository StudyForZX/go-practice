package ltmath

/**
 * 暴力方式：得出所有子数组的索引，返回相加=k的数组的数量
 */

// func SubArraySum(nums []int, k int) int {

// 	count := 0

// 	for start := 0; start < len(nums); start++ {
// 		sum := 0
// 		for end := start; end >= 0; end-- {
// 			sum += nums[end]
// 			if sum == k {
// 				count++
// 			}
// 		}
// 	}

// 	return count
// }

func SubArraySum(nums []int, k int) int {

	count, preSum := 0, 0
	numsMap := map[int]int{}
	numsMap[0] = 1

	for i := 0; i < len(nums); i++ {

		preSum += nums[i]

		if _, ok := numsMap[preSum-k]; ok {
			count += numsMap[preSum-k]
		}

		numsMap[preSum]++
	}

	return count
}
