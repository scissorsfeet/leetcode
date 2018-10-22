package Stack

type backspaceComparer struct {
	str string
	s *ArrayStack
}

func (this *backspaceComparer) init() {
	if nil == this.s {
		this.s = NewArrayStack()
	}
	for i := range this.str {
		if string(this.str[i]) == "#" {
			if !this.s.IsEmpty() {
				this.s.Pop()
			}
		} else {
			this.s.Push(this.str[i])
		}
	}
}

func backspaceCompare(S string, T string) bool {
	sComparer := backspaceComparer{str:S}
	sComparer.init()
	tComparer := backspaceComparer{str:T}
	tComparer.init()

	for !sComparer.s.IsEmpty() && !tComparer.s.IsEmpty() {
		s := sComparer.s.Pop()
		t := tComparer.s.Pop()
		if s.(uint8) != t.(uint8) {
			return false
		}
	}

	if !sComparer.s.IsEmpty() || !tComparer.s.IsEmpty() {
		return false
	}

	return true
}
