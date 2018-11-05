package Sort

import "testing"

func TestIsAnagram(t *testing.T) {
	if !isAnagram("cat", "tca") {
		t.Fail()
	}
	if !isAnagram("abcde", "ecdab") {
		t.Fail()
	}
	if isAnagram("aaaaaa", "bbbbb") {
		t.Fail()
	}
}
