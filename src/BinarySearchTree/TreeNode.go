package _4_tree

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func NewNode(v int) *TreeNode {
	return &TreeNode{Val: v}
}

func (this *TreeNode) String() string {
	return fmt.Sprintf("v:%+v, left:%+v, right:%+v", this.Val, this.Left, this.Right)
}
