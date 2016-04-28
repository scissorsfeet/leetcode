package adddigits

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var ret int

	ret = addDigits(3)
	if ret != 3 {
		t.Error("3 is err")
	} else {
		t.Log("pass")
	}

	ret = addDigits(10)
	if ret != 1 {
		t.Error("10 is err")
	} else {
		t.Log("pass")
	}

	ret = addDigits(38)
	if ret != 2 {
		t.Error("38 is err")
	} else {
		t.Log("pass")
	}

	ret = addDigits(128)
	if ret != 2 {
		t.Error("128 is err")
	} else {
		t.Log("pass")
	}
}
