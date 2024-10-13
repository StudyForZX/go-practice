package ltmath

import (
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func Test_LT122_MaxProfit(t *testing.T) {

	prices := []int{7, 1, 5, 3, 6, 4}

	expected := 7

	res := ltmath.LT122_MaxProfit(prices)

	if expected != res {
		t.Errorf("excepted %v, got %v", expected, res)
	}

}
