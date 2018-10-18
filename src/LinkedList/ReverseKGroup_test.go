package LinkedList

import (
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	arr1 := []int{1,2,3,4,5}
	linkedList1 := CreateListByArray(arr1)
	linkedList1.Print()
	h := reverseKGroup(linkedList1.Head.Next,2)
	l := &LinkedList{Head:&ListNode{Next:h}}
	l.Print()

	arr2 := []int{1,2,3,4,5}
	linkedList2 := CreateListByArray(arr2)
	linkedList2.Print()
	h2 := reverseKGroup(linkedList2.Head.Next,3)
	l2 := &LinkedList{Head:&ListNode{Next:h2}}
	l2.Print()
}