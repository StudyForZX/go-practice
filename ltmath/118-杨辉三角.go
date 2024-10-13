package ltmath

// 写法优雅
func LT118_GenerateByLeetCode(numsRow int) [][]int {

	res := make([][]int, numsRow)

	for i := range res {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][i] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}

	return res
}

// 第一次写的写法
func LT118_Generate(numRows int) [][]int {

	res := make([][]int, numRows)

	if numRows == 1 {
		return [][]int{
			{1},
		}
	}

	if numRows == 2 {
		return [][]int{
			{1},
			{1, 1},
		}
	}

	res[0] = []int{1}
	res[1] = []int{1, 1}

	for i := 2; i < numRows; i++ {

		tmp := []int{1}

		for j := 0; j < len(res[i-1])-1; j++ {

			sumNum := res[i-1][j] + res[i-1][j+1]

			tmp = append(tmp, sumNum)

		}

		tmp = append(tmp, 1)

		res[i] = tmp

	}

	return res
}
