package ltmath_test

import (
	"testing"

	"github.com/studyforzx/lt/datastructure"
	"github.com/studyforzx/lt/ltmath"
)

func TestLT234IsPalindromeByMyself(t *testing.T) {

	list := &datastructure.ListNode[int]{}
	list = list.ArrayToListNode([]int{1, 2, 2, 1})

	expected := true

	res := ltmath.LT234IsPalindromeByMyself(list)

	if expected != res {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
