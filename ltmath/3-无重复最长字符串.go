package ltmath

func LT3_LengthOfLongestSubstring(s string) int {

	left := 0
	right := 0
	maxLen := 0
	length := 0
	strMap := map[byte]bool{}

	sArr := []byte(s)

	for _, c := range sArr {

		if !strMap[c] {

			strMap[c] = true
			right++
			length++
			maxLen = max(maxLen, length)

		} else {

			for strMap[c] {
				delete(strMap, sArr[left])
				left++
				length--
			}

			strMap[c] = true
			right++
			length++
		}
	}

	return maxLen
}
