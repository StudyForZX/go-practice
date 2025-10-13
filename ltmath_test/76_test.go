package ltmath

import (
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func TestXxx(t *testing.T) {

	sStr := "ADOBECODEBANC"
	tStr := "ABC"

	expected := "BANC"
	res := ltmath.LT76MinWindow(sStr, tStr)

	if expected != res {
		t.Errorf("expected %s, res %s", expected, res)
	}
}
