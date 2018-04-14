package LinkedList

/*
删除链表中等于给定值 val 的所有元素。

示例
给定: 1 --> 2 --> 6 --> 3 --> 4 --> 5 --> 6, val = 6
返回: 1 --> 2 --> 3 --> 4 --> 5
*/
func removeElements(head *ListNode, val int) *ListNode {
    if nil != head {

        //先把头部的值吃掉
        for nil != head && val == head.Val {
            head = head.Next
        }

        if nil != head {
            //保证头指针已经不是要删除的元素了
            pre := head
            cur := head.Next
            for nil != cur {
                if val == cur.Val {//删除该节点
                    pre.Next = cur.Next
                    cur = cur.Next
                } else { //不需要删除时持续前进
                    pre = cur
                    cur = pre.Next
                }
            }
        }
    }
    return head
}