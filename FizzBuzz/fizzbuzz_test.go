package FizzBuzz

import "testing"

func TestA(t *testing.T) {
	var ret []string
	ret = fizzBuzz(15)
	t.Log(ret)
}
