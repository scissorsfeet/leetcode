package BinaryTree

func preorderTraversal(root *TreeNode) []int {
	res := preorderHelper(root)
	return res
}

func preorderHelper(root *TreeNode) []int {
	if nil == root {
		return nil
	}

	res := []int{root.Val}
	left := preorderHelper(root.Left)
	if nil != left {
		res = append(res, left...)
	}
	right := preorderHelper(root.Right)
	if nil != right {
		res = append(res, right...)
	}
	return res
}

func preorderTraversalIter(root *TreeNode) []int {
	stack := NewArrayStack()
	p := root
	res := make([]int, 0, 1024)

	for !stack.IsEmpty() || nil != p {
		if nil != p {
			stack.Push(p)
			res = append(res, p.Val)
			p = p.Left
		} else {
			p = stack.Pop().(*TreeNode)
			p = p.Right
		}
	}

	return res
}