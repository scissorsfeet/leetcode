package BinaryTree

func findBottomLeftValue(root *TreeNode) int {
	var res int
	var resultDepth = 0
	findBottomLeftValueHelper(root, &res, &resultDepth, 1)
	return res
}

func findBottomLeftValueHelper(root *TreeNode, res, resultDepth *int, curDepth int) {
	if nil == root {
		return
	}
	findBottomLeftValueHelper(root.Left, res, resultDepth, curDepth+1)
	if curDepth > *resultDepth {
		*resultDepth = curDepth
		*res = root.Val
	}
	findBottomLeftValueHelper(root.Right, res, resultDepth, curDepth+1)
}
