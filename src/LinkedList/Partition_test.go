package LinkedList

import(
    "testing"
)

//go test -v LinkedList.go Partition.go Partition_test.go --test.run=TestPartition1
func TestPartition1(t *testing.T) {
    arr1 := []int{2,1,2,3,1,4,0,5}
    linkedList1 := CreateListByArray(arr1)
    h := partition(linkedList1.Head.Next, 3)
    tmpList := NewLinkedList(&ListNode{Next:h})
    tmpList.Print()
}

//go test -v LinkedList.go Partition.go Partition_test.go --test.run=TestPartition2
func TestPartition2(t *testing.T) {
    arr1 := []int{2,1}
    linkedList1 := CreateListByArray(arr1)
    h := partition(linkedList1.Head.Next, 2)
    tmpList := NewLinkedList(&ListNode{Next:h})
    tmpList.Print()
}

//go test -v LinkedList.go Partition.go Partition_test.go --test.run=TestPartition2
func TestPartition3(t *testing.T) {
    arr1 := []int{1}
    linkedList1 := CreateListByArray(arr1)
    h := partition(linkedList1.Head.Next, 0)
    tmpList := NewLinkedList(&ListNode{Next:h})
    tmpList.Print()
}