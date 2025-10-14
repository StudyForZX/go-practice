package ltmath_test

import (
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func TestLT55CanJump(t *testing.T) {
	nums := []int{0}
	expected := true
	// nums := []int{3, 2, 1, 0, 4}
	// expected := false
	// nums := []int{2, 3, 1, 1, 4}
	// expected := true

	res := ltmath.LT55CanJump(nums)

	if expected != res {
		t.Errorf("expected %t, got %t", expected, res)
	}
}
