package LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//先把两个链表的值相加到一个新的链表中
	l1Len := getListLen(l1)
	l2Len := getListLen(l2)
	var long, short *ListNode
	var lenDiff int
	if l1Len > l2Len {
		long, short = l1, l2
		lenDiff = l1Len - l2Len
	} else {
		long, short = l2, l1
		lenDiff = l2Len - l1Len
	}
	//先把长链表多出来的部分拿出来
	head := new(ListNode)
	var pre *ListNode = head
	for i:=0;i<lenDiff;i++{
		newNode := &ListNode{Val:long.Val}
		pre.Next = newNode
		pre = newNode
		long = long.Next
	}
	//重合部分相加
	for nil != long && nil != short {
		newNode := &ListNode{Val:long.Val+short.Val}
		pre.Next = newNode
		pre = newNode
		long = long.Next
		short = short.Next
	}

	getCarry(head)

	if head.Val == 0 {
		return head.Next
	}
	return head
}

//获取链表长度
func getListLen(p *ListNode) int {
	var len = 0
	for nil != p {
		p = p.Next
		len++
	}
	return len
}

//获取进位值
func getCarry(p *ListNode) int {
	if nil == p {
		return 0
	}

	tmp := p.Val + getCarry(p.Next)
	p.Val = tmp % 10
	return tmp / 10
}