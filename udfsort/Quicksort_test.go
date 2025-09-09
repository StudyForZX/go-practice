package udfsort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {

	expected := []int{1, 2, 3, 6, 8, 14, 23, 31, 111}

	res := QuickSort([]int{1, 3, 14, 2, 6, 8, 111, 23, 31})

	if reflect.DeepEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
