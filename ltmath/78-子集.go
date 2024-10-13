package ltmath

func LT78_SubsetsByRecursion(nums []int) [][]int {

	res := [][]int{}

	tmp := []int{}

	var dfs func(index int)
	dfs = func(index int) {

		if index == len(nums) {
			res = append(res, append([]int{}, tmp...))
			return
		}

		tmp = append(tmp, nums[index])
		dfs(index + 1)
		tmp = tmp[:len(tmp)-1]
		dfs(index + 1)
	}

	dfs(0)
	return res
}

func LT78_SubsetsByFor(nums []int) [][]int {

	res := [][]int{
		{},
	}

	for i := range len(nums) {
		for j := range len(res) {
			tmp := append([]int{}, res[j]...)
			tmp = append(tmp, nums[i])
			res = append(res, tmp)
		}
	}

	return res
}
