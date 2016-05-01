package BinaryTree

import (
	"testing"
)

func TestCreate(t *testing.T) {
	arr := []int{1, 4, 3, -1, 2, -3, 4}
	root := CreateByArray(arr)
	root.LevelTraverse()
}
