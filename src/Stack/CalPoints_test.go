package Stack

import "testing"

func TestCallPoints(t *testing.T) {
	ops := []string{"5","2","C","D","+"}
	t.Log(calPoints(ops))

	ops = []string{"5","-2","4","C","D","9","+","+"}
	t.Log(calPoints(ops))
}