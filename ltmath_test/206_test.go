package ltmath_test

import (
	"reflect"
	"testing"

	"github.com/studyforzx/lt/datastructure"
	"github.com/studyforzx/lt/ltmath"
)

func Test_LT206_ReversePart(t *testing.T) {

	nums := []int{1, 2, 3, 4, 5, 7}
	expected := []int{1, 2, 4, 3, 5, 7}

	m := 2
	n := 3

	ln := &datastructure.ListNode[int]{}
	ln = ln.ArrayToListNode(nums)

	res := ltmath.Test_LT206_ReversePart(ln, m, n)

	resArr := res.ListNodeToArray()

	if !reflect.DeepEqual(resArr, expected) {
		t.Errorf("expected %v, got %v", expected, resArr)
	}
}
