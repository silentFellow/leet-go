package leetcode

import (
	"container/heap"
	"sort"
)

type heapType [][]int

func (h heapType) Len() int { return len(h) }

func (h heapType) Less(i, j int) bool { return h[i][1] < h[j][1] }

func (h heapType) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *heapType) Push(v any) {
	*h = append(*h, v.([]int))
}

func (h *heapType) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	minHeap := &heapType{}
	heap.Init(minHeap)

	lastDay := 0
	for _, v := range events {
		lastDay = max(lastDay, v[1])
	}

	count := 0
	idx := 0
	for day := 1; day <= lastDay; day++ {
		for idx < len(events) && events[idx][0] == day {
			heap.Push(minHeap, events[idx])
			idx++
		}

		for minHeap.Len() > 0 && (*minHeap)[0][1] < day {
			heap.Pop(minHeap)
		}

		if minHeap.Len() > 0 {
			heap.Pop(minHeap)
			count++
		}
	}

	return count
}
