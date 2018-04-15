package LinkedList

/*
给定一个链表，判断链表中否有环。
*/

func hasCycle(head *ListNode) bool {
	var result bool

	if nil != head && nil != head.Next {
		slow := head
		fast := head.Next

		for nil != fast && nil != fast.Next {
			if slow == fast {
				result = true
				break
			}
			slow = slow.Next
			fast = fast.Next.Next
		}
	}

	return result
}
