package Recursive

import "testing"

func TestLongestUnivaluePath1(t *testing.T) {
	root := &TreeNode{Val:3}
	root.Left = &TreeNode{Val:4}
	root.Right = &TreeNode{Val:3}

	root.Left.Left = &TreeNode{Val:4}
	root.Left.Right = &TreeNode{Val:4}

	root.Left.Left.Left = &TreeNode{Val:4}

	root.Right.Right = &TreeNode{Val:3}
	root.Right.Right.Left = &TreeNode{Val:3}
	root.Right.Right.Right = &TreeNode{Val:3}
	root.Right.Right.Right.Right = &TreeNode{Val:3}
	root.Right.Right.Right.Left = &TreeNode{Val:3}

	t.Log(longestUnivaluePath(root))
}

func TestLongestUnivaluePath2(t *testing.T) {
	root := &TreeNode{Val:2}
	root.Left = &TreeNode{Val:4}
	root.Right = &TreeNode{Val:5}

	root.Left.Left = &TreeNode{Val:1}
	root.Left.Right = &TreeNode{Val:1}

	root.Right.Left = &TreeNode{Val:5}
	root.Right.Right = &TreeNode{Val:5}

	t.Log(longestUnivaluePath(root))
}

func TestLongestUnivaluePath3(t *testing.T) {
	root := &TreeNode{Val:1}

	t.Log(longestUnivaluePath(root))
}

func TestLongestUnivaluePath4(t *testing.T) {
	root := &TreeNode{Val:1}
	root.Left = &TreeNode{Val:1}
	root.Right = &TreeNode{Val:1}

	root.Left.Left = &TreeNode{Val:1}
	root.Right.Right = &TreeNode{Val:1}
	root.Right.Right.Right = &TreeNode{Val:1}

	t.Log(longestUnivaluePath(root))
}