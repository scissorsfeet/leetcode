package MaximumDepthOfBinaryTree

import (
	"testing"

	binaryTree "github.com/liutianyi1989/leetcode/BinaryTree"
)

func TestA(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5, -1, -1, 7}
	root := binaryTree.CreateByArray(ints)

	max := maxDepth(root)
	if max != 4 {
		t.Error("error")
	} else {
		t.Log(max)
	}
}

func TestB(t *testing.T) {
	ints := []int{}
	root := binaryTree.CreateByArray(ints)

	max := maxDepth(root)
	if max != 0 {
		t.Error("error")
	} else {
		t.Log(max)
	}
}

func TestC(t *testing.T) {
	ints := []int{3}
	root := binaryTree.CreateByArray(ints)

	max := maxDepth(root)
	if max != 1 {
		t.Error("error")
	} else {
		t.Log(max)
	}
}
