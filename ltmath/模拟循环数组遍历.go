package ltmath

func MockArrayForeach(nums []int) [][]int {

	res := [][]int{}
	tmp := make([]int, len(nums))

	var dfs func(index int)
	dfs = func(index int) {

		if index == len(nums) {
			res = append(res, append([]int{}, tmp...))
			return
		}

		for i := 0; i < nums[index]; i++ {
			tmp[index] = i
			dfs(index + 1)
		}

	}

	dfs(0)

	return res
}

// func MockArrayForeach(arr []int) [][]int {

// 	arrLen := len(arr)
// 	res := [][]int{}

// 	for i := 0; i < arrLen; i++ {
// 		res = MockArrayForeachHandle(arr[0], arr[i], res)
// 	}

// 	return res
// }

// func MockArrayForeachHandle(first int, max int, res [][]int) [][]int {

// 	tmp := [][]int{}

// 	if len(res) < first {

// 		for i := 0; i < max; i++ {
// 			tmp = append(tmp, []int{i})
// 		}

// 	} else {

// 		for _, item := range res {
// 			for i := 0; i < max; i++ {
// 				tmp = append(tmp, append(item, i))
// 			}
// 		}
// 	}

// 	return tmp
// }
