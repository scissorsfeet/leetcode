package Array

import (
	"container/heap"
)

func merge(nums1 []int, m int, nums2 []int, n int)  {
	if m == 0 {
		for i:=0;i<n;i++ {
			nums1[i] = nums2[i]
		}
		return
	}
	if n == 0 {
		return
	}

	h := IntHeap(nums2)
	heap.Init(&h)
	var i = 0
	for i < m {
		hMin := heap.Pop(&h).(int)
		for i < m && nums1[i] <= hMin {
			i++
		}
		if i >= m {
			heap.Push(&h, hMin)
			break
		}
		tmp := nums1[i]
		nums1[i] = hMin
		heap.Push(&h, tmp)
		i++
	}
	for j:=0;j<n;j++ {
		x := heap.Pop(&h).(int)
		nums1[i] = x
		i++
	}
	return
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
