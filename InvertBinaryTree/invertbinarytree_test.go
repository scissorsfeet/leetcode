package InvertBinaryTree

import (
	"testing"

	binaryTree "github.com/liutianyi1989/leetcode/BinaryTree"
)

func TestA(t *testing.T) {
	arr := []int{1, 2, 4, 5, 3, -1, -1, 3}
	root := binaryTree.CreateByArray(arr)
	root.LevelTraverse()
	invertTree(root)
	root.LevelTraverse()
}

func TestB(t *testing.T) {
	arr := []int{}
	root := binaryTree.CreateByArray(arr)
	root.LevelTraverse()
	invertTree(root)
	root.LevelTraverse()
}

func TestC(t *testing.T) {
	arr := []int{1, 2}
	root := binaryTree.CreateByArray(arr)
	root.LevelTraverse()
	invertTree(root)
	root.LevelTraverse()
}
