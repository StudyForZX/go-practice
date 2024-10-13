package ltmath

func LT238_ProductExceptSelfByFor(nums []int) []int {

	numsLen := len(nums)
	numsMap := map[int]bool{}
	res := make([]int, numsLen)

	for i := 0; i < len(nums); i++ {

		tmp := 1
		numsMap[i] = true

		for j := 0; j < len(nums); j++ {

			if !numsMap[j] {
				tmp = tmp * nums[j]
			}

			if j == len(nums)-1 {
				res[i] = tmp
			}
		}

		numsMap[i] = false
	}

	return res
}
