package ltmath_test

import (
	"reflect"
	"testing"

	"github.com/studyforzx/lt/ltmath"
)

func TestLT49GroupAnagramsByMyself(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	expected := [][]string{
		{"bat"},
		{"nat", "tan"},
		{"ate", "eat", "tea"},
	}

	res := ltmath.LT49GroupAnagramsByMyself(strs)

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
