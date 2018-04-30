package LinkedList

/*
对链表进行插入排序。
*/

func InsertionSortList(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}

	//定义排序好的链表头和尾
	sortedHead := head
	sortedTail := head
	cur := head.Next

	for nil != cur {
		next := cur.Next
		if cur.Val >= sortedTail.Val {//如果当前节点的值大于排序链表尾节点，直接插入到尾部
			sortedTail.Next = cur
			sortedTail = cur
		} else if cur.Val <= sortedHead.Val {//如果当前节点的值小于排序链表头结点，直接插头
			cur.Next = sortedHead
			sortedHead = cur
		} else {//在头尾区间内需要遍历查找
			tmpPre := sortedHead
			tmpCur := sortedHead.Next
			for cur.Val > tmpCur.Val {
				tmpPre = tmpCur
				tmpCur = tmpCur.Next
			}
			tmpPre.Next = cur
			cur.Next = tmpCur
		}
		cur = next
	}

	//断尾
	sortedTail.Next = nil

	//取头
	return sortedHead
}