package ltmath_test

import (
	"reflect"
	"testing"

	"github.com/studyforzx/lt/datastructure"
	"github.com/studyforzx/lt/ltmath"
)

func TestLT25ReverseKGroupByMyself(t *testing.T) {
	listNodeIns := &datastructure.ListNode[int]{}
	ln := listNodeIns.ArrayToListNode([]int{1, 2, 3, 4, 5})

	expected := []int{2, 1, 4, 3, 5}

	res := ltmath.LT25ReverseKGroup(ln, 2).ListNodeToArray()

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
