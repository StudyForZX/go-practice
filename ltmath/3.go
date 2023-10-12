package ltmath

func LengthOfLongestSubstring(s string) string {

	sLen := len(s)

	for i := 0; i < sLen; i++ {

	}

}

// func LengthOfLongestSubstring(s string) string {

// 	sLen := len(s)

// 	if 0 == sLen {
// 		return ""
// 	}

// 	sMap := make(map[byte][]int, sLen)

// 	for i := 0; i < sLen; i++ {

// 		sMap[s[i]] = append(sMap[s[i]], i)

// 	}

// 	var maxLen int
// 	var maxStr string

// 	for i := 0; i < len(sMap); i++ {

// 		var max int

// 		item := sMap[s[i]]

// 		if 1 == len(item) {
// 			max = sLen - item[0]
// 		} else {
// 			max = item[len(item)-1] - item[0]
// 		}

// 		if maxLen < max {
// 			maxLen = max
// 			maxStr = s[i:maxLen]
// 		}
// 	}

// 	return maxStr
// }
