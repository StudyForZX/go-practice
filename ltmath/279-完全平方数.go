package ltmath

import "math"

func LT279NumSquares(n int) int {
	nSlice := make([]int, n+1)

	for i := 1; i <= n; i++ {
		minNum := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minNum = min(minNum, nSlice[i-j*j])
		}

		nSlice[i] = minNum + 1
	}

	return nSlice[n]
}
