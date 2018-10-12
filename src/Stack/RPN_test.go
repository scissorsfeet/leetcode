package Stack

import "testing"

func TestRPN(t *testing.T) {
	s := []string{"2", "1", "+", "3", "*"}
	if evalRPN(s) != 9 {
		t.Fail()
	}
	s = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	if evalRPN(s) != 22 {
		t.Fail()
	}
	s = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	if evalRPN(s) != 22 {
		t.Fail()
	}
}
