package leetcode

// TODO: verify once after sometime

import (
	"container/heap"
	"strings"
)

type heapImp []string

func (h heapImp) Less(i, j int) bool {
	if len(h[i]) != len(h[j]) {
		return len(h[i]) < len(h[j])
	} else {
    lex := strings.Compare(h[i], h[j])
    if lex < 0 { return true }
    return false
	}
}

func (h heapImp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h heapImp) Len() int {
	return len(h)
}

func (h *heapImp) Push(val any) {
	*h = append(*h, val.(string))
}

func (h *heapImp) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthLargestNumber(nums []string, k int) string {
	heapNums := &heapImp{}
	heap.Init(heapNums)

	for _, val := range nums {
		heap.Push(heapNums, val)

		if heapNums.Len() > k {
			heap.Pop(heapNums)
		}
	}

	return (*heapNums)[0]
}
