package Stack

import "testing"

func TestBracket(t *testing.T) {
	s := "(){}"
	if !isValid(s) {
		t.Fail()
	}
	s = ""
	if !isValid(s) {
		t.Fail()
	}
	s = "[{}]"
	if !isValid(s) {
		t.Fail()
	}
	s = "[{]}"
	if isValid(s) {
		t.Fail()
	}
	s = "[{}])"
	if isValid(s) {
		t.Fail()
	}
	s = "[)"
	if isValid(s) {
		t.Fail()
	}
}
