package LinkedList

import "testing"

//go test -v LinkedList.go GetIntersectionNode.go GetIntersectionNode_test.go --test.run=TestGetIntersectionNode
func TestGetIntersectionNode(t *testing.T) {
	arr1 := []int{2,8,1,3,4,5}
	linkedList1 := CreateListByArray(arr1)
	arr2 := []int{9,10,11}
	linkedList2 := CreateListByArray(arr2)

	linkedList2.Head.Next.Next.Next.Next = linkedList1.Head.Next.Next.Next
	result := getIntersectionNode(linkedList1.Head.Next, linkedList2.Head.Next)
	t.Log(result.Val)
}