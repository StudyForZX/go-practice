package ltmath

func LT20_IsValid(s string) bool {

	if len(s) == 1 {
		return false
	}

	stack := []byte{}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {

		if pairs[s[i]] > 0 {

			if len(stack) == 0 || pairs[s[i]] != stack[len(stack)-1] {
				return false
			}

			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
