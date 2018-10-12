package Queue

import "fmt"

type MyCircularQueue struct {
	q []int
	capacity int
	head int
	tail int
}


/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	if k <= 0 {
		return MyCircularQueue{}
	}
	capacity := k+1
	return MyCircularQueue{q:make([]int, capacity), capacity:capacity}
}


/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.q[this.tail] = value
	this.tail = (this.tail+1)%this.capacity
	return true
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head+1)%this.capacity
	return true
}


/** Get the front item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return ^int(^uint(0)>>1)
	}
	i := (this.tail-1+this.capacity)%this.capacity
	return this.q[i]
}


/** Get the last item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return ^int(^uint(0)>>1)
	}
	fmt.Println(this.q, this.head, this.tail)
	return this.q[this.head]
}


/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	if this.head == this.tail {
		return true
	}
	return false
}


/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if (this.tail+1)%this.capacity == this.head {
		return true
	}
	return false
}
