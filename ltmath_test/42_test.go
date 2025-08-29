package ltmath_test

import (
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func Test_LT42_Trap(t *testing.T) {

	expected := 6
	heights := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// heights := []int{4, 2, 0, 3, 2, 5}

	res := ltmath.LT42_Trap_ByMySelf(heights)

	if res != expected {
		t.Errorf("excepted %d, got %d", expected, res)
	}
}
