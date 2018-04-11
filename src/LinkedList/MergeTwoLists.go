package LinkedList

//合并两个已排序的链表，并将其作为一个新列表返回。新列表应该通过拼接前两个列表的节点来完成。

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if nil == l1 {
		return l2
	} else if nil == l2 {
		return l1
	}
	//p1是第一个结点较小的链表
	var p1, p2, ret *ListNode
	if l1.Val > l2.Val {
		p1 = l2
		p2 = l1
		ret = l2
	} else {
		p1 = l1
		p2 = l2
		ret = l1
	}

	var preP1 *ListNode
	for nil != p1 && nil != p2 {
		for nil != p1 && nil != p2 && p1.Val <= p2.Val {
			preP1 = p1
			p1 = p1.Next
		}
		preP1.Next = p2

		tmp := p1
		p1 = p2
		p2 = tmp
	}

	return ret
}