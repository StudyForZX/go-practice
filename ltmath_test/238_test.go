package ltmath_test

import (
	"reflect"
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func Test_LT238_ProductExceptSelfByFor(t *testing.T) {

	nums := []int{1, 2, 3, 4}
	expected := []int{24, 12, 8, 6}

	res := ltmath.LT238_ProductExceptSelfByFor(nums)

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
