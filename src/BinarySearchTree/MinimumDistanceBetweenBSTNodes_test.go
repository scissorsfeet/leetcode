package _4_tree

import (
	"testing"
)

var compareFunc = func(v, nodeV interface{}) int {
	vInt := v.(int)
	nodeVInt := nodeV.(int)

	if vInt > nodeVInt {
		return 1
	} else if vInt < nodeVInt {
		return -1
	}
	return 0
}

func TestMinDiffInBST(t *testing.T) {
	bst := NewBST(20, compareFunc)

	bst.Insert(10)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(5)

	t.Log(minDiffInBST(bst.root))
}

func TestMinDiffInBSTII(t *testing.T) {
	bst := NewBST(20, compareFunc)

	bst.Insert(10)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(5)

	t.Log(minDiffInBSTII(bst.root))
}