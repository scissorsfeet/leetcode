package LinkedList

/*
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:

输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:

输入: 1->1->1->2->3
输出: 2->3
*/

func deleteDuplicatesII(head *ListNode) *ListNode {
    helper := &ListNode{Next:head}

    pre := helper
    cur := helper.Next

    var repeatTimes = 0
    for nil != cur && nil != cur.Next {
        if cur.Val != cur.Next.Val {//值不等
            if repeatTimes > 0 {
                pre.Next = cur.Next
                repeatTimes = 0
            } else {
                pre = cur
            }
        } else {
            repeatTimes++
        }
        cur = cur.Next
    }

    //收个尾
    if nil != cur {
        if repeatTimes > 0 {
            pre.Next = nil
        } else {
            pre.Next = cur
        }
    }

    return helper.Next
}