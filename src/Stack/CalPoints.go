package Stack

import "strconv"

func calPoints(ops []string) int {
	stack := NewArrayStack()

	for i := range ops {
		switch ops[i] {
		case "+":
			a := stack.Pop().(int)
			b := stack.Pop().(int)
			stack.Push(b)
			stack.Push(a)
			stack.Push(a+b)
		case "C":
			stack.Pop()
		case "D":
			stack.Push(stack.Top().(int)*2)
		default:
			point, _ := strconv.Atoi(ops[i])
			stack.Push(point)
		}
	}

	ret := 0
	for !stack.IsEmpty() {
		ret += stack.Pop().(int)
	}
	return ret
}
