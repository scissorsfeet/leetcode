package Stack

import "testing"

func TestScoreOfParentheses(t *testing.T) {
	var s string

	s = "(()(()))"
	if scoreOfParentheses(s) != 6 {
		t.Fatal(1)
	}

	s = "(())"
	if scoreOfParentheses(s) != 2 {
		t.Fatal(2)
	}

	s = "()()"
	if scoreOfParentheses(s) != 2 {
		t.Fatal(3)
	}

	s = "()"
	if scoreOfParentheses(s) != 1 {
		t.Fatal(4)
	}

	s = "(())()"
	if scoreOfParentheses(s) != 3 {
		t.Fatal(5)
	}
}