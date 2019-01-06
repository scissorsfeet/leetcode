package Array

import (
	"fmt"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	var nums1 []int

	nums1 = []int{4,5,6,0,0,0}
	merge(nums1,3,[]int{1,2,3},3)
	fmt.Println(nums1)

	nums1 = []int{0,0,0}
	merge(nums1,0,[]int{2,5,6},3)
	fmt.Println(nums1)

	nums1 = []int{1,0}
	merge(nums1,1,[]int{2},1)
	fmt.Println(nums1)
}
