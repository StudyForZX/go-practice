package ltmath

func LT3LengthOfLongestSubstring(s string) int {

	length := 0
	maxLength := 0
	left := 0
	right := 0
	charMap := make(map[rune]bool, len(s))

	for _, c := range s {

		if !charMap[c] {
			charMap[c] = true
			right++
			length++
			maxLength = max(maxLength, length)
		} else {
			for charMap[c] {
				charMap[rune(s[left])] = false
				length--
				left++
			}
			charMap[c] = true
			right++
			length++
		}
	}

	return maxLength

}

func LT3LengthOfLongestSubstring2(s string) int {

	left := 0
	maxLen := 0
	charMap := make(map[rune]int)

	for i, char := range s {

		if lastIndex, exists := charMap[char]; exists && lastIndex >= left {
			left = lastIndex + 1
		}

		charMap[char] = i

		maxLen = max(maxLen, i-left+1)
	}

	return maxLen

}
