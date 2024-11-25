package leetcode

import "container/heap"

type heapImp []int

// sort interface
func (h heapImp) Less (i, j int) bool {
  return h[i] < h[j]
}
func (h heapImp) Swap (i, j int) {
  h[i], h[j] = h[j], h[i]
}
func (h heapImp) Len () int { return len(h) }

// heap interface
func (h *heapImp) Push (val interface{}) {
  *h = append(*h, val.(int))
}

func (h *heapImp) Pop () interface{} {
  old := *h
  n := len(old)
  x := old[n-1]
  *h = old[0:n-1]
  return x
}

func findKthLargest(nums []int, k int) int {
  heapNums := &heapImp{}
  heap.Init(heapNums)

  for _, val := range nums {
    heap.Push(heapNums, val)
    
    // just maintain only k elements in heap rather than adding all and pushing it
    if heapNums.Len() > k {
      heap.Pop(heapNums)
    }
  }

  return (*heapNums)[0]
}
