package ltmath_test

import (
	"reflect"
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func Test_LT118_Generate(t *testing.T) {

	numsRow := 5
	expected := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}

	res := ltmath.LT118_GenerateByLeetCode(numsRow)

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}

}
