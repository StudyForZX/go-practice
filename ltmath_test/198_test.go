package ltmath_test

import (
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func Test_LT198_Rob(t *testing.T) {

	nums := []int{1, 2, 3, 1}
	expected := 4

	res := ltmath.LT198_Rob(nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
