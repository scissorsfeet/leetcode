package SameTree

import (
	binaryTree "github.com/liutianyi1989/leetcode/BinaryTree"
)

func isSameTree(p *binaryTree.TreeNode, q *binaryTree.TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if (p == nil && q != nil) || (p != nil && q == nil) || (p.Val != q.Val) {
		return false
	}
	isLeftSame := isSameTree(p.Left, q.Left)
	if !isLeftSame {
		return false
	}
	isRightSame := isSameTree(p.Right, q.Right)
	if !isRightSame {
		return false
	}
	return true
}
