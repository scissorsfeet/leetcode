package BinaryTree

func inorderTraversal(root *TreeNode) []int {
	if nil == root {
		return nil
	}

	res := make([]int, 0, 1024)
	left := inorderTraversal(root.Left)
	if nil != left {
		res = left
	}
	res = append(res, root.Val)
	right := inorderTraversal(root.Right)
	if nil != right {
		res = append(res, right...)
	}
	return res
}

func inorderTraversalIter(root *TreeNode) []int {
	res := make([]int, 0, 1024)
	p := root
	stack := NewArrayStack()

	for !stack.IsEmpty() || nil != p {
		if nil != p {
			stack.Push(p)
			p = p.Left
		} else {
			p = stack.Pop().(*TreeNode)
			res = append(res, p.Val)
			p = p.Right
		}
	}

	return res
}