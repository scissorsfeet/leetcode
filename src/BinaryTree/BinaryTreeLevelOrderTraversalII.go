package BinaryTree

func levelOrderBottom(root *TreeNode) [][]int {
	if nil == root {
		return nil
	}
	s := NewArrayStack()
	q := NewMyCircularQueue(1024)

	q.EnQueue(root)

	for !q.IsEmpty() {
		size := q.size
		oneLevel := make([]int, 0, size)
		for i:=0;i<size;i++ {
			cur := q.Front().(*TreeNode)
			q.DeQueue()
			oneLevel = append(oneLevel, cur.Val)
			if nil != cur.Left {
				q.EnQueue(cur.Left)
			}
			if nil != cur.Right {
				q.EnQueue(cur.Right)
			}
		}
		s.Push(oneLevel)
	}

	res := make([][]int, 0, 1024)
	for !s.IsEmpty() {
		res = append(res, s.Pop().([]int))
	}

	return res
}