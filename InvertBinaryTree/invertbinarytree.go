package InvertBinaryTree

import (
	binaryTree "github.com/liutianyi1989/leetcode/BinaryTree"
)

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
