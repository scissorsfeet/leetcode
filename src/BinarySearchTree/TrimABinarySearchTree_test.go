package _4_tree

import "testing"

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

func TestTrimBST(t *testing.T) {
	bst := NewBST(2, compareFunc)

	bst.Insert(1)
	bst.Insert(3)

	t.Log(trimBST(bst.root, 3, 4))
	bst.InOrderTraverse()
}