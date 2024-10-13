package ltmath

import (
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func Test_LT123_MaxProfit(t *testing.T) {

	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	expected := 6

	res := ltmath.LT123_MaxProfit(prices)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
