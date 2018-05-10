package LinkedList

/**
给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	helper := &ListNode{}

	p := l1
	q := l2

	pre := helper
	carry := 0//进位字段

	//先把两个链表交集部分相加
	for nil != p && nil != q {
		pre.Next, carry = getCarryAndNode(p.Val+q.Val+carry)
		pre = pre.Next
		p = p.Next
		q = q.Next
	}

	//获取较长链表的剩余部分
	var leftNode *ListNode
	if nil != p {
		leftNode = p
	} else if nil != q {
		leftNode = q
	}

	//将剩余部分与进位值相加
	for nil != leftNode {
		pre.Next, carry = getCarryAndNode(leftNode.Val+carry)
		pre = pre.Next
		leftNode = leftNode.Next
	}

	//有可能最后还有一个进位，需要收尾
	if carry > 0 {
		pre.Next = &ListNode{Val:1}
	}

	return helper.Next
}

func getCarryAndNode(val int) (*ListNode, int) {
	carry := val / 10
	node := &ListNode{Val:val%10}
	return node, carry
}