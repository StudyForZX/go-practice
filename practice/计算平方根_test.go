package practice

import (
	"testing"
)

func TestPracticeSquareRoot(t *testing.T) {
	expected := 10
	res := practiceSquareRoot(100)
	if expected != res {
		t.Errorf("expected %d, res is %d", expected, res)
	}
}
