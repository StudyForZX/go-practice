package main

import (
	"fmt"

	"github.com/studyforzx/lt/datastructure"
	"github.com/studyforzx/lt/ltmath"
)

func main() {

	listNode := &datastructure.ListNode[int]{}
	listNode = listNode.ArrayToListNode([]int{1, 2, 3, 4, 5})
	res := ltmath.LT206_ReverseListByFor(listNode).ListNodeToArray()
	fmt.Println(res)
}
