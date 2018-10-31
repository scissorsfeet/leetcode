package BinarySearch

import "testing"

//go test -v PeakIndexInMountainArray.go PeakIndexInMountainArray_test.go --test.run=TestPeakIndexInMountainArray1
func TestPeakIndexInMountainArray1(t *testing.T) {
	arr1 := []int{1,2,3,4,3,2,1}
	i := peakIndexInMountainArray(arr1)
	t.Log(i)
}

//go test -v PeakIndexInMountainArray.go PeakIndexInMountainArray_test.go --test.run=TestPeakIndexInMountainArray2
func TestPeakIndexInMountainArray2(t *testing.T) {
	arr1 := []int{1,4,3,2,1}
	i := peakIndexInMountainArray(arr1)
	t.Log(i)
}

//go test -v PeakIndexInMountainArray.go PeakIndexInMountainArray_test.go --test.run=TestPeakIndexInMountainArray3
func TestPeakIndexInMountainArray3(t *testing.T) {
	arr1 := []int{1,2,3,4,1}
	i := peakIndexInMountainArray(arr1)
	t.Log(i)
}
