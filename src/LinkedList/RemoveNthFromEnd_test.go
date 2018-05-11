package LinkedList
import "testing"

//go test -v RemoveNthFromEnd.go RemoveNthFromEnd_test.go LinkedList.go --test.run=TestRemoveNthFromEnd1
func TestRemoveNthFromEnd1(t *testing.T)  {
    arr1 := []int{1,3,3,3,5,7,7,9}
    linkedList1 := CreateListByArray(arr1)

    l := removeNthFromEnd(linkedList1.Head.Next, 3)
    tmpList := NewLinkedList(&ListNode{Next:l})
    tmpList.Print()
}

//go test -v RemoveNthFromEnd.go RemoveNthFromEnd_test.go LinkedList.go --test.run=TestRemoveNthFromEnd2
func TestRemoveNthFromEnd2(t *testing.T)  {
    arr1 := []int{1,3,3,3,5,7,7,9}
    linkedList1 := CreateListByArray(arr1)

    l := removeNthFromEnd(linkedList1.Head.Next, 1)
    tmpList := NewLinkedList(&ListNode{Next:l})
    tmpList.Print()
}

//go test -v RemoveNthFromEnd.go RemoveNthFromEnd_test.go LinkedList.go --test.run=TestRemoveNthFromEnd3
func TestRemoveNthFromEnd3(t *testing.T)  {
    arr1 := []int{1,3,3,3,5,7,7,9}
    linkedList1 := CreateListByArray(arr1)

    l := removeNthFromEnd(linkedList1.Head.Next, 8)
    tmpList := NewLinkedList(&ListNode{Next:l})
    tmpList.Print()
}
//go test -v RemoveNthFromEnd.go RemoveNthFromEnd_test.go LinkedList.go --test.run=TestRemoveNthFromEnd4
func TestRemoveNthFromEnd4(t *testing.T)  {
    arr1 := []int{1}
    linkedList1 := CreateListByArray(arr1)

    l := removeNthFromEnd(linkedList1.Head.Next, 1)
    tmpList := NewLinkedList(&ListNode{Next:l})
    tmpList.Print()
}
