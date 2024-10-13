package notes

import "fmt"

func GoArrayNotes() {

	//
	// go数组笔记:
	// https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-array/
	//
	// 特点：内存连续、定长、无法改变
	//
	//

	// initialize
	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3}

	fmt.Println(arr1, arr2)
}
