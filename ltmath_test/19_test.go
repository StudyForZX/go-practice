package ltmath_test

import (
	"reflect"
	"testing"

	"github.com/studyforzx/lt/datastructure"
	"github.com/studyforzx/lt/ltmath"
)

func Test_LT19_RemoveNthFromEndByFor(t *testing.T) {

	expected := []int{1, 2, 3, 5}
	var list *datastructure.ListNode[int]

	list = list.ArrayToListNode([]int{1, 2, 3, 4, 5})

	res := ltmath.LT19_RemoveNthFromEndByFor(list, 2)
	resArr := res.ListNodeToArray()

	if !reflect.DeepEqual(resArr, expected) {
		t.Errorf("expected %v, got %v", expected, resArr)
	}

}

func Test_LT19_RemoveNthFromEndByStack(t *testing.T) {

	expected := []int{1}
	var list *datastructure.ListNode[int]

	list = list.ArrayToListNode([]int{1, 2})

	res := ltmath.LT19_RemoveNthFromEndByStack(list, 1)
	resArr := res.ListNodeToArray()

	if !reflect.DeepEqual(resArr, expected) {
		t.Errorf("expected %v, got %v", expected, resArr)
	}

}
