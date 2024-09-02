package ltmath

func MockArrayForeach(arr []int) [][]int {

	arrLen := len(arr)
	res := [][]int{}

	for i := 0; i < arrLen; i++ {
		res = MockArrayForeachHandle(arr[0], arr[i], res)
	}

	return res
}

func MockArrayForeachHandle(first int, max int, res [][]int) [][]int {

	tmp := [][]int{}

	if len(res) < first {

		for i := 0; i < max; i++ {
			tmp = append(tmp, []int{i})
		}

	} else {

		for _, item := range res {
			for i := 0; i < max; i++ {
				tmp = append(tmp, append(item, i))
			}
		}
	}

	return tmp
}
