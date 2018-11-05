package BinarySearch

import "testing"

//go test -v Intersection.go Intersection_test.go --test.run=TestIntersection1
func TestIntersection1(t *testing.T) {
	arr1 := []int{1}
	arr2 := []int{1,2}
	i := intersection(arr1, arr2)
	t.Log(i)
}
