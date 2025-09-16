package ltmath_test

import (
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func TestLT128LongestConsecutive(t *testing.T) {

	nums := []int{100, 4, 200, 1, 3, 2}

	expected := 4

	res := ltmath.LT128LongestConsecutive(nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

}
