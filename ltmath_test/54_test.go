package ltmath_test

import (
	"reflect"
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func Test_LT54_SpiraOrder(t *testing.T) {

	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	expected := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}

	res := ltmath.LT54_SpiraOrder(matrix)

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
