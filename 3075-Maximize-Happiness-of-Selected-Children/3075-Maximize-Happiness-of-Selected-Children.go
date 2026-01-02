package leetcode

import "container/heap"

type heapType []int

// sort interface
func (h heapType) Len() int { return len(h) }

func (h heapType) Less(i, j int) bool { return h[i] > h[j] }

func (h heapType) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// heap interface
func (h *heapType) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *heapType) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// main method
func maximumHappinessSum(happiness []int, k int) int64 {
	heapData := &heapType{}
	heap.Init(heapData)

	for _, h := range happiness {
		heap.Push(heapData, h)
	}

	ans := 0
	round := 0
	for range k {
		top := heap.Pop(heapData).(int)
		actual := top - round
		if actual < 0 {
			break
		}
        ans += actual

		round++
	}

	return int64(ans)
}
