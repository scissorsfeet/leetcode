package BinaryTree

func pruneTree(root *TreeNode) *TreeNode {
	if nil == root {
		return nil
	}
	isDelete := pruneTreeHelper(root)
	if isDelete {
		return nil
	}
	return root
}

func pruneTreeHelper(root *TreeNode) bool {
	if nil == root {
		return false
	}

	leftDelete := pruneTreeHelper(root.Left)
	if leftDelete {
		root.Left = nil
	}
	rightDelete := pruneTreeHelper(root.Right)
	if rightDelete {
		root.Right = nil
	}
	return root.Val == 0 && leftDelete && rightDelete
}
