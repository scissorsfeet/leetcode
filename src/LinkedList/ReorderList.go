package LinkedList

func reorderList(head *ListNode)  {
	if nil == head || nil == head.Next || nil == head.Next.Next {
		return
	}

	//先找到中间节点
	slow := head
	fast := head
	for nil != fast && nil != fast.Next {
		slow = slow.Next
		fast = fast.Next.Next
	}

	//反转中间节点后面的节点
	pre := slow
	cur := slow.Next
	for nil != cur {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	//两端向中间走
	p := head
	q := pre
	for p != slow {
		tmpLeft := p.Next
		tmpRight := q.Next

		p.Next = q
		q.Next = tmpLeft

		p = tmpLeft
		q = tmpRight
	}

	p.Next = nil

	return
}