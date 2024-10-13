package ltmath

func LT54_SpiraOrder(matrix [][]int) []int {

	if len(matrix) == 0 {
		return []int{}
	}

	m, n := 0, 0
	res := []int{}
	m = len(matrix)
	n = len(matrix[0])

	matrixMap := make([][]bool, m)
	for i := range matrixMap {
		matrixMap[i] = make([]bool, n)
	}

	startM := 0
	startN := 0

	for count := 0; count < m*n; count++ {

		if !matrixMap[startM][startN] {

			res = append(res, matrix[startM][startN])

			matrixMap[startM][startN] = true

		}

		if startN < n-1 && !matrixMap[startM][startN] {

			if startM <= m-1 {
				startM--
			} else {
				startN++
			}

		} else {
			if startM == m-1 {
				startN--
			} else {
				startM++
			}
		}

	}

	return res
}
