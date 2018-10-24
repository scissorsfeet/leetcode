package BinarySearch

import "testing"

func TestIsPerfectSquare(t *testing.T) {
	if !isPerfectSquare(16) {
		t.Fatal(1)
	}

	if isPerfectSquare(14) {
		t.Fatal(2)
	}

	if !isPerfectSquare(1) {
		t.Fatal(3)
	}
}
