package LinkedList

import (
	"testing"
)

func TestA(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	head := CreateByArray(arr)
	head.ToString()
}
