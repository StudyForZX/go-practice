package ltmath_test

import (
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func Test_LT1_TwoSum(t *testing.T) {

	target := 9
	nums := []int{2, 7, 11, 15}
	expected := []int{0, 1}

	res := ltmath.LT1_TwoSum(nums, target)

	if len(res) != 2 || res[0] != expected[0] || res[1] != expected[1] {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
