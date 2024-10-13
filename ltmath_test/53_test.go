package ltmath_test

import (
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func Test_LT53_MaxSubArray(t *testing.T) {

	nums := []int{-2, -1}
	expected := -1

	res := ltmath.LT53_MaxSubArray(nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

}
