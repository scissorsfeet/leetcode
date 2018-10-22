package Stack

import "testing"

func TestBackspaceCompare(t *testing.T) {
	s := "ab#c"
	tt := "ad#c"
	if !backspaceCompare(s, tt) {
		t.Fail()
	}

	s = "ab##"
	tt = "c#d#"
	if !backspaceCompare(s, tt) {
		t.Fail()
	}

	s = "a#c"
	tt = "b"
	if backspaceCompare(s, tt) {
		t.Fail()
	}
}
