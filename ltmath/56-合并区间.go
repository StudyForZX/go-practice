package ltmath

import "sort"

func LT56_Merge(intervals [][]int) [][]int {

	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}
	minStart, maxEnd := 0, 0

	for i := 0; i < len(intervals); i++ {

		start, end := intervals[i][0], intervals[i][1]

		if len(res) == 0 || res[len(res)-1][1] < start {
			res = append(res, []int{start, end})
		} else {

			if res[len(res)-1][0] > start {
				minStart = start
			} else {
				minStart = res[len(res)-1][0]
			}

			if res[len(res)-1][1] > end {
				maxEnd = res[len(res)-1][1]
			} else {
				maxEnd = end
			}

			res[len(res)-1] = []int{minStart, maxEnd}
		}
	}

	return res
}
