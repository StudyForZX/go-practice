package ltmath

import (
	"reflect"
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func Test_LT78_SubsetsByFor(t *testing.T) {

	nums := []int{1, 2, 3}
	expected := [][]int{
		{},
		{1},
		{2},
		{1, 2},
		{3},
		{1, 3},
		{2, 3},
		{1, 2, 3},
	}

	res := ltmath.LT78_SubsetsByFor(nums)

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}

}

func Test_LT78_SubsetsByRecursion(t *testing.T) {

	nums := []int{1, 2, 3}
	expected := [][]int{
		{},
		{1},
		{2},
		{1, 2},
		{3},
		{1, 3},
		{2, 3},
		{1, 2, 3},
	}

	res := ltmath.LT78_SubsetsByRecursion(nums)

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}

}
