package LinkedList

/*
Given a singly linked list, group all odd nodes together followed by the even nodes. Please note here we are talking about the node number and not the value in the nodes.

You should try to do it in place. The program should run in O(1) space complexity and O(nodes) time complexity.

Example:
Given 1->2->3->4->5->NULL,
return 1->3->5->2->4->NULL.
*/

func oddEvenList(head *ListNode) *ListNode {
    if nil == head || nil == head.Next || nil == head.Next.Next {
        return head
    }

    insertPosNode := head
    pre := head.Next
    cur := head.Next.Next

    for {
        //将cur从链表删除
        pre.Next = cur.Next

        //将cur插入到insertPos后
        tmp := insertPosNode.Next
        insertPosNode.Next = cur
        cur.Next = tmp

        //insertPos节点前进
        insertPosNode = insertPosNode.Next

        //pre走一步，cur走两步
        if nil == pre.Next || nil == pre.Next.Next {
            break
        }
        pre = pre.Next
        cur = pre.Next
    }

    return head
}