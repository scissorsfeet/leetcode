package LinkedList
import "testing"

//go test -v LinkedList.go OddEvenList.go OddEvenList_test.go --test.run=TestOddEvenList1
func TestOddEvenList1(t *testing.T) {
    arr1 := []int{1,2,3,4,5,6,7,8}
    linkedList1 := CreateListByArray(arr1)
    linkedList1.Head.Next = oddEvenList(linkedList1.Head.Next)
    linkedList1.Print()
}

//go test -v LinkedList.go OddEvenList.go OddEvenList_test.go --test.run=TestOddEvenList2
func TestOddEvenList2(t *testing.T) {
    arr1 := []int{1,2,3}
    linkedList1 := CreateListByArray(arr1)
    linkedList1.Head.Next = oddEvenList(linkedList1.Head.Next)
    linkedList1.Print()
}

//go test -v LinkedList.go OddEvenList.go OddEvenList_test.go --test.run=TestOddEvenList3
func TestOddEvenList3(t *testing.T) {
    arr1 := []int{1,2}
    linkedList1 := CreateListByArray(arr1)
    linkedList1.Head.Next = oddEvenList(linkedList1.Head.Next)
    linkedList1.Print()
}

//go test -v LinkedList.go OddEvenList.go OddEvenList_test.go --test.run=TestOddEvenList4
func TestOddEvenList4(t *testing.T) {
    arr1 := []int{1}
    linkedList1 := CreateListByArray(arr1)
    linkedList1.Head.Next = oddEvenList(linkedList1.Head.Next)
    linkedList1.Print()
}