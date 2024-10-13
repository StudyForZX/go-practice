package ltmath_test

import (
	"reflect"
	"testing"

	"github.com/studyforzx/lt/datastructure"
	"github.com/studyforzx/lt/ltmath"
)

func Test_LT102_LevelOrder(t *testing.T) {

	expected := [][]int{
		{3},
		{9, 20},
		{15, 7},
	}

	treeNodeArr := []int{3, 9, 20, -1, -1, 15, 7}
	var tn *datastructure.TreeNode[int]
	tn = tn.ArrayToTreeNode(treeNodeArr, -1)

	res := ltmath.LT102_LevelOrder(tn)

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
