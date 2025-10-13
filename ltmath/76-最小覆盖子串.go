package ltmath

func LT76MinWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	window := make(map[byte]int)
	have, needCount := 0, len(need)
	res, resLen := []int{-1, -1}, len(s)+1

	left := 0
	for right := 0; right < len(s); right++ {
		sChar := s[right]
		window[sChar]++
		if need[sChar] > 0 && window[sChar] == need[sChar] {
			have++
		}

		for have == needCount {
			// 更新最小区间
			if right-left+1 < resLen {
				res = []int{left, right}
				resLen = right - left + 1
			}

			// 左边收缩
			window[s[left]]--
			if need[s[left]] > 0 && window[s[left]] < need[s[left]] {
				have--
			}
			left++
		}
	}

	if resLen == len(s)+1 {
		return ""
	}
	return s[res[0] : res[1]+1]
}
