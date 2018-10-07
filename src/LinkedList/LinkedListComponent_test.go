package LinkedList

import "testing"

//go test -v LinkedList.go LinkedListComponent.go LinkedListComponent_test.go --test.run=TestLinkedListComponent1
func TestLinkedListComponent1(t *testing.T) {
	arr1 := []int{1,2,3,4,5}
	linkedList1 := CreateListByArray(arr1)
	h := numComponents(linkedList1.Head.Next, []int{1,2,4})
	t.Log(h)
}

//go test -v LinkedList.go LinkedListComponent.go LinkedListComponent_test.go --test.run=TestLinkedListComponent2
func TestLinkedListComponent2(t *testing.T) {
	arr1 := []int{1,2,3,4,5}
	linkedList1 := CreateListByArray(arr1)
	h := numComponents(linkedList1.Head.Next, []int{1,2,3})
	t.Log(h)
}
