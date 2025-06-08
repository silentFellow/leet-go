package leetcode

import (
	"container/heap"
	"strings"
)

type Pair struct {
	val rune
	idx int
}

type heapStruct []Pair

// sort interface
func (h heapStruct) Len() int {
	return len(h)
}

func (h heapStruct) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h heapStruct) Less(i, j int) bool {
	if h[i].val != h[j].val {
		return h[i].val < h[j].val
	}
	return h[i].idx > h[j].idx
}

// heap interface
func (h *heapStruct) Push(v any) {
	*h = append(*h, v.(Pair))
}

func (h *heapStruct) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func clearStars(s string) string {
	// heap
	minHeap := &heapStruct{}
	heap.Init(minHeap)

	removed := make(map[int]struct{})
	for i, v := range s {
		if v == '*' {
			for minHeap.Len() > 0 {
				if _, ok := removed[(*minHeap)[0].idx]; ok {
					heap.Pop(minHeap)
				} else {
					break
				}
			}

			if minHeap.Len() > 0 {
				removed[(*minHeap)[0].idx] = struct{}{}
			}

			removed[i] = struct{}{}
		} else {
			heap.Push(minHeap, Pair{val: v, idx: i})
		}
	}

	var ans strings.Builder
	for i, v := range s {
		if _, ok := removed[i]; !ok {
			ans.WriteRune(v)
		}
	}

	return ans.String()
}
