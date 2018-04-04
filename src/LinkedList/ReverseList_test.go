package LinkedList

import(
    "testing"
)

//go test -v LinkedList.go ReverseList.go ReverseList_test.go --test.run=TestReverseList
func TestReverseList(t *testing.T) {
    arr1 := []int{1,2,3,4,5}
    linkedList1 := CreateListByArray(arr1)
    linkedList1.Head.Next = reverseList(linkedList1.Head.Next)
    linkedList1.Print()
}