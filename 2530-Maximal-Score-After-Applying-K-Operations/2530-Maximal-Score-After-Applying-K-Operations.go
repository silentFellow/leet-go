package leetcode

import (
	"container/heap"
	"math"
)

type maxElement []int

// sort interface
func (m maxElement) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m maxElement) Less(i, j int) bool { return m[i] > m[j] }
func (m maxElement) Len() int { return len(m) }

// heap interface
func (m *maxElement) Push(val any) {
  *m = append(*m, val.(int))
}

func (m *maxElement) Pop() any {
  old := *m
  n := len(old)
  x := old[n-1]
  *m = old[:n-1]
  return x
}

func maxKelements(nums []int, k int) int64 {
  nh := &maxElement{}

  heap.Init(nh)
  for _, val := range nums {
    heap.Push(nh, val)
  }

  sum := 0
  for range k {
    val := heap.Pop(nh).(int)
    sum += val
    neoVal := int(math.Ceil(float64(val) / 3))
    heap.Push(nh, neoVal)
  }

  return int64(sum)
}
