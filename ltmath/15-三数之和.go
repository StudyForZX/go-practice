package ltmath

import "sort"

// 暴力破解 + map
// func LT15_ThreeSum(nums []int) [][]int {

// 	res := [][]int{}
// 	tmp := map[string][]int{}

// 	sort.Ints(nums)

// 	numsLen := len(nums)

// 	for i := 0; i < numsLen; i++ {
// 		for j := i + 1; j < numsLen; j++ {
// 			for k := j + 1; k < numsLen; k++ {
// 				if nums[i]+nums[j]+nums[k] == 0 {
// 					index := fmt.Sprintf("%d-%d-%d", nums[i], nums[j], nums[k])
// 					tmp[index] = []int{nums[i], nums[j], nums[k]}
// 				}
// 			}
// 		}
// 	}

// 	for _, item := range tmp {
// 		res = append(res, item)
// 	}

// 	return res
// }

// 暴力破解 + 不适用map
// func LT15_ThreeSum(nums []int) [][]int {

// 	res := [][]int{}
// 	sort.Ints(nums)

// 	numsLen := len(nums)

// 	for i := 0; i < numsLen; i++ {

// 		if i > 0 && nums[i] == nums[i-1] {
// 			continue
// 		}

// 		for j := i + 1; j < numsLen; j++ {

// 			if j > i+1 && nums[j] == nums[j-1] {
// 				continue
// 			}

// 			for k := j + 1; k < numsLen; k++ {

// 				if k > j+1 && nums[k] == nums[k-1] {
// 					continue
// 				}

// 				if nums[i]+nums[j]+nums[k] == 0 {
// 					res = append(res, []int{nums[i], nums[j], nums[k]})
// 				}
// 			}
// 		}
// 	}

// 	return res
// }

// n^2
func LT15_ThreeSum(nums []int) [][]int {

	res := [][]int{}
	sort.Ints(nums)

	numsLen := len(nums)

	for i := 0; i < numsLen; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		k := numsLen - 1
		for j := i + 1; j < numsLen; j++ {

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for j < k && nums[i]+nums[j]+nums[k] > 0 {
				k--
			}

			if j == k {
				break
			}

			if nums[i]+nums[j]+nums[k] == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
			}
		}
	}

	return res
}
