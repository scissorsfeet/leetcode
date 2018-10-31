package Recursive

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func allPossibleFBT(N int) []*TreeNode {
	if N < 0 || N % 2 == 0 {
		return []*TreeNode{}
	}
	if N == 1 {
		return []*TreeNode{&TreeNode{Val:0}}
	}
	nodes := make([]*TreeNode, 0, 8)
	N=N-1
	for i:=1;i<N;i+=2 {
		lNodes := allPossibleFBT(i)
		rNodes := allPossibleFBT(N-i)
		for l := range lNodes {
			for r := range rNodes {
				root := &TreeNode{Val:0}
				root.Left = lNodes[l]
				root.Right = rNodes[r]
				nodes = append(nodes, root)
			}
		}
	}
	return nodes
}


