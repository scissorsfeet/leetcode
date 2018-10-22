package Queue

type MyStack struct {
	q *MyCircularQueue
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	q := NewMyCircularQueue(1000000)
	return MyStack{&q}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	size := this.q.size
	this.q.EnQueue(x)
	for ;size>0;size-- {
		this.q.EnQueue(this.q.Front())
		this.q.DeQueue()
	}
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.q.IsEmpty() {
		return -1
	}
	v := this.q.Front()
	this.q.DeQueue()
	return v
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.q.Front()
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.q.IsEmpty()
}