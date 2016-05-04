package MoveZeroes

import (
	"testing"
)

func TestA(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	t.Log(nums)
}

func TestB(t *testing.T) {
	nums := []int{}
	moveZeroes(nums)
	t.Log(nums)
}

func TestC(t *testing.T) {
	nums := []int{0, 0, 0}
	moveZeroes(nums)
	t.Log(nums)
}

func TestD(t *testing.T) {
	nums := []int{0, 0, 0, 1}
	moveZeroes(nums)
	t.Log(nums)
}
