package ltmath

func LT17_LetterCombinations(digits string) []string {

	res := []string{}
	var tmp string

	if len(digits) == 0 {
		return res
	}

	digitsMap := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	var dfs func(index int)
	dfs = func(index int) {

		if index == len(digits) {

			res = append(res, tmp)

			return
		}

		for _, item := range digitsMap[string(digits[index])] {
			tmp = tmp + item
			dfs(index + 1)
			tmp = tmp[:len(tmp)-1]
		}
	}

	dfs(0)

	return res
}
