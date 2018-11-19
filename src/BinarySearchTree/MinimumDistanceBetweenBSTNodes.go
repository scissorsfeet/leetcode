package _4_tree

import "math"

func minDiffInBST(root *TreeNode) int {
	s := NewArrayStack()
	var pre *TreeNode = nil
	p := root
	res := math.MaxInt32

	for !s.IsEmpty() || nil != p {
		if nil != p {
			s.Push(p)
			p = p.Left
		} else {
			tmp := s.Pop().(*TreeNode)
			if nil != pre {
				diff := tmp.Val - pre.Val
				if diff < res {
					res = diff
				}
			}
			pre = tmp
			p = tmp.Right
		}
	}

	return res
}

func minDiffInBSTII(root *TreeNode) int {
	var res = new(int)
	var pre *int
	*res = math.MaxInt32
	helper(root, pre, res)
	return *res
}

func helper(p *TreeNode, pre, res *int){
	if nil == p {
		return
	}

	helper(p.Left, pre, res)
	if nil != pre {
		if p.Val - *pre < *res {
			*res = p.Val - *pre
		}
	}
	if nil == pre {
		pre = new(int)
	}
	*pre = p.Val
	helper(p.Right, pre, res)
}