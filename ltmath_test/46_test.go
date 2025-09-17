package ltmath_test

import (
	"reflect"
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func Test_LT46_Permute(t *testing.T) {
	nums := []int{1, 2, 3}
	expected := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}

	res := ltmath.LT46Permute(nums)

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
