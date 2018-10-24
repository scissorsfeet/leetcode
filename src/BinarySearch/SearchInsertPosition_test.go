package BinarySearch

import "testing"

func TestSearchInsert(t *testing.T) {
	t.Log(searchInsert([]int{1,3,5,7}, 4))
	t.Log(searchInsert([]int{1,3,5,7}, 8))
	t.Log(searchInsert([]int{1,3,5,7}, -1))
}
