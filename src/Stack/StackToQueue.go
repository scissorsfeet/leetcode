package Stack

type MyQueue struct {
	tailStack *ArrayStack
	headStack *ArrayStack
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{NewArrayStack(), NewArrayStack()}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.tailStack.Push(x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	var ret = -1
	if this.headStack.IsEmpty() {
		for !this.tailStack.IsEmpty() {
			this.headStack.Push(this.tailStack.Pop())
		}
	}
	if !this.headStack.IsEmpty() {
		ret = this.headStack.Pop().(int)
	}
	return ret
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	var ret = -1
	if this.headStack.IsEmpty() {
		for !this.tailStack.IsEmpty() {
			this.headStack.Push(this.tailStack.Pop())
		}
	}
	if !this.headStack.IsEmpty() {
		ret = this.headStack.Top().(int)
	}
	return ret
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if this.headStack.IsEmpty() && this.tailStack.IsEmpty() {
		return true
	}
	return false
}