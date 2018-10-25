package Stack

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	res := make([]int, n)
	for i:=0;i<n;i++{
		res[i] = -1
	}

	s := NewArrayStack()
	//第一轮，从后往前，分段大数沉底
	for i:=n-1;i>=0;i-- {
		for !s.IsEmpty() && s.Top().(int) <= nums[i] {
			s.Pop()
		}
		s.Push(nums[i])
	}

	//第二轮，从后往前，后面的数和前面的大数进行比较，起到循环作用
	for i:=n-1;i>=0;i-- {
		for !s.IsEmpty() && s.Top().(int) <= nums[i] {
			s.Pop()
		}
		if !s.IsEmpty() {
			res[i] = s.Top().(int)
		}
		s.Push(nums[i])
	}

	return res
}
