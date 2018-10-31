package Recursive

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	if nil == root {
		return 0
	}
	var res int
	helper(root, &res)
	return res
}

func helper(root *TreeNode, pRes *int) int {
	if nil == root {
		return 0
	}
	left := helper(root.Left, pRes) + 1
	right := helper(root.Right, pRes) + 1
	if nil == root.Left || root.Val != root.Left.Val {
		left = 0
	}
	if nil == root.Right || root.Val != root.Right.Val {
		right = 0
	}
	if *pRes < (left + right) {
		*pRes = left + right
	}
	ret := left
	if left < right {
		ret = right
	}
	return ret
}