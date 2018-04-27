package LinkedList

func splitListToParts(root *ListNode, k int) []*ListNode {
    ret := make([]*ListNode, k)

    if nil != root {
        slotNumArr := make([]int, k)
        cur := root
        var pos = 0
        for nil != cur {
            slotNumArr[pos]++
            pos++
            pos = pos % k
            cur = cur.Next
        }

        cur = root
        for i, num := range slotNumArr {
            if num == 0 {
                ret[i] = nil
            } else {
                groupHead := cur
                for j := 1; j<num; j++ {
                    cur = cur.Next
                }
                tmp := cur.Next
                cur.Next = nil
                cur = tmp
                ret[i] = groupHead
            }
        }
    } else {
        for i:=0;i<k;i++{
            ret[i] = nil
        }
    }

    return ret
}