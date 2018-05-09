package LinkedList

import "testing"

//go test -v LinkedList.go DetectCycle.go DetectCycle_test.go --test.run=TestDetectCycle1
func TestDetectCycle1(t *testing.T) {
	arr1 := []int{3,2,1}
	linkedList1 := CreateListByArray(arr1)

	linkedList1.Head.Next.Next.Next = linkedList1.Head.Next.Next
	result := detectCycle(linkedList1.Head)
	t.Log(result.Val)
}

//go test -v LinkedList.go DetectCycle.go DetectCycle_test.go --test.run=TestDetectCycle2
func TestDetectCycle2(t *testing.T) {
	arr1 := []int{2,2,1}
	linkedList1 := CreateListByArray(arr1)

	//linkedList1.Head.Next.Next.Next = linkedList1.Head.Next.Next
	result := detectCycle(linkedList1.Head)
	if nil != result {
		t.Fatal(result.Val)
	}
}

//go test -v LinkedList.go DetectCycle.go DetectCycle_test.go --test.run=TestDetectCycle3ss
func TestDetectCycle3(t *testing.T) {
	arr1 := []int{}
	linkedList1 := CreateListByArray(arr1)

	//linkedList1.Head.Next.Next.Next = linkedList1.Head.Next.Next
	result := detectCycle(linkedList1.Head)
	if nil != result {
		t.Fatal(result.Val)
	}
}
