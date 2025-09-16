package ltmath

import "sort"

func LT49GroupAnagramsByMyself(strs []string) [][]string {

	res := [][]string{}
	resMap := map[string][]string{}

	for _, str := range strs {
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool { return int(runes[i]) < int(runes[j]) })
		resMap[string(runes)] = append(resMap[string(runes)], str)
	}

	for _, v := range resMap {
		res = append(res, v)
	}

	return res

}
