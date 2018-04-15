package LinkedList

/*
请检查一个链表是否为回文链表。

进阶：
你能在 O(n) 的时间和 O(1) 的额外空间中做到吗？
*/

func isPalindrome(head *ListNode) bool {
	var result bool = true

	if nil != head {
		//首先找到中间节点
		//使用快慢指针法
		slow := head
		fast := head
		var step int = 0
		for nil != fast && nil != fast.Next {
			slow = slow.Next
			fast = fast.Next.Next
			step++
		}

		//new一个栈，用来存放链表的后半段
		nodeStack := make([]*ListNode, 0, step)
		cur := slow
		for nil != cur {
			nodeStack = append(nodeStack, cur)
			cur = cur.Next
		}

		//对比
		stackLen := len(nodeStack)
		cur = head
		for i:=stackLen-1; i>=0; i-- {
			if cur.Val != nodeStack[i].Val {
				result = false
				break
			}
			cur = cur.Next
		}
	}

	return result
}