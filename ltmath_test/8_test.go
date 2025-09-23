package ltmath_test

import (
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func TestLT8MyAtoi(t *testing.T) {
	// str := "  +1337c0d3"
	// expected := 1337

	// str := "42"
	// expected := 42

	str := "   -042"
	expected := -42

	res := ltmath.LT8MyAtoi(str)

	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
