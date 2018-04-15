package LinkedList

import "testing"

//go test -v LinkedList.go HasCycle.go HasCycle_test.go --test.run=TestHasCycle1
func TestHasCycle1(t *testing.T) {
	arr1 := []int{2,2,1}
	linkedList1 := CreateListByArray(arr1)

	linkedList1.Head.Next.Next.Next = linkedList1.Head.Next.Next
	result := hasCycle(linkedList1.Head)
	if result != true {
		t.Fail()
	}
}

//go test -v LinkedList.go HasCycle.go HasCycle_test.go --test.run=TestHasCycle2
func TestHasCycle2(t *testing.T) {
	arr1 := []int{2,2,1}
	linkedList1 := CreateListByArray(arr1)

	//linkedList1.Head.Next.Next.Next = linkedList1.Head.Next.Next
	result := hasCycle(linkedList1.Head)
	if result != false {
		t.Fail()
	}
}
