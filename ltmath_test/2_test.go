package ltmath_test

import (
	"reflect"
	"testing"

	"github.com/studyforzx/lt/datastructure"
	"github.com/studyforzx/lt/ltmath"
)

func TestLT2AddTwoNumbersByMyself(t *testing.T) {

	listIns := &datastructure.ListNode[int]{}
	l1 := listIns.ArrayToListNode([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := listIns.ArrayToListNode([]int{9, 9, 9, 9})

	expected := []int{8, 9, 9, 9, 0, 0, 0, 1}

	res := ltmath.LT2AddTwoNumbersByMyself(l1, l2).ListNodeToArray()

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
