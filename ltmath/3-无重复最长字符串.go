package ltmath

func LT3_LengthOfLongestSubstring(s string) int {

	length := 0
	maxLen := 0
	left := 0
	right := 0
	sMap := map[byte]int{}

	for right < len(s) {

		_, ok := sMap[s[right]]

		if !ok {
			sMap[s[right]] = 1
			length++
			if length > maxLen {
				maxLen = length
			}
			right++

		} else {

			for _, ok := sMap[s[right]]; ok; _, ok = sMap[s[right]] {
				delete(sMap, s[left])
				left++
				length--
			}

			sMap[s[right]] = 1
			length++
			right++
		}
	}

	return maxLen
}
