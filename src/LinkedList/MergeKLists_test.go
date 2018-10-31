package LinkedList

import "testing"

func TestMergeKLists(t *testing.T) {
	lists := make([]*ListNode, 0, 10)
	l1 := CreateListByArray([]int{1,3,5,6})
	l2 := CreateListByArray([]int{2,4,7})
	l3 := CreateListByArray([]int{-1,4,10})
	l4 := CreateListByArray([]int{2,5,19})
	lists = append(lists, l1.Head.Next)
	lists = append(lists, l2.Head.Next)
	lists = append(lists, l3.Head.Next)
	lists = append(lists, l4.Head.Next)

	l := &LinkedList{Head:&ListNode{Next:mergeKLists(lists)}}
	l.Print()
}