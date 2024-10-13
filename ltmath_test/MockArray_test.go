package ltmath_test

import (
	"reflect"
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func TestMockArrayForeach(t *testing.T) {

	nums := []int{1, 2, 3}
	expected := [][3]int{
		{0, 0, 0},
		{0, 0, 1},
		{0, 0, 2},
		{0, 1, 0},
		{0, 1, 1},
		{0, 1, 2},
	}

	res := ltmath.MockArrayForeach(nums)

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
