package BinarySearch

import "testing"

func TestBinarySearch(t *testing.T) {
	var nums []int

	nums = []int{2,5,6,7,10}

	if search(nums, 5) != 1 {
		t.Fatal(1)
	}

	if search(nums, 10) != 4 {
		t.Fatal(1)
	}

	if search(nums, 11) != -1 {
		t.Fatal(1)
	}
}