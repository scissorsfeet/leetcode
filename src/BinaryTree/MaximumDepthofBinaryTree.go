package BinaryTree

func maxDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	var max int
	if left > right {
		max = left
	} else {
		max = right
	}
	return max + 1
}
