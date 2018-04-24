package LinkedList

func getIntersectionNode(headA *ListNode, headB *ListNode) *ListNode {
	var intersectionNode *ListNode = nil
	if nil != headA && nil != headB {
		//首先找到A链的tail节点
		tailA := headA
		for nil != tailA.Next {
			tailA = tailA.Next
		}

		//把tailA与headB连接形成环
		tailA.Next = headB

		//快慢指针相遇
		var slow = headA
		var fast = headA
		var hasCycle bool = false
		for nil != fast && nil != fast.Next {
			slow = slow.Next
			fast = fast.Next.Next
			if slow == fast {
				hasCycle = true
				break
			}
		}


		//如果有环
		//慢指针回到headA
		//slow和fast同时前进，再相遇即是交点
		if hasCycle {
			slow = headA
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			intersectionNode = slow
		}

		//tailA与headB断链，恢复原始结构
		tailA.Next = nil
	}

	return intersectionNode
}