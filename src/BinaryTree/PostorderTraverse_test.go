package BinaryTree

import "testing"

func TestPostorderTraversal(t *testing.T) {
	root := &TreeNode{Val:1}
	root.Left = &TreeNode{Val:2}
	root.Right = &TreeNode{Val:3}

	root.Left.Left = &TreeNode{Val:4}

	t.Log(postorderTraversal(root))
}

func TestPostorderTraversalIter(t *testing.T) {
	root := &TreeNode{Val:1}
	root.Left = &TreeNode{Val:2}
	root.Right = &TreeNode{Val:3}

	root.Left.Left = &TreeNode{Val:4}

	t.Log(postorderTraversalIter(root))
}