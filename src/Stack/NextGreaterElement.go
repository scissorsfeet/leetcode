package Stack

func nextGreaterElement(findNums []int, nums []int) []int {
	m := make(map[int]int)
	stack := NewArrayStack()

	for i := range nums {
		for !stack.IsEmpty() && stack.Top().(int) < nums[i] {
			top := stack.Pop().(int)
			m[top] = nums[i]
		}
		stack.Push(nums[i])
	}

	ret := make([]int, len(findNums))
	for i := range findNums {
		if _, ok := m[findNums[i]]; ok {
			ret[i] = m[findNums[i]]
		} else {
			ret[i] = -1
		}
	}
	return ret
}

