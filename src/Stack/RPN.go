package Stack

import "strconv"

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	stack := NewArrayStack()
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			oper1, _ := stack.Pop().(int)
			oper2, _ := stack.Pop().(int)
			var result int
			switch(token) {
			case "+":
				result = oper2 + oper1
			case "-":
				result = oper2 - oper1
			case "*":
				result = oper2 * oper1
			case "/":
				result = oper2 / oper1
			}
			stack.Push(result)
		} else {
			iToken, _ := strconv.Atoi(token)
			stack.Push(iToken)
		}
	}

	return stack.Pop().(int)
}
