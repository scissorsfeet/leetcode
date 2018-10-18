package LinkedList

/*
给出一个链表，每 k 个节点一组进行翻转，并返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么将最后剩余节点保持原有顺序。

示例 :

给定这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5

说明 :

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	if nil == head{
		return head
	}
	helper := &ListNode{Next:head}
	head = helper
	for nil != head && nil != head.Next {
		slow := head.Next
		fast := slow
		noTraverse := false
		for i:=0;i<k-1;i++ {
			if nil == fast.Next {
				noTraverse = true
				break
			}
			fast = fast.Next
		}
		if !noTraverse {
			tail := fast.Next
			traverse(slow, fast)
			head.Next = fast
			slow.Next = tail
			head = slow
		} else {
			break
		}
	}
	return helper.Next
}

//翻转中间部分
//head:该区间内的起始节点
//tail:该区间内的终止节点
func traverse(head, tail *ListNode) {
	nextZoneHead := tail.Next
	var pre *ListNode = nil
	cur := head
	for nil != cur && cur != nextZoneHead {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
}