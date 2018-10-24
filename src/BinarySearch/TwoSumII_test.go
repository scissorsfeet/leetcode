package BinarySearch

import "testing"

func TestTwoSum(t *testing.T) {
	t.Log(twoSum([]int{2, 7, 11, 15}, 9))
	t.Log(twoSum([]int{2}, 2))
	t.Log(twoSum([]int{2,7,11,15}, 22))
}