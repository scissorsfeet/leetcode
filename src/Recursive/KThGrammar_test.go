package Recursive

import "testing"

func TestKThGrammar1(t *testing.T) {
	if v := kthGrammar(1,1); v != 0 {
		t.Fail()
	}
	if v := kthGrammar(2,2); v != 1 {
		t.Fail()
	}
	if v := kthGrammar(3,1); v != 0 {
		t.Fail()
	}
	if v := kthGrammar(4,5); v != 1 {
		t.Fail()
	}
}
