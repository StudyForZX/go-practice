package ltmath_test

import (
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func Test_LT3_LengthOfLongestSubstring(t *testing.T) {

	s := "abcabcbb"
	expected := 3

	res := ltmath.LT3_LengthOfLongestSubstring(s)

	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}

}
