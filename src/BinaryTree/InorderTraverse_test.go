package BinaryTree

import "testing"

func TestInorderTraversal(t *testing.T) {
	root := &TreeNode{Val:1}
	root.Left = &TreeNode{Val:2}
	root.Right = &TreeNode{Val:3}

	root.Left.Left = &TreeNode{Val:4}

	t.Log(inorderTraversal(root))
}

func TestInorderTraversalIter(t *testing.T) {
	root := &TreeNode{Val:1}
	root.Left = &TreeNode{Val:2}
	root.Right = &TreeNode{Val:3}

	root.Left.Left = &TreeNode{Val:4}

	t.Log(inorderTraversalIter(root))
}