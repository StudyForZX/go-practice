package ltmath_test

import (
	"reflect"
	"testing"

	"github.com/studyforzx/lt/datastructure"
	"github.com/studyforzx/lt/ltmath"
)

func Test_LT21_MergeTwoLists(t *testing.T) {

	expected := []int{1, 1, 2, 3, 4, 4}
	var list1, list2 *datastructure.ListNode[int]

	list1 = list1.ArrayToListNode([]int{1, 2, 4})
	list2 = list2.ArrayToListNode([]int{1, 3, 4})

	res := ltmath.LT21_MergeTwoLists(list1, list2)
	resArr := res.ListNodeToArray()

	if !reflect.DeepEqual(resArr, expected) {
		t.Errorf("expected %v, got %v", expected, resArr)
	}

}
