package LinkedList
import "testing"

//go test -v RemoveLinkedListElements.go RemoveLinkedListElements_test.go LinkedList.go --test.run=TestRemoveElements1
func TestRemoveElements1(t *testing.T)  {
    arr1 := []int{1,3,3,3,5,7,7,9}
    linkedList1 := CreateListByArray(arr1)

    l := removeElements(linkedList1.Head.Next, 3)
    tmpList := NewLinkedList(&ListNode{Next:l})
    tmpList.Print()
}

//go test -v RemoveLinkedListElements.go RemoveLinkedListElements_test.go LinkedList.go --test.run=TestRemoveElements2
func TestRemoveElements2(t *testing.T)  {
    arr1 := []int{3,3,3,5,7,7,9}
    linkedList1 := CreateListByArray(arr1)

    l := removeElements(linkedList1.Head.Next, 3)
    tmpList := NewLinkedList(&ListNode{Next:l})
    tmpList.Print()
}

//go test -v RemoveLinkedListElements.go RemoveLinkedListElements_test.go LinkedList.go --test.run=TestRemoveElements3
func TestRemoveElements3(t *testing.T)  {
    arr1 := []int{3}
    linkedList1 := CreateListByArray(arr1)

    l := removeElements(linkedList1.Head.Next, 3)
    tmpList := NewLinkedList(&ListNode{Next:l})
    tmpList.Print()
}