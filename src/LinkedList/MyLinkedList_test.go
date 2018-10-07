package LinkedList

import (
	"fmt"
	"testing"
)

func (this *MyLinkedList) Print(isFromTail bool) {
	var s string
	if !isFromTail {
		cur := this.head.next
		for nil != cur {
			s += fmt.Sprintf("%+v", cur.val)
			if nil != cur.next {
				s += "->"
			}
			cur = cur.next
		}
	} else {
		cur := this.tail
		for this.head != cur {
			s += fmt.Sprintf("%+v", cur.val)
			if this.head != cur.prev {
				s += "->"
			}
			cur = cur.prev
		}
	}
	fmt.Println(s)
}

func TestMyLinkedList_AddAtHead(t *testing.T) {
	l := Constructor()
	l.AddAtHead(1)
	l.AddAtHead(2)
	l.AddAtHead(3)
	l.Print(false)
	l.Print(true)
}

func TestMyLinkedList_AddAtTail(t *testing.T) {
	l := Constructor()
	l.AddAtTail(1)
	l.AddAtTail(2)
	l.AddAtTail(3)
	l.Print(false)
	l.Print(true)
}

func TestMyLinkedList_Get(t *testing.T) {
	l := Constructor()
	l.AddAtTail(1)
	l.AddAtTail(2)
	l.AddAtTail(3)
	t.Log(l.Get(0))
	t.Log(l.Get(1))
	t.Log(l.Get(2))
	t.Log(l.Get(3))
}

func TestMyLinkedList_AddAtIndex(t *testing.T) {
	l := Constructor()
	l.AddAtTail(1)
	l.AddAtTail(2)
	l.AddAtTail(3)
	l.Print(false)
	l.Print(true)
	t.Log(l.length)
	l.AddAtIndex(0,10)
	l.Print(false)
	l.Print(true)
	t.Log(l.length)
	l.AddAtIndex(4,11)
	l.Print(false)
	l.Print(true)
	t.Log(l.length)
}

func TestMyLinkedList_DeleteAtIndex(t *testing.T) {
	l := Constructor()
	l.AddAtTail(1)
	l.AddAtTail(2)
	l.AddAtTail(3)
	l.Print(false)
	l.Print(true)
	t.Log(l.length)
	l.DeleteAtIndex(0)
	l.Print(false)
	l.Print(true)
	t.Log(l.length)
	l.DeleteAtIndex(1)
	l.Print(false)
	l.Print(true)
	t.Log(l.length)
	l.DeleteAtIndex(0)
	l.Print(false)
	l.Print(true)
	t.Log(l.length)
}

func TestA(t *testing.T) {
	l := Constructor()
	t.Log(l.Get(0))
	l.AddAtIndex(1,2)
	l.Print(false)
	t.Log(l.Get(0))
	t.Log(l.Get(1))
	l.AddAtIndex(0,1)
	l.Print(false)
	t.Log(l.Get(0))
	t.Log(l.Get(1))
}