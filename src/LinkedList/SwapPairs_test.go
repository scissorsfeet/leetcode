package LinkedList

import "testing"

//go test -v LinkedList.go SwapPairs.go SwapPairs_test.go --test.run=TestSwapPairs1
func TestSwapPairs1(t *testing.T) {
	arr1 := []int{1,2,3,4,5}
	linkedList1 := CreateListByArray(arr1)
	linkedList1.Head.Next = swapPairs(linkedList1.Head.Next)
	linkedList1.Print()
}

//go test -v LinkedList.go SwapPairs.go SwapPairs_test.go --test.run=TestSwapPairs2
func TestSwapPairs2(t *testing.T) {
	arr1 := []int{1,2,3}
	linkedList1 := CreateListByArray(arr1)
	linkedList1.Head.Next = swapPairs(linkedList1.Head.Next)
	linkedList1.Print()
}

//go test -v LinkedList.go SwapPairs.go SwapPairs_test.go --test.run=TestSwapPairs3
func TestSwapPairs3(t *testing.T) {
	arr1 := []int{1,2}
	linkedList1 := CreateListByArray(arr1)
	linkedList1.Head.Next = swapPairs(linkedList1.Head.Next)
	linkedList1.Print()
}

//go test -v LinkedList.go SwapPairs.go SwapPairs_test.go --test.run=TestSwapPairs4
func TestSwapPairs4(t *testing.T) {
	arr1 := []int{1}
	linkedList1 := CreateListByArray(arr1)
	linkedList1.Head.Next = swapPairs(linkedList1.Head.Next)
	linkedList1.Print()
}