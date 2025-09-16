package ltmath_test

import (
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func Test_LT35_SearchInsertByFor(t *testing.T) {

	nums := []int{1, 3, 5, 6}
	target := 5
	expected := 2

	res := ltmath.LT35SearchInsertByMyself(nums, target)

	if expected != res {
		t.Errorf("excepted %d, got %d", expected, res)
	}
}
