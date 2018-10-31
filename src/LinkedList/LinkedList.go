package LinkedList
import (
    "fmt"
    "strings"
)

//链表节点
type ListNode struct {
    Val int
    Next *ListNode
}

//创建节点
func NewListNode(value int, next *ListNode) *ListNode {
    return &ListNode{value, next}
}

//链表长度
func ListLen(head *ListNode) int {
    var length = 0
    for nil != head {
        length++
        head = head.Next
    }
    return length
}

type LinkedList struct {
    Head *ListNode
}

func NewLinkedList(head *ListNode) *LinkedList {
    return &LinkedList{head}
}

//通过数据创建链表
func CreateListByArray(items []int) *LinkedList {
    if len(items) <= 0 {
        return &LinkedList{Head:&ListNode{}}
    }

    head := NewListNode(0, nil)
    cur := head
    for _, v := range items {
        newNode := NewListNode(v, nil)
        cur.Next = newNode
        cur = cur.Next
    }

    return NewLinkedList(head)
}

func(this *LinkedList) Print() {
    items := make([]string, 0, 8)
    cur := this.Head.Next
    for nil != cur {
        items = append(items, fmt.Sprintf("%+v",cur.Val))
        cur = cur.Next
    }
    fmt.Println(strings.Join(items, "->"))
}