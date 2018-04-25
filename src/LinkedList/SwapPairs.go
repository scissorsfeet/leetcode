package LinkedList

/*
两两交换链表中的节点
*/
func swapPairs(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}
	//直接保存新的头结点用于返回
	retNode := head.Next

	//四元组,first、second为需要交换的两个结点，pre为first前驱结点，third为second后继结点
	var pre, first, second, third *ListNode = nil, head, head.Next, head.Next.Next
	for nil != first && nil != second {
		//先交换first和second
		if nil != pre {
			pre.Next = second
		}
		second.Next = first
		first.Next = third

		if nil == third || nil == third.Next{ //无需继续交换
			break
		}
		//前进
		pre = first
		first = third
		second = first.Next
		third = second.Next
	}

	return retNode
}