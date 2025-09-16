package ltmath

func LT3LengthOfLongestSubstring(s string) int {

	length := 0
	maxLength := 0
	left := 0
	right := 0
	charMap := make(map[rune]struct{}, len(s))

	for _, c := range s {

		if _, ok := charMap[c]; !ok {

			charMap[c] = struct{}{}
			right++
			length++
			maxLength = max(maxLength, length)

		} else {

			for {
				if _, ok := charMap[c]; ok {
					delete(charMap, rune(s[left]))
					length--
					left++
				} else {
					break
				}
			}

			charMap[c] = struct{}{}
			right++
			length++
		}

	}

	return maxLength

}
