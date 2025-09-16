package ltmath_test

import (
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func TestLT11MaxAreaByMyself(t *testing.T) {

	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	expected := 49

	res := ltmath.LT11MaxAreaByMyself(nums)

	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}

}
