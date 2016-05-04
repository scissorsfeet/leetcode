/*
Given a binary tree, find its maximum depth.
找到所给出的二叉树的最大深度
The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
最大深度：从根节点到最远叶子节点所经过的距离
*/
package MaximumDepthOfBinaryTree

import (
	binaryTree "github.com/liutianyi1989/leetcode/BinaryTree"
)

func maxDepth(root *binaryTree.TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth((root.Right))
	if leftDepth >= rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}
