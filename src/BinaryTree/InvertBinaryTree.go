package BinaryTree

func invertTree(root *TreeNode) *TreeNode {
	if nil == root {
		return nil
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}