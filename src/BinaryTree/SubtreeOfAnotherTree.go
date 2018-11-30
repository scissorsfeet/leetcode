package BinaryTree

func isSubtree(s *TreeNode, t *TreeNode) bool {
	sWindow := tree2window(s)
	tWindow := tree2window(t)
	if len(sWindow) < len(tWindow) {
		return false
	}
	var start, end = 0, len(tWindow)-1
	var isSubtree = false
	for end < len(tWindow) {
		for i := start; i <= end; i++ {
			if (sWindow[i] != nil && tWindow[i] == nil) || (sWindow[i] == nil && tWindow[i] != nil) || sWindow[i].Val != tWindow[i].Val {
				break
			}
		}
		start += 1
		end += 1
	}
	return isSubtree
}

func tree2window(t *TreeNode) []*TreeNode {
	s := NewArrayStack()
	res := make([]*TreeNode, 0, 1024)
	if !s.IsEmpty() || nil != t {
		if nil != t {
			res = append(res, t)
			s.Push(t)
			t = t.Left
		} else {
			res = append(res, nil)
			top := s.Pop().(*TreeNode)
			t = top.Right
		}
	}
	res = append(res, nil)

	return res
}
