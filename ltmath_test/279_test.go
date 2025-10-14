package ltmath_test

import (
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func TestLT279NumSquares(t *testing.T) {

	num := 82
	expected := 2
	res := ltmath.LT279NumSquares(num)
	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
