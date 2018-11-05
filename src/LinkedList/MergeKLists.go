package LinkedList

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	m := make(map[int]*ListNode, len(lists))
	for i := range lists {
		if nil == lists[i] {
			continue
		}
		m[i] = lists[i]
	}
	if len(m) == 0 {
		return nil
	}

	helper := &ListNode{}
	cur := helper
	for len(m) > 1 {
		var minPos = -1
		var minValue = 0
		for i := range m {
			if minPos < 0 || minValue > m[i].Val {
				minPos = i
				minValue = m[i].Val
			}
		}
		cur.Next = m[minPos]
		cur = cur.Next
		m[minPos] = m[minPos].Next
		if nil == m[minPos] {
			delete(m, minPos)
		}
	}
	for p := range m {
		cur.Next = m[p]
	}

	return helper.Next
}