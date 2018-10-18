package Stack

import "testing"

func TestStackToQueue(t *testing.T) {
	q := MyQueue{NewArrayStack(), NewArrayStack()}
	q.Push(1)
	q.Push(2)
	t.Log(q.Peek())
	t.Log(q.Pop())
	t.Log(q.Empty())
}