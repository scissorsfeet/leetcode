package LinkedList

/*
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。
*/

func partition(head *ListNode, x int) *ListNode {
    //辅助头结点,持有整个链表
    helper := &ListNode{Next:head}

    //定位第一个大于等于x的节点
    var preM, m *ListNode
    pre := helper
    cur := head
    for nil != cur {
        if x <= cur.Val {
            preM = pre
            m = cur
            break
        }
        pre = cur
        cur = cur.Next
    }

    //如果没有这个节点，直接返回
    if nil != m {
        //将小于x的值前插
        for nil != cur {
            next := cur.Next
            if x > cur.Val {
                //从原位置删除改节点
                pre.Next = next
                //插入m节点前
                preM.Next = cur
                cur.Next = m
                //更新preM
                preM = cur
            } else { //没有删除节点，需要更新pre
                pre = cur
            }
            //更新cur
            cur = next
        }
    }

    return helper.Next
}