package LinkedList

func middleNode(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}
	if nil == head.Next {
		return head
	}
	helper := &ListNode{Next:head}
	slow := helper.Next
	fast := helper.Next
	for nil != fast && nil != fast.Next {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}