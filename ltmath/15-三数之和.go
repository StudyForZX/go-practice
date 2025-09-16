package ltmath

import (
	"sort"
)

func LT15ThreeSumByMyselfWithN3(nums []int) [][]int {

	sort.Ints(nums)

	numsLen := len(nums)

	res := [][]int{}

	for i := 0; i < numsLen-2; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < numsLen-1; j++ {

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for k := j + 1; k < numsLen; k++ {

				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}

				if nums[i]+nums[j]+nums[k] == 0 {
					// res = append(res, []int{i, j, k})
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}

	return res

}

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
