package ltmath_test

import (
	"reflect"
	"testing"

	"github.com/studyforzx/lt/datastructure"
	"github.com/studyforzx/lt/ltmath"
)

func Test_LT141_HasCycleByStep(t *testing.T) {

	nums := []int{3, 2, 0, -4}
	expected := true

	l := &datastructure.ListNode[int]{}
	l.ArrayToListNode(nums)
	res := ltmath.LT141_HasCycleByStep(l)

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %t, got %t", expected, res)
	}
}
