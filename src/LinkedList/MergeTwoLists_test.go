package LinkedList

import "testing"

//合并两个已排序的链表，并将其作为一个新列表返回。新列表应该通过拼接前两个列表的节点来完成。

//go test -v LinkedList.go MergeTwoLists.go MergeTwoLists_test.go --test.run=TestMergeTwoLists1
func TestMergeTwoLists1(t *testing.T) {
	arr1 := []int{1,3,5,7,9}
	linkedList1 := CreateListByArray(arr1)

	arr2 := []int{2,4,6,8,10}
	linkedList2 := CreateListByArray(arr2)

	l := mergeTwoLists(linkedList1.Head.Next, linkedList2.Head.Next)
	tmpList := NewLinkedList(&ListNode{Next:l})
	tmpList.Print()
}

//go test -v LinkedList.go MergeTwoLists.go MergeTwoLists_test.go --test.run=TestMergeTwoLists2
func TestMergeTwoLists2(t *testing.T) {
	arr1 := []int{1}
	linkedList1 := CreateListByArray(arr1)

	arr2 := []int{2,4,6,8,10}
	linkedList2 := CreateListByArray(arr2)

	l := mergeTwoLists(linkedList1.Head.Next, linkedList2.Head.Next)
	tmpList := NewLinkedList(&ListNode{Next:l})
	tmpList.Print()
}

//go test -v LinkedList.go MergeTwoLists.go MergeTwoLists_test.go --test.run=TestMergeTwoLists3
func TestMergeTwoLists3(t *testing.T) {
	arr1 := []int{1,3,5,7,9}
	linkedList1 := CreateListByArray(arr1)

	arr2 := []int{2}
	linkedList2 := CreateListByArray(arr2)

	l := mergeTwoLists(linkedList1.Head.Next, linkedList2.Head.Next)
	tmpList := NewLinkedList(&ListNode{Next:l})
	tmpList.Print()
}