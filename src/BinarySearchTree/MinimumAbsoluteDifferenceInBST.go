package _4_tree

import "math"

func getMinimumDifference(root *TreeNode) int {
	var res = math.MaxInt32
	stack := NewArrayStack()
	p := root
	var pre *TreeNode

	for !stack.IsEmpty() || nil != p {
		if nil != p {
			stack.Push(p)
			p = p.Left
		} else {
			t := stack.Pop().(*TreeNode)
			if nil != pre && t.Val-pre.Val<res{
				res = t.Val-pre.Val
			}
			pre = t
		}
	}

	return res
}
