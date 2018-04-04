package LinkedList
//Write a function to delete a node (except the tail) in a singly linked list, given only access to that node.

//Supposed the linked list is 1 -> 2 -> 3 -> 4 and you are given the third node with value 3,
//the linked list should become 1 -> 2 -> 4 after calling your function.


/*
解法：直接将该节点的next节点的值和后继指针覆盖待删除节点

疑惑：简单的int类型或者指针表高效
*/
func DeleteNodeInALinkedList(node *ListNode) {
    if nil == node.Next {
        return
    }
    node.Val = node.Next.Val
    node.Next = node.Next.Next
}