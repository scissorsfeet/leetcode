package Stack

type Item struct {
	val int
	shiftDays int
	index int
}

/**
 时间复杂度:O(N)
 空间复杂度:O(N)
*/
func dailyTemperatures(temperatures []int) []int {
	if len(temperatures) == 0 {
		return temperatures
	}

	stack := NewArrayStack()
	result := make([]int, len(temperatures))
	for i, temperature := range temperatures {
		top := stack.Top()
		if nil != top {
			iTop := top.(*Item)
			for iTop.val < temperature {
				stack.Pop()
				result[iTop.index] = iTop.shiftDays
				if !stack.IsEmpty() {
					stack.Top().(*Item).shiftDays += iTop.shiftDays
					iTop = stack.Top().(*Item)
				} else {
					break
				}
			}
		}
		stack.Push(&Item{temperature, 1, i})
	}

	return result
}
