package Stack

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2!=0 {
		return false
	}

	stack := NewArrayStack()
	for _, c := range s {
		cstr := string(c)
		if cstr == "{" || cstr == "(" || cstr == "[" {
			stack.Push(cstr)
		} else {
			if stack.IsEmpty() {
				return false
			}
			v := stack.Pop()
			if (cstr == "}" && v.(string) != "{") ||
				(cstr == "]" && v.(string) != "[") ||
				(cstr == ")" && v.(string) != "(") {
					return false
			}
		}
	}
	if !stack.IsEmpty() {
		return false
	}
	return true
}