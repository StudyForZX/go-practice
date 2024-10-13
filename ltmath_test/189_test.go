package ltmath_test

import (
	"reflect"
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func Test_LT189_RotateByExchange(t *testing.T) {

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	expected := []int{5, 6, 7, 1, 2, 3, 4}

	res := ltmath.LT189_RotateByExchange(nums, k)

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("excepted %v,got %v", expected, res)
	}
}

func Test_LT189_RotateBySlice(t *testing.T) {

	nums := []int{1, 2}
	k := 3

	expected := []int{2, 1}

	res := ltmath.LT189_RotateBySlice(nums, k)

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("excepted %v,got %v", expected, res)
	}
}
