package LinkedList

import(
    "testing"
)

//go test -v LinkedList.go InsertionSortList.go InsertionSortList_test.go --test.run=TestSortList1
func TestInsertionSortList1(t *testing.T) {
    arr1 := []int{2,3,1,5,4}
    linkedList1 := CreateListByArray(arr1)
    linkedList1.Head.Next = sortList(linkedList1.Head.Next)
    linkedList1.Print()
}

//go test -v LinkedList.go SortList.go SortList_test.go --test.run=TestSortList2
func TestInsertionSortList2(t *testing.T) {
    arr1 := []int{2,1}
    linkedList1 := CreateListByArray(arr1)
    linkedList1.Head.Next = sortList(linkedList1.Head.Next)
    linkedList1.Print()
}

//go test -v LinkedList.go SortList.go SortList_test.go --test.run=TestSortList3
func TestInsertionSortList3(t *testing.T) {
    arr1 := []int{6,5,4,3,2,1}
    linkedList1 := CreateListByArray(arr1)
    linkedList1.Head.Next = sortList(linkedList1.Head.Next)
    linkedList1.Print()
}

//go test -v LinkedList.go SortList.go SortList_test.go --test.run=TestSortList4
func TestInsertionSortList4(t *testing.T) {
    arr1 := []int{1,2,3,4,5,6}
    linkedList1 := CreateListByArray(arr1)
    linkedList1.Head.Next = sortList(linkedList1.Head.Next)
    linkedList1.Print()
}