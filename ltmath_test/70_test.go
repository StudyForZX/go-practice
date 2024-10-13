package ltmath_test

import (
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func Test_LT70_ClimbStairsByFor(t *testing.T) {

	n := 11
	expected := 89

	res := ltmath.LT70_ClimbStairsByFor(n)

	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}

}
