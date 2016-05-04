package DeleteNodeInLinkedList

import (
	"testing"

	linkedlist "github.com/liutianyi1989/leetcode/LinkedList"
)

func TestA(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	head := linkedlist.CreateByArray(arr)
	head.ToString()
	solution(head.Next.Next)
	head.ToString()
}
