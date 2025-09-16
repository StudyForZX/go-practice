package ltmath

import (
	"reflect"
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func TestLT283MoveZeroesByMyself(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	expected := []int{1, 3, 12, 0, 0}

	res := ltmath.LT283MoveZeroesByMyself(nums)

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestLT283MoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	expected := []int{1, 3, 12, 0, 0}

	res := ltmath.LT283MoveZeroes(nums)

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
