package InvertBinaryTree

import (
	binaryTree "github.com/liutianyi1989/leetcode/BinaryTree"
)

/**
 * Definition for a binary tree node.
 * 二叉树定义
 */
type TreeNode struct {
	binaryTree.TreeNode
}

func invertTree(root *binaryTree.TreeNode) *binaryTree.TreeNode {
	if root == nil {
		return nil
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
