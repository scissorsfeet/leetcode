package LinkedList

import "testing"

//go test -v AddTwoNumbersII.go AddTwoNumbersII_test.go LinkedList.go --test.run=TestAddTwoNumbers1
func TestAddTwoNumbers1(t *testing.T) {
	arr1 := []int{9,9,9,9}
	linkedList1 := CreateListByArray(arr1)

	arr2 := []int{1}
	linkedList2 := CreateListByArray(arr2)

	l := addTwoNumbers(linkedList1.Head.Next, linkedList2.Head.Next)
	tmpList := NewLinkedList(&ListNode{Next:l})
	tmpList.Print()
}

//go test -v AddTwoNumbersII.go AddTwoNumbersII_test.go LinkedList.go --test.run=TestAddTwoNumbers2
func TestAddTwoNumbers2(t *testing.T) {
	arr1 := []int{9,9,9,9}
	linkedList1 := CreateListByArray(arr1)

	l := addTwoNumbers(linkedList1.Head.Next, nil)
	tmpList := NewLinkedList(&ListNode{Next:l})
	tmpList.Print()
}

//go test -v AddTwoNumbersII.go AddTwoNumbersII_test.go LinkedList.go --test.run=TestAddTwoNumbers3
func TestAddTwoNumbers3(t *testing.T) {
	l := addTwoNumbers(nil, nil)
	tmpList := NewLinkedList(&ListNode{Next:l})
	tmpList.Print()
}

//go test -v AddTwoNumbersII.go AddTwoNumbersII_test.go LinkedList.go --test.run=TestAddTwoNumbers4
func TestAddTwoNumbers4(t *testing.T) {
	arr1 := []int{9,8,9,9}
	linkedList1 := CreateListByArray(arr1)

	arr2 := []int{1}
	linkedList2 := CreateListByArray(arr2)

	l := addTwoNumbers(linkedList1.Head.Next, linkedList2.Head.Next)
	tmpList := NewLinkedList(&ListNode{Next:l})
	tmpList.Print()
}

//go test -v AddTwoNumbersII.go AddTwoNumbersII_test.go LinkedList.go --test.run=TestAddTwoNumbers5
func TestAddTwoNumbers5(t *testing.T) {
	arr1 := []int{7,2,4,3}
	linkedList1 := CreateListByArray(arr1)

	arr2 := []int{5,6,4}
	linkedList2 := CreateListByArray(arr2)

	l := addTwoNumbers(linkedList1.Head.Next, linkedList2.Head.Next)
	tmpList := NewLinkedList(&ListNode{Next:l})
	tmpList.Print()
}