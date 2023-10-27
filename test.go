package main

import (
	"fmt"
)

// func permute(letters map[byte][]int) [][]int {
// 	var helper func(map[byte][]int, byte, []int)
// 	res := [][]int{}

// 	helper = func(letters map[byte][]int, letter byte, current []int) {
// 		if len(current) == len(letters) {
// 			tmp := make([]int, len(current))
// 			copy(tmp, current)
// 			res = append(res, tmp)
// 			return
// 		}

// 		for _, num := range letters[letter] {
// 			current = append(current, num)
// 			helper(letters, letter, current)
// 			current = current[:len(current)-1]
// 		}
// 	}

// 	helper(letters, 'a', []int{})
// 	return res
// }

func permute(letters map[byte][]int) [][]int {

	var res [][]int
	var helper func(map[byte][]int, []int)

	helper = func(m map[byte][]int, i []int) {

		for _, index := range letters {

		}

	}
}

func main() {
	letters := map[byte][]int{
		'a': []int{1, 2, 3},
		'b': []int{4, 5},
		'c': []int{6, 7},
	}
	res := permute(letters)
	fmt.Println(res)
}
