package LinkedList

//import "fmt"

/*
We are given head, the head node of a linked list containing unique integer values.

We are also given the list G, a subset of the values in the linked list.

Return the number of connected components in G, where two values are connected if they appear consecutively in the linked list
*/

func numComponents(head *ListNode, G []int) int {
	m := make(map[int]bool, len(G))
	for _, v := range G {
		m[v] = true
	}

	var res = 0
	cur := head
	for nil != cur {
		//fmt.Println(cur.Val)
		if _, ok := m[cur.Val]; ok {
			res++
			cur = cur.Next
			for nil != cur {
				if _, ok := m[cur.Val]; ok {
					cur = cur.Next
				} else {
					break
				}
			}
		} else {
			cur = cur.Next
		}
	}

	return res
}
