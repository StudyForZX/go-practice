package ltmath_test

import (
	"reflect"
	"testing"

	"github.com/studyforzx/lt/datastructure"
	"github.com/studyforzx/lt/ltmath"
)

func Test_LT148_SortList(t *testing.T) {

	nums := []int{4, 2, 1, 3}
	expected := []int{1, 2, 3, 4}

	l := &datastructure.ListNode[int]{}
	l = l.ArrayToListNode(nums)
	res := ltmath.LT148_SortListByTransToArray(l)
	resArr := res.ListNodeToArray()

	if !reflect.DeepEqual(resArr, expected) {
		t.Errorf("expected %v, got %v", expected, resArr)
	}
}
