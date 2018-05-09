package LinkedList

/*
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

说明：不允许修改给定的链表。
*/

func detectCycle(head *ListNode) *ListNode {
	helper := &ListNode{Next:head}
	//快慢指针找相遇节点
	slow := helper
	fast := helper

	for nil != fast && nil != fast.Next {
		if slow == fast && slow != helper {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	var ret *ListNode

	if nil != fast && nil != fast.Next {
		slow = helper
		for slow != fast {
			slow = slow.Next
			fast = fast.Next
		}
		ret = slow
	}

	return ret
}