package BinaryTree

func postorderTraversal(root *TreeNode) []int {
	if nil == root {
		return nil
	}
	left := postorderTraversal(root.Left)
	right := postorderTraversal(root.Right)
	res := []int{}
	if nil != left {
		res = left
	}
	if nil != right {
		res = append(res, right...)
	}
	res = append(res, root.Val)
	return res
}

func postorderTraversalIter(root *TreeNode) []int {
	if nil == root {
		return nil
	}
	res := make([]int, 0, 1024)
	p := root
	s1 := NewArrayStack()
	s2 := NewArrayStack()
	s1.Push(p)
	for !s1.IsEmpty() {
		t := s1.Pop().(*TreeNode)
		s2.Push(t)
		if nil != t.Left {
			s1.Push(t.Left)
		}
		if nil != t.Right {
			s1.Push(t.Right)
		}
	}

	for !s2.IsEmpty() {
		t := s2.Pop().(*TreeNode)
		res = append(res, t.Val)
	}

	return res
}