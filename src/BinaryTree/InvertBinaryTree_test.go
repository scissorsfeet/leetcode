package BinaryTree

import "testing"

func TestInvertTree(t *testing.T) {
	root := &TreeNode{Val:1}
	root.Left = &TreeNode{Val:2}
	root.Left.Left = &TreeNode{Val:3}

	root.InOrderTraverse()
}
