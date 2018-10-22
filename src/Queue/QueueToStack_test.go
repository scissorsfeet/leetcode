package Queue

import "testing"

func TestQueue2Stack(t *testing.T) {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	t.Log(q.Top())
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Pop())
}