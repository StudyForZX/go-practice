package ltmath

import "sort"

// 暴力破解 + map
// func FindAnagrams(s string, p string) []int {

// 	res := []int{}
// 	index := 0
// 	sArray := []byte(s)
// 	pArray := []byte(p)

// 	for i := 0; i < len(s); i++ {

// 		pHash := map[byte]int{}

// 		for _, v := range pArray {
// 			pHash[v] = 0
// 		}

// 		value, ok := pHash[sArray[i]]

// 		if ok {

// 			for index < len(s) && value == 0 {
// 				pHash[sArray[i]]++

// 				sum := 0
// 				for _, v := range pHash {
// 					sum += v
// 				}

// 				if sum == len(p) {
// 					res = append(res, index)
// 				}
// 			}

// 		}

// 		index++
// 	}

// 	return res
// }

// 排序对比
func LT438_FindAnagrams(s string, p string) []int {

	res := []int{}
	tmp := []byte{}

	pSlice := sortBytes([]byte(p))

	for i := 0; i <= len(s)-len(p); i++ {

		tmp = tmp[:0]
		tmp = append(tmp, s[i:i+len(p)]...)
		tmp = sortBytes([]byte(tmp))

		if string(tmp) == string(pSlice) {
			res = append(res, i)
		}
	}

	return res
}

func sortBytes(b []byte) []byte {

	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	return b
}
