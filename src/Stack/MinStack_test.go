package Stack

import "testing"

func TestMinStack_GetMin(t *testing.T) {
	s := Constructor()
	s.Push(5)
	s.Push(4)
	s.Push(1)
	s.Push(3)
	t.Log(s.GetMin())
	s.Pop()
	t.Log(s.GetMin())
	s.Pop()
	t.Log(s.GetMin())
}
