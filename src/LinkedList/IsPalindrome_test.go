package LinkedList

import "testing"

//go test -v LinkedList.go IsPalindrome.go IsPalindrome_test.go --test.run=TestIsPalindrome1
func TestIsPalindrome1(t *testing.T) {
	arr1 := []int{1,2,3,2,1}
	linkedList1 := CreateListByArray(arr1)
	result := isPalindrome(linkedList1.Head.Next)
	t.Log(result)
}

//go test -v LinkedList.go IsPalindrome.go IsPalindrome_test.go --test.run=TestIsPalindrome2
func TestIsPalindrome2(t *testing.T) {
	arr1 := []int{1}
	linkedList1 := CreateListByArray(arr1)
	result := isPalindrome(linkedList1.Head.Next)
	t.Log(result)
}

//go test -v LinkedList.go IsPalindrome.go IsPalindrome_test.go --test.run=TestIsPalindrome3
func TestIsPalindrome3(t *testing.T) {
	arr1 := []int{1,2,2,1}
	linkedList1 := CreateListByArray(arr1)
	result := isPalindrome(linkedList1.Head.Next)
	t.Log(result)
}

//go test -v LinkedList.go IsPalindrome.go IsPalindrome_test.go --test.run=TestIsPalindrome4
func TestIsPalindrome4(t *testing.T) {
	arr1 := []int{2,2}
	linkedList1 := CreateListByArray(arr1)
	result := isPalindrome(linkedList1.Head.Next)
	t.Log(result)
}

//go test -v LinkedList.go IsPalindrome.go IsPalindrome_test.go --test.run=TestIsPalindrome5
func TestIsPalindrome5(t *testing.T) {
	arr1 := []int{2,2,1}
	linkedList1 := CreateListByArray(arr1)
	result := isPalindrome(linkedList1.Head.Next)
	t.Log(result)
}

//go test -v LinkedList.go IsPalindrome.go IsPalindrome_test.go --test.run=TestIsPalindrome6
func TestIsPalindrome6(t *testing.T) {
	arr1 := []int{1,0,3,4,0,1}
	linkedList1 := CreateListByArray(arr1)
	result := isPalindrome(linkedList1.Head.Next)
	t.Log(result)
}

//go test -v LinkedList.go IsPalindrome.go IsPalindrome_test.go --test.run=TestIsPalindrome7
func TestIsPalindrome7(t *testing.T) {
	arr1 := []int{1,2}
	linkedList1 := CreateListByArray(arr1)
	result := isPalindrome(linkedList1.Head.Next)
	t.Log(result)
}