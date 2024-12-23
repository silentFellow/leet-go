package leetcode

import (
	"container/heap"
	"sort"
)

type MaxQuery struct {
	value int
	index int
	queryIndex int
}

type MaxQueryHeap []MaxQuery

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	results := make([]int, len(queries))
	unansweredQueries := make([]MaxQuery, 0)

	// pre-process
	// work through direct values
	for i, query := range queries {
		l, r := query[0], query[1]
		if l > r {
			l, r = r, l
		}

		if l == r || heights[r] > heights[l] {
			results[i] = r
		} else {
			results[i] = -1
			unansweredQueries = append(unansweredQueries, MaxQuery{heights[l]+1, r, i})
		}
	}

	sort.Slice(unansweredQueries, func(i, j int) bool {
		return unansweredQueries[i].index < unansweredQueries[j].index
	})

	mqh := &MaxQueryHeap{}
	heap.Init(mqh)
	for i, currentBuilding := range heights {
		for len(unansweredQueries) > 0 && unansweredQueries[0].index < i {
			heap.Push(mqh, unansweredQueries[0])
			unansweredQueries = unansweredQueries[1:]
		}

		for mqh.Len() > 0 && (*mqh)[0].value <= currentBuilding {
			top := heap.Pop(mqh).(MaxQuery)
			results[top.queryIndex] = i
		}
	}

	return results
}

func (mqh MaxQueryHeap) Len() int           { return len(mqh) }
func (mqh MaxQueryHeap) Swap(i, j int)      { mqh[i], mqh[j] = mqh[j], mqh[i] }
func (mqh MaxQueryHeap) Less(i, j int) bool { return mqh[i].value < mqh[j].value }

func (mqh *MaxQueryHeap) Push(v any) {
	*mqh = append(*mqh, v.(MaxQuery))
}

func (mqh *MaxQueryHeap) Pop() any {
	old := *mqh
	n := len(old)
	x := old[n-1]
	*mqh = old[:n-1]
	return x
}
