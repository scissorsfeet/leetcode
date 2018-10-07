package LinkedList

import "testing"

func TestMiddleNode(t *testing.T) {
	arr1 := []int{1,2,3,4,5,6}
	linkedList1 := CreateListByArray(arr1)
	t.Log(middleNode(linkedList1.Head.Next))
}