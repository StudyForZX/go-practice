package practice

import "testing"

func Test(t *testing.T) {
	expected := 10
	res := PracticeAppendScliceAndRetrunIndex(5)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
