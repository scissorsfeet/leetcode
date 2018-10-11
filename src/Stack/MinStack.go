package Stack

type MinStack struct {
	data []int
	min []int
	top int
	minVal int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data : make([]int,0,32),
		min : make([]int, 0, 32),
		top : -1,
		minVal : int(^uint(0)>>1),
	}
}

func (this *MinStack) isEmpty() bool {
	if this.top < 0 {
		return true
	}
	return false
}

func (this *MinStack) Push(x int) {
	if this.isEmpty() {
		this.top = 0
	} else {
		this.top += 1
	}

	if this.minVal > x {
		this.minVal = x
	}

	if this.top > len(this.data)-1 {
		this.data = append(this.data, x)
		this.min = append(this.min, this.minVal)
	} else {
		this.data[this.top] = x
		this.min[this.top] = this.minVal
	}
}

func (this *MinStack) Pop()  {
	if this.isEmpty() {
		return
	}
	this.top -= 1
	if !this.isEmpty() {
		this.minVal = this.min[this.top]
	} else {
		this.minVal = int(^uint(0)>>1)
	}
}

func (this *MinStack) Top() int {
	if this.isEmpty() {
		return ^int(^uint(0)>>1)
	}
	return this.data[this.top]
}


func (this *MinStack) GetMin() int {
	if this.isEmpty() {
		return ^int(^uint(0)>>1)
	}
	return this.minVal
}