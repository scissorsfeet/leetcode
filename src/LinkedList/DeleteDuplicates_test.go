package LinkedList
import "testing"

//go test -v DeleteDuplicates.go DeleteDuplicates_test.go LinkedList.go --test.run=TestDeleteDuplicates1
func TestDeleteDuplicates1(t *testing.T) {
    arr1 := []int{1,3,3,3,5,7,7,9}
    linkedList1 := CreateListByArray(arr1)

    l := deleteDuplicates(linkedList1.Head.Next)
    tmpList := NewLinkedList(&ListNode{Next:l})
    tmpList.Print()
}

//go test -v DeleteDuplicates.go DeleteDuplicates_test.go LinkedList.go --test.run=TestDeleteDuplicates2
func TestDeleteDuplicates2(t *testing.T) {
    arr1 := []int{1,3,3,3,5,7,7}
    linkedList1 := CreateListByArray(arr1)

    l := deleteDuplicates(linkedList1.Head.Next)
    tmpList := NewLinkedList(&ListNode{Next:l})
    tmpList.Print()
}

//go test -v DeleteDuplicates.go DeleteDuplicates_test.go LinkedList.go --test.run=TestDeleteDuplicates3
func TestDeleteDuplicates3(t *testing.T) {
    arr1 := []int{1,1}
    linkedList1 := CreateListByArray(arr1)

    l := deleteDuplicates(linkedList1.Head.Next)
    tmpList := NewLinkedList(&ListNode{Next:l})
    tmpList.Print()
}