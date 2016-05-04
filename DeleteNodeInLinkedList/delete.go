package DeleteNodeInLinkedList

import (
	//"fmt"

	linkedlist "github.com/liutianyi1989/leetcode/LinkedList"
)

/**
 *Write a function to delete a node (except the tail) in a singly linked list, given only access to that node.
 *删除单链表中给定的节点(除了尾节点之外）
 *Supposed the linked list is 1 -> 2 -> 3 -> 4 and you are given the third node with value 3, the linked list should become 1 -> 2 -> 4 after calling your function.
 */
func solution(node *linkedlist.Node) {
	if node != nil {
		if node.Next != nil {
			*node = *node.Next
		}
	}
}
