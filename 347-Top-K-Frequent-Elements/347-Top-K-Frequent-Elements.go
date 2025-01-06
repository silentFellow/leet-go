package leetcode

import "container/heap"

type pairs struct {
	val  int
	freq int
}

type maxHeap []pairs

func (m maxHeap) Len() int { return len(m) }

func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m maxHeap) Less(i, j int) bool {
	return m[i].freq > m[j].freq
}

func (m *maxHeap) Push(v any) {
	*m = append(*m, v.(pairs))
}

func (m *maxHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	hmap := make(map[int]int)
	for _, val := range nums {
		hmap[val]++
	}

	mh := &maxHeap{}
	heap.Init(mh)
	for key, val := range hmap {
		pair := pairs{
			val:  key,
			freq: val,
		}

		heap.Push(mh, pair)
	}

	ans := []int{}
	for range k {
		top := heap.Pop(mh).(pairs)
		ans = append(ans, top.val)
	}

	return ans
}
