package reverse

import (
	"testing"
)

func TestNotEmpty(t *testing.T) {
	ret1 := reverseString("abcdef")
	if ret1 != "fedcba" {
		t.Error("reverse abcdef fail")
	} else {
		t.Log(ret1)
	}

	ret2 := reverseString("a")
	if ret2 != "a" {
		t.Error("reverse a fail")
	} else {
		t.Log(ret2)
	}
}

func TestEmpty(t *testing.T) {
	ret := reverseString("")
	if ret != "" {
		t.Error("reverse empty fail")
	} else {
		t.Log(ret)
	}
}

func TestNotEmptyV2(t *testing.T) {
	ret1 := reverseStringV2("abcdef")
	if ret1 != "fedcba" {
		t.Error("reverse abcdef fail")
	} else {
		t.Log(ret1)
	}

	ret2 := reverseStringV2("a")
	if ret2 != "a" {
		t.Error("reverse a fail")
	} else {
		t.Log(ret2)
	}
}

func TestEmptyV2(t *testing.T) {
	ret := reverseStringV2("")
	if ret != "" {
		t.Error("reverse empty fail")
	} else {
		t.Log(ret)
	}
}
