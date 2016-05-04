package SameTree

import (
	"testing"

	binaryTree "github.com/liutianyi1989/leetcode/BinaryTree"
)

func TestA(t *testing.T) {
	arrA := []int{1, 2, 3, 4}
	arrB := []int{1, 2, 3, 3}
	nodeA := binaryTree.CreateByArray(arrA)
	nodeB := binaryTree.CreateByArray(arrB)
	isSame := isSameTree(nodeA, nodeB)
	if isSame {
		t.Fail()
	}
}

func TestB(t *testing.T) {
	arrA := []int{}
	arrB := []int{}
	nodeA := binaryTree.CreateByArray(arrA)
	nodeB := binaryTree.CreateByArray(arrB)
	isSame := isSameTree(nodeA, nodeB)
	if !isSame {
		t.Fail()
	}
}

func TestC(t *testing.T) {
	arrA := []int{1, 2, 3, -1, 4}
	arrB := []int{1, 2, 3, -1, 4}
	nodeA := binaryTree.CreateByArray(arrA)
	nodeB := binaryTree.CreateByArray(arrB)
	isSame := isSameTree(nodeA, nodeB)
	if !isSame {
		t.Fail()
	}
}
