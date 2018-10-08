package LinkedList

type Node struct {
	val int
	next *Node
	prev *Node
	child *Node
}

func flatten(head *Node) *Node {
	cur := head
	for nil != cur {
		if nil != cur.child {
			oldNext := cur.next
			cur.next = cur.child
			cur.child.prev = cur
			cur.child = nil
			if nil != oldNext {
				tmp := cur.next
				for nil != tmp.next {
					tmp = tmp.next
				}
				tmp.next = oldNext
				oldNext.prev = tmp
			}
		}
		cur = cur.next
	}
	return head
}