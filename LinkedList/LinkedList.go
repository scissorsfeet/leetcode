package LinkedList

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

func newNode(val int, next *Node) (node *Node) {
	return &Node{val, next}
}

func CreateByArray(items []int) (head *Node) {
	itemsLen := len(items)
	if itemsLen <= 0 {
		head = nil
	} else if itemsLen == 1 {
		head = newNode(items[0], nil)
	} else {
		var pre *Node
		for _, v := range items {
			cur := newNode(v, nil)
			if pre == nil {
				pre = cur
				head = cur
			} else {
				pre.Next = cur
				pre = cur
			}
		}
	}
	return
}

func (this *Node) ToString() {
	var arr []int
	pNode := this
	for pNode != nil {
		arr = append(arr, pNode.Val)
		pNode = pNode.Next
	}
	fmt.Println(arr)
}
