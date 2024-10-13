package ltmath_test

import (
	"reflect"
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func Test_LT17_LetterCombinations(t *testing.T) {
	digits := "23"
	expected := []string{
		"ad",
		"ae",
		"af",
		"bd",
		"be",
		"bf",
		"cd",
		"ce",
		"cf",
	}

	res := ltmath.LT17_LetterCombinations(digits)

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
