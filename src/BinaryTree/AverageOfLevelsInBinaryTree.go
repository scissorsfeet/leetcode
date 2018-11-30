package BinaryTree

func averageOfLevels(root *TreeNode) []float64 {
	q := NewMyCircularQueue(10240)
	q.EnQueue(root)
	res := make([]float64,0,32)

	for !q.IsEmpty() {
		levelSum := 0
		levelNum := q.size
		for i:=0;i<levelNum;i++ {
			cur := q.Front().(*TreeNode)
			q.DeQueue()
			levelSum += cur.Val
			if nil != cur.Left {
				q.EnQueue(cur.Left)
			}
			if nil != cur.Right {
				q.EnQueue(cur.Right)
			}
		}
		avg := float64(levelSum)/float64(levelNum)
		res = append(res, avg)
	}

	return res
}
