package ltmath

import (
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func Test_LT121_MaxProfit(t *testing.T) {

	prices := []int{7, 1, 5, 3, 6, 4}

	expected := 5

	res := ltmath.LT121MaxProfit(prices)

	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}

}
