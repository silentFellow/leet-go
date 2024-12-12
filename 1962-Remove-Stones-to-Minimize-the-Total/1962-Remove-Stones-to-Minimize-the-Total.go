package leetcode

import (
	"container/heap"
	"math"
)

type stones []int

// sort interface
func (s stones) Len() int { return len(s) }
func (s stones) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s stones) Less(i, j int) bool { return s[i] > s[j] }

// heap interface
func (s *stones) Push(val any) {
  *s = append(*s, val.(int))
}

func (s *stones) Pop() any {
  old := *s
  n := len(old)
  x := old[n-1]
  *s = old[:n-1]
  return x
}


func minStoneSum(piles []int, k int) int {
  sh := &stones{}

  heap.Init(sh)
  for _, val := range piles {
    heap.Push(sh, val)
  }

  for range k {
    val := heap.Pop(sh).(int)
    neoVal := int(math.Ceil((float64(val) / 2)))
    heap.Push(sh, neoVal)
  }

  sum := 0
  for sh.Len() > 0 {
    sum += heap.Pop(sh).(int)
  }

  return sum
}
