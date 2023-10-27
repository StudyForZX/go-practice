package ltmath

import "strings"

func MinWindows(s string, t string) string {

	var allC [][]int
	var tIndex map[byte][]int

	// 索引探测
	for i := 0; i < len(t); i++ {

		tmpIndex := strings.IndexByte(s, t[i])

		if -1 == tmpIndex {
			return ""
		}

		tIndex := append(tIndex[t[i]], tmpIndex)

	}

	// 排列组合
	for _, v := range tIndex {
		for i := 0; i < len(v); i++ {

		}
	}

	return ""
}
