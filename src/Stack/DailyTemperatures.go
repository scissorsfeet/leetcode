package Stack

type Item struct {
	val int
	shiftDays int
	index int
}

func dailyTemperatures(temperatures []int) []int {
	if len(temperatures) == 0 {
		return temperatures
	}

	stack := NewArrayStack()
	assist := NewArrayStack()
	result := make([]int, len(temperatures))
	for i, temperature := range temperatures {
		top := stack.Top()
		if nil != top {
			iTop := top.(*Item)
			if iTop.val < temperature {
				for iTop.val < temperature {
					stack.Pop()
					assist.Push(iTop)
					if !stack.IsEmpty() {
						stack.Top().(*Item).shiftDays += iTop.shiftDays
						iTop = stack.Top().(*Item)
					} else {
						break
					}
				}
			} else {
				assist.Print()
				for !assist.IsEmpty() {
					result = append(result, assist.Pop().(*Item).shiftDays)
				}
			}
		}
		stack.Push(&Item{temperature, 1, i})
	}

	for !assist.IsEmpty() {
		top := assist.Pop().(*Item)
		result[top.index] = top.shiftDays
	}

	for !stack.IsEmpty() {
		top := stack.Pop().(*Item)
		result[top.index] = 0
	}

	return result
}
