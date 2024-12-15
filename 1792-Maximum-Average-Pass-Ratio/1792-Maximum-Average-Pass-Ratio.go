package leetcode

import (
	"container/heap"
)

// heap type
type class [][]int

// sort interface
func (c class) Len() int { return len(c) }

func (c class) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

func (c class) Less(i, j int) bool {
  currenti := float64(c[i][0]) / float64(c[i][1])
  potentiali := float64(c[i][0] + 1) / float64(c[i][1] + 1)

  currentj := float64(c[j][0]) / float64(c[j][1])
  potentialj := float64(c[j][0] + 1) / float64(c[j][1] + 1)

	return (potentiali - currenti) > (potentialj - currentj)
}

// heap interface
func (c *class) Push(val any) {
	*c = append(*c, val.([]int))
}

func (c *class) Pop() any {
	old := *c
	n := len(old)
	x := old[n-1]
	*c = old[:n-1]
	return x
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	scores := &class{}
	heap.Init(scores)

	for _, val := range classes {
		heap.Push(scores, val)
	}

	for range extraStudents {
		top := heap.Pop(scores).([]int)
		top[0]++
		top[1]++
		heap.Push(scores, top)
	}

	var ans float64
	for scores.Len() > 0 {
		top := heap.Pop(scores).([]int)
		ans += (float64(top[0]) / float64(top[1]))
	}

	return ans / float64(len(classes))
}
