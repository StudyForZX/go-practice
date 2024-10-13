package ltmath_test

import (
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func Test_LT560_SubarraySum(t *testing.T) {

	nums := []int{3, 4, 7, 2, -3, 1, 4, 2}
	k := 7
	expected := 4

	res := ltmath.LT560_SubarraySumByPreSum(nums, k)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

}
