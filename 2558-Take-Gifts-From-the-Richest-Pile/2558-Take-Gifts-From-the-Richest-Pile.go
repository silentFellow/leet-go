package leetcode

import (
	"container/heap"
	"math"
)

type giftHeap []int

// sort interface
func (g giftHeap) Swap(i, j int) { g[i], g[j] = g[j], g[i] }
func (g giftHeap) Less(i, j int) bool { return g[i] > g[j] }
func (g giftHeap) Len() int { return len(g) }

// heap interface
func (g *giftHeap) Push(i any) { *g = append(*g, i.(int)) }
func (g *giftHeap) Pop() any {
  old := *g
  n := len(old)
  x := old[n-1]
  *g = old[:n-1]
  return x
}

func pickGifts(gifts []int, k int) int64 {
  gh := &giftHeap{}

  heap.Init(gh)
  for _, val := range gifts {
    heap.Push(gh, val)
  }

  for range k {
    val := heap.Pop(gh).(int)
    neoVal := int(math.Sqrt(float64(val)))
    heap.Push(gh, neoVal)
  }

  sum := 0
  for gh.Len() > 0 {
    sum += heap.Pop(gh).(int)
  }

  return int64(sum)
}
