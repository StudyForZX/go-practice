package ltmath

var first, second int

func TwoSum(nums []int, target int) (int, int) {

	for i := 0; i < len(nums); i++ {

		diff := target - nums[i]

		for j := 1; j < len(nums); j++ {

			if diff == nums[j] {

				first, second = i, j

				break

			}
		}
	}

	return first, second

}
