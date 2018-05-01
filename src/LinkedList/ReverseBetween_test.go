package LinkedList
import "testing"

//go test -v LinkedList.go ReverseBetween.go ReverseBetween_test.go --test.run=TestReverseBetweenList1
func TestReverseBetweenList1(t *testing.T) {
    arr1 := []int{1,2,3,4,5}
    linkedList1 := CreateListByArray(arr1)
    linkedList1.Head.Next = reverseBetween(linkedList1.Head.Next,2,4)
    linkedList1.Print()
}

//go test -v LinkedList.go ReverseBetween.go ReverseBetween_test.go --test.run=TestReverseBetweenList2
func TestReverseBetweenList2(t *testing.T) {
    arr1 := []int{1}
    linkedList1 := CreateListByArray(arr1)
    linkedList1.Head.Next = reverseBetween(linkedList1.Head.Next,1,1)
    linkedList1.Print()
}

//go test -v LinkedList.go ReverseBetween.go ReverseBetween_test.go --test.run=TestReverseBetweenList3
func TestReverseBetweenList3(t *testing.T) {
    arr1 := []int{3,5}
    linkedList1 := CreateListByArray(arr1)
    linkedList1.Head.Next = reverseBetween(linkedList1.Head.Next,1,2)
    linkedList1.Print()
}