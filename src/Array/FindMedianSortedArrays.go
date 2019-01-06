package Array

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1Len := len(nums1)
	nums2Len := len(nums2)
	numsLen := nums1Len + nums2Len
	if 0 == numsLen {
		return float64(0)
	}
	if numsLen & 1 == 0 {
		return float64(findKth(nums1, nums2, numsLen/2) + findKth(nums1, nums2, (numsLen/2)+1))/2
	}
	return float64(findKth(nums1, nums2, (numsLen/2)+1))
}

func findKth(nums1, nums2 []int, k int) int {
	nums1Len := len(nums1)
	nums2Len := len(nums2)
	if nums1Len == 0 {
		return nums2[k-1]
	} else if nums2Len == 0 {
		return nums1[k-1]
	}
	if k == 1 {
		if nums1[0] > nums2[0] {
			return nums2[0]
		} else {
			return nums1[0]
		}
	}
	if nums1Len > nums2Len {
		return findKth(nums2, nums1, k)
	}

	var x = nums1Len
	if x > k/2 {
		x = k/2
	}
	y := k-x
	fmt.Println(x, y, k)
	if nums1[x-1] <= nums2[y-1] {
		return findKth(nums1[x:], nums2, k-x)
	}
	return findKth(nums1, nums2[y:], k-y)
}
