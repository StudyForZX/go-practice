package ltmath

func LT46Permute(nums []int) [][]int {

	res := [][]int{}
	tmp := []int{}
	numsMap := make(map[int]bool, len(nums))

	var dfs func()
	dfs = func() {

		if len(nums) == len(tmp) {
			res = append(res, append([]int{}, tmp...))
			return
		}

		for _, num := range nums {

			if !numsMap[num] {
				tmp = append(tmp, num)
				numsMap[num] = true
				dfs()
				numsMap[num] = false
				tmp = tmp[:len(tmp)-1]
			}

		}

	}

	dfs()

	return res

}
