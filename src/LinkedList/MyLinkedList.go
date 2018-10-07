package LinkedList

type MyListNode struct {
	val int
	next *MyListNode
	prev *MyListNode
}

type MyLinkedList struct {
	head *MyListNode
	tail *MyListNode
	length int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{&MyListNode{}, nil,0}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.length {
		return -1
	}

	cur := this.head
	for i:=0;i<=index;i++ {
		cur = cur.next
	}

	return cur.val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	if nil == this.head {
		this.head = &MyListNode{}
	}

	newNode := &MyListNode{val:val}
	oldNext := this.head.next
	this.head.next = newNode
	newNode.prev = this.head
	newNode.next = oldNext
	if nil != oldNext {
		oldNext.prev = newNode
	} else {
		this.tail = newNode
	}
	this.length++
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	newNode := &MyListNode{val:val}
	if nil == this.tail {
		this.head.next = newNode
		newNode.prev = this.head
	} else {
		this.tail.next = newNode
		newNode.prev = this.tail
	}
	this.tail = newNode
	this.length++
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index < 0 || index > this.length {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.length {
		this.AddAtTail(val)
		return
	}

	newNode := &MyListNode{val:val}
	cur := this.head
	for i:=0;i<=index;i++{
		cur = cur.next
	}

	oldPrev := cur.prev
	oldPrev.next = newNode
	newNode.prev = oldPrev
	newNode.next = cur
	cur.prev = newNode
	this.length++
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 || index >= this.length {
		return
	}

	cur := this.head
	for i:=0;i<=index;i++ {
		cur = cur.next
	}

	prev := cur.prev
	next := cur.next
	prev.next = next
	if nil != next {
		next.prev = prev
	} else {
		this.tail = prev
	}
	this.length--
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */