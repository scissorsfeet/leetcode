package LinkedList

import(
    "testing"
)

//go test -v LinkedList.go ReorderList.go ReorderList_test.go --test.run=TestReorderList1
func TestReorderList1(t *testing.T) {
    arr1 := []int{1,2,3,4,5}
    linkedList1 := CreateListByArray(arr1)
    reorderList(linkedList1.Head.Next)
    linkedList1.Print()
}

//go test -v LinkedList.go ReorderList.go ReorderList_test.go --test.run=TestReorderList2
func TestReorderList2(t *testing.T) {
    arr1 := []int{1,2}
    linkedList1 := CreateListByArray(arr1)
    reorderList(linkedList1.Head.Next)
    linkedList1.Print()
}

//go test -v LinkedList.go ReorderList.go ReorderList_test.go --test.run=TestReorderList3
func TestReorderList3(t *testing.T) {
    arr1 := []int{1,2,3,4,5,6}
    linkedList1 := CreateListByArray(arr1)
    reorderList(linkedList1.Head.Next)
    linkedList1.Print()
}