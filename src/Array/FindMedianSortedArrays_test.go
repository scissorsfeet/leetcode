package Array

import "testing"

func TestFindMedianSortedArray(t *testing.T) {
	nums1 := []int{2,3,4}
	nums2 := []int{1}
	t.Log(findMedianSortedArrays(nums1, nums2))
}