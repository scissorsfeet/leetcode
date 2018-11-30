package BinaryTree

func increasingBST(root *TreeNode) *TreeNode {
	if nil == root {
		return nil
	}

	s := NewArrayStack()
	p := root
	var newRoot *TreeNode = nil
	var q *TreeNode = nil
	for !s.IsEmpty() || nil != p {
		if nil != p {
			s.Push(p)
			p = p.Left
		} else {
			t := s.Pop().(*TreeNode)
			if nil == newRoot {
				newRoot = t
				q = t
			} else {
				q.Right = t
			}
			p = t.Right
		}
	}

	return newRoot
}
