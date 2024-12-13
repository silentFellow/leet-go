package leetcode

import (
	"container/heap"
)

type score struct {
	val int
	i   int
}

type scores []score

func (s scores) Len() int { return len(s) }

func (s scores) Less(i, j int) bool {
	if s[i].val == s[j].val {
		return s[i].i < s[j].i
	}
	return s[i].val < s[j].val
}

func (s scores) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s *scores) Push(val any) {
	*s = append(*s, val.(score))
}

func (s *scores) Pop() any {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[:n-1]
	return x
}

func findScore(nums []int) int64 {
	sh := &scores{}
	heap.Init(sh)

	for i, val := range nums {
		s := score{
			val: val,
			i:   i,
		}

		heap.Push(sh, s)
	}

	ans := 0
	for range len(nums) {
		top := heap.Pop(sh).(score)
		topIndex := top.i
		if nums[topIndex] == -1 {
			continue
		}

		ans += top.val
		if topIndex > 0 {
			nums[topIndex-1] = -1
		}

		if topIndex < len(nums)-1 {
			nums[topIndex+1] = -1
		}
	}

	return int64(ans)
}
