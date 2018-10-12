package Queue

import "testing"

func TestMyCircularQueue(t *testing.T) {
	q := Constructor(3)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)

	//t.Log(q.Front())
	//t.Log(q.Rear())
	//t.Log("--------------")

	q.DeQueue()
	//t.Log(q.Front())
	//t.Log(q.Rear())
	//t.Log("--------------")

	q.EnQueue(6)
	//t.Log(q.Front())
	//t.Log(q.Rear())
	//t.Log("--------------")

	q.DeQueue()
	t.Log(q.Front())
	t.Log(q.Rear())
}
