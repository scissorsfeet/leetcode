package LinkedList
import "testing"

//go test -v LinkedList.go SplitListToParts.go SplitListToParts_test.go --test.run=TestSplitListToParts1
func TestSplitListToParts1(t *testing.T) {
    arr1 := []int{1,2,3,4,5}
    linkedList1 := CreateListByArray(arr1)
    ret := splitListToParts(linkedList1.Head.Next, 2)
    for i, _ := range ret {
        list := LinkedList{Head:&ListNode{}}
        list.Head.Next = ret[i]
        list.Print()
    }
}

//go test -v LinkedList.go SplitListToParts.go SplitListToParts_test.go --test.run=TestSplitListToParts2
func TestSplitListToParts2(t *testing.T) {
    arr1 := []int{1,2}
    linkedList1 := CreateListByArray(arr1)
    ret := splitListToParts(linkedList1.Head.Next, 2)
    for i, _ := range ret {
        list := LinkedList{Head:&ListNode{}}
        list.Head.Next = ret[i]
        list.Print()
    }
}

//go test -v LinkedList.go SplitListToParts.go SplitListToParts_test.go --test.run=TestSplitListToParts3
func TestSplitListToParts3(t *testing.T) {
    arr1 := []int{1,2,3,4,5,1,2,3,4,5}
    linkedList1 := CreateListByArray(arr1)
    ret := splitListToParts(linkedList1.Head.Next, 3)
    for i, _ := range ret {
        list := LinkedList{Head:&ListNode{}}
        list.Head.Next = ret[i]
        list.Print()
    }
}

//go test -v LinkedList.go SplitListToParts.go SplitListToParts_test.go --test.run=TestSplitListToParts4
func TestSplitListToParts4(t *testing.T) {
    arr1 := []int{1,2,3,4}
    linkedList1 := CreateListByArray(arr1)
    ret := splitListToParts(linkedList1.Head.Next, 5)
    for i, _ := range ret {
        list := LinkedList{Head:&ListNode{}}
        list.Head.Next = ret[i]
        list.Print()
    }
}

//go test -v LinkedList.go SplitListToParts.go SplitListToParts_test.go --test.run=TestSplitListToParts5
func TestSplitListToParts5(t *testing.T) {
    arr1 := []int{}
    linkedList1 := CreateListByArray(arr1)
    ret := splitListToParts(linkedList1.Head.Next, 3)
    for i, _ := range ret {
        list := LinkedList{Head:&ListNode{}}
        list.Head.Next = ret[i]
        list.Print()
    }
}