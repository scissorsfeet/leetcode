package LinkedList

/*
给定一个排序链表，删除所有重复的元素使得每个元素只留下一个。

案例：

给定 1->1->2，返回 1->2

给定 1->1->2->3->3，返回 1->2->3
*/

func deleteDuplicates(head *ListNode) *ListNode {
    if nil != head && nil != head.Next {
        //每一个新值的初始指针
        initP := head
        //当前指针
        cur := head.Next
        for nil != cur {
            //一旦出现不相等的元素，直接删除初始指针后面的重复元素
            if initP.Val != cur.Val {
                initP.Next = cur
                initP = cur
            }
            //当前指针前进
            cur = cur.Next
        }
        //将最后的初始指针的后继置空
        initP.Next = nil
    }

    return head
}