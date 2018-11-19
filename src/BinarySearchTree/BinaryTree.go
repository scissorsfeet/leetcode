package _4_tree

import "fmt"

type BinaryTree struct {
	root *TreeNode
}

func NewBinaryTree(rootV int) *BinaryTree {
	return &BinaryTree{NewNode(rootV)}
}

func (this *BinaryTree) InOrderTraverse() {
	p := this.root
	s := NewArrayStack()

	for !s.IsEmpty() || nil != p {
		if nil != p {
			s.Push(p)
			p = p.Left
		} else {
			tmp := s.Pop().(*TreeNode)
			fmt.Printf("%+v ", tmp.Val)
			p = tmp.Right
		}
	}
	fmt.Println()
}

func (this *BinaryTree) PreOrderTraverse() {
	p := this.root
	s := NewArrayStack()

	for !s.IsEmpty() || nil != p {
		if nil != p {
			fmt.Printf("%+v ", p.Val)
			s.Push(p)
			p = p.Left
		} else {
			p = s.Pop().(*TreeNode).Right
		}
	}

	fmt.Println()
}

func (this *BinaryTree) PostOrderTraverse() {
	s1 := NewArrayStack()
	s2 := NewArrayStack()
	s1.Push(this.root)
	for !s1.IsEmpty() {
		p := s1.Pop().(*TreeNode)
		s2.Push(p)
		if nil != p.Left {
			s1.Push(p.Left)
		}
		if nil != p.Right {
			s1.Push(p.Right)
		}
	}

	for !s2.IsEmpty() {
		fmt.Printf("%+v ", s2.Pop().(*TreeNode).Val)
	}
}
