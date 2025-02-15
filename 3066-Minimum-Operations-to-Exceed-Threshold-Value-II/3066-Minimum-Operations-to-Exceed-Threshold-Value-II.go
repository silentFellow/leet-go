package leetcode

import "container/heap"

func minOperations(nums []int, k int) int {
	ans := 0

	h := &HeapStruct{}
	heap.Init(h)
	for _, v := range nums {
		heap.Push(h, v)
	}

	for h.Len() >= 2 && (*h)[0] < k {
		ans++
		f := heap.Pop(h).(int)
		s := heap.Pop(h).(int)

		v := min(f, s)*2 + max(f, s)
		heap.Push(h, v)
	}

	return ans
}

// heap struct
type HeapStruct []int

// sort interface
func (h HeapStruct) Len() int { return len(h) }

func (h HeapStruct) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h HeapStruct) Less(i, j int) bool { return h[i] < h[j] }

// heap interface
func (h *HeapStruct) Push(v any) {
	*h = append(*h, v.(int))
}

func (h *HeapStruct) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
