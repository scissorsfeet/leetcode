package LinkedList

/*
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。
*/

func reverseBetween(head *ListNode, m int, n int) *ListNode {
    helper := &ListNode{Next:head}
    pre := helper
    cur := helper.Next

    //找到起始位置
    for i:=1;i<m;i++ {
        pre = pre.Next
        cur = cur.Next
    }

    //m前的节点、m节点
    preM := pre
    mNode := cur

    //翻转中间链表
    nextM := cur.Next
    for i:=0;i<n-m&&nil!=nextM;i++{
        temp := nextM.Next
        nextM.Next = cur
        cur = nextM
        nextM = temp
    }

    //头尾对调
    mNode.Next = nextM
    preM.Next = cur

    return helper.Next
}