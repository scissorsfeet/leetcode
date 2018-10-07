package LinkedList

import(
    "testing"
)

//go test -v LinkedList.go RotateRight.go RotateRight_test.go --test.run=TestRotateRight1
func TestRotateRight1(t *testing.T) {
    arr1 := []int{1,2,3,4,5}
    linkedList1 := CreateListByArray(arr1)
    linkedList1.Head.Next = rotateRight(linkedList1.Head.Next, 2)
    linkedList1.Print()
}

//go test -v LinkedList.go RotateRight.go RotateRight_test.go --test.run=TestRotateRight2
func TestRotateRight2(t *testing.T) {
    arr1 := []int{1,2}
    linkedList1 := CreateListByArray(arr1)
    linkedList1.Head.Next = rotateRight(linkedList1.Head.Next, 1)
    linkedList1.Print()
}