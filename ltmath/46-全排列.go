package ltmath

func LT46_Permute(nums []int) [][]int {

	res := [][]int{}
	combine := []int{}
	used := map[int]bool{}

	var dfs func(index int)
	dfs = func(index int) {

		for _, num := range nums {

			if index == len(nums) {
				res = append(res, append([]int{}, combine...))
				return
			}

			if !used[num] {
				combine = append(combine, num)
				used[num] = true
				dfs(index + 1)
				used[num] = false
				combine = combine[:len(combine)-1]
			}

		}

	}

	dfs(0)

	return res
}
