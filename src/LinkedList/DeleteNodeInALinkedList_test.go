package LinkedList

import(
    "testing"
)

//go test -v LinkedList.go DeleteNodeInALinkedList.go DeleteNodeInALinkedList_test.go --test.run=TestDeleteNodeInALinkedList
func TestDeleteNodeInALinkedList(t *testing.T) {
    arr1 := []int{1,2,3,4,5}
    linkedList1 := CreateListByArray(arr1)
    DeleteNodeInALinkedList(linkedList1.Head.Next.Next.Next)
    linkedList1.Print()

    arr2 := []int{1}
    linkedList2 := CreateListByArray(arr2)
    DeleteNodeInALinkedList(linkedList2.Head.Next)
    linkedList2.Print()
}