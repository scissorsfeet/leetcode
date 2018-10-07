package LinkedList
import "testing"

//go test -v DeleteDuplicatesII.go DeleteDuplicatesII_test.go LinkedList.go --test.run=TestDeleteDuplicatesII1
func TestDeleteDuplicatesII1(t *testing.T) {
    arr1 := []int{1,3,3,3,5,7,7,9}
    linkedList1 := CreateListByArray(arr1)

    l := deleteDuplicatesII(linkedList1.Head.Next)
    tmpList := NewLinkedList(&ListNode{Next:l})
    tmpList.Print()
}

//go test -v DeleteDuplicatesII.go DeleteDuplicatesII_test.go LinkedList.go --test.run=TestDeleteDuplicatesII2
func TestDeleteDuplicatesII2(t *testing.T) {
    arr1 := []int{1,2,3,3,4,4,6}
    linkedList1 := CreateListByArray(arr1)

    l := deleteDuplicatesII(linkedList1.Head.Next)
    tmpList := NewLinkedList(&ListNode{Next:l})
    tmpList.Print()
}

//go test -v DeleteDuplicatesII.go DeleteDuplicatesII_test.go LinkedList.go --test.run=TestDeleteDuplicatesII3
func TestDeleteDuplicatesII3(t *testing.T) {
    arr1 := []int{1}
    linkedList1 := CreateListByArray(arr1)

    l := deleteDuplicatesII(linkedList1.Head.Next)
    tmpList := NewLinkedList(&ListNode{Next:l})
    tmpList.Print()
}