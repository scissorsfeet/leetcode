package LinkedList

import(
    "testing"
)

func TestCreateListByArray(t *testing.T) {
    arr1 := []int{1,2,3,4,5}
    linkedList := CreateListByArray(arr1)
    linkedList.Print()
}