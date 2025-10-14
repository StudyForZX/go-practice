package ltmath_test

import (
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func TestLT45Jump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	expected := 2

	res := ltmath.LT45Jump(nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
