package BinaryTree

func levelOrder(root *TreeNode) [][]int {
	if nil == root {
		return nil
	}
	q := NewMyCircularQueue(1000000)
	q.EnQueue(root)
	res := make([][]int, 0, 1024)
	levelNum := 0
	for !q.IsEmpty() {
		n := q.size
		res = append(res, make([]int,0,32))
		for i:=0;i<n;i++ {
			p := q.Front().(*TreeNode)
			q.DeQueue()
			res[levelNum] = append(res[levelNum], p.Val)
			if nil != p.Left {
				q.EnQueue(p.Left)
			}
			if nil != p.Right {
				q.EnQueue(p.Right)
			}
		}
		levelNum++
	}
	return res
}
