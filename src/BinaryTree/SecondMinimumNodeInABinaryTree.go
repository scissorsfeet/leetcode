package BinaryTree

import "math"

func findSecondMinimumValue(root *TreeNode) int {
	if nil == root {
		return -1
	}
	var first, second = math.MaxInt32, math.MaxInt32
	dfs(root, &first, &second)
	if first == second {
		return -1
	}
	return second
}

func dfs(root *TreeNode, pFirst, pSecond *int) {
	if nil == root {
		return
	}
	if root.Val != *pFirst || root.Val != *pSecond {
		if root.Val < *pFirst {
			*pSecond = *pFirst
			*pFirst = root.Val
		} else if root.Val < *pSecond {
			*pSecond = root.Val
		}
	}
	dfs(root.Left, pFirst, pSecond)
	dfs(root.Right, pFirst, pSecond)
}