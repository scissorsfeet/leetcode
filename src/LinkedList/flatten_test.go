package LinkedList

import (
	"fmt"
	"testing"
)

func (this *Node) Print() {
	var s string
	cur := this
	for nil != cur {
		s += fmt.Sprintf("%+v", cur.val)
		if nil != cur.next {
			s += "->"
		}
		cur = cur.next
	}
	fmt.Println(s)
}

func TestFlatten(t *testing.T) {
	n5 := &Node{val:5}
	n4 := &Node{val:4,next:n5}
	n5.prev = n4
	n3 := &Node{val:3,next:n4}
	n4.prev = n3
	n2 := &Node{val:2,next:n3}
	n3.prev = n2
	n1 := &Node{val:1,next:n2}
	n2.prev = n1
	n1.Print()

	n7 := &Node{val:7}
	n6 := &Node{val:6,next:n7}
	n7.prev = n6
	n2.child = n6

	n9 := &Node{val:9}
	n8 := &Node{val:8,next:n9}
	n9.prev = n8
	n5.child = n8

	n11 := &Node{val:11}
	n10 := &Node{val:10,next:n11}
	n11.prev = n10
	n6.child = n10

	flatten(n1)
	n1.Print()





}
