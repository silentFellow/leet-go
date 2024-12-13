package leetcode

// TODO: optimize it with a hamp => lazy remove stale statergy

import (
	"container/heap"
)

type maxHeap []int

// sort interface
func (m maxHeap) Less(i, j int) bool { return m[i] > m[j] }
func (m maxHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m maxHeap) Len() int { return len(m) }

// heap interface
func (m *maxHeap) Push(val any) {
  *m = append(*m, val.(int))
}

func (m *maxHeap) Pop() any {
  old := *m
  n := len(old)
  x := old[n-1]
  *m = old[:n-1]
  return x
}

func maxSlidingWindow(nums []int, k int) []int {
  ans := []int{}

  sh := &maxHeap{}
  heap.Init(sh)

  // initial
  for i := range k {
    heap.Push(sh, nums[i])
  }
  ans = append(ans, (*sh)[0])

  for i:=k; i<len(nums); i++ {
    // remove val
    val := nums[i-k]
    for j:=0; j<sh.Len(); j++ {
      if (*sh)[j] == val {
        heap.Remove(sh, j)
        break
      }
    }

    heap.Push(sh, nums[i])
    ans = append(ans, (*sh)[0])
  }

  return ans
}

// package leetcode
//
// import (
// 	"container/heap"
// )
//
// // maxHeap defines a max-heap
// type maxHeap struct {
// 	data  []int
// 	index map[int]int // Tracks the index of values
// }
//
// // Implement sort.Interface
// func (m maxHeap) Len() int           { return len(m.data) }
// func (m maxHeap) Less(i, j int) bool { return m.data[i] > m.data[j] } // Max heap
// func (m maxHeap) Swap(i, j int) {
// 	m.data[i], m.data[j] = m.data[j], m.data[i]
// 	m.index[m.data[i]] = i
// 	m.index[m.data[j]] = j
// }
//
// // Implement heap.Interface
// func (m *maxHeap) Push(val any) {
// 	v := val.(int)
// 	m.index[v] = len(m.data)
// 	m.data = append(m.data, v)
// }
//
// func (m *maxHeap) Pop() any {
// 	old := m.data
// 	n := len(old)
// 	val := old[n-1]
// 	delete(m.index, val)
// 	m.data = old[:n-1]
// 	return val
// }
//
// // RemoveStale removes the top element if it is no longer valid
// func (m *maxHeap) RemoveStale(stale map[int]int) {
// 	for m.Len() > 0 {
// 		top := m.data[0]
// 		if count, ok := stale[top]; ok && count > 0 {
// 			stale[top]--
// 			heap.Pop(m)
// 		} else {
// 			break
// 		}
// 	}
// }
//
// // maxSlidingWindow calculates the maximum of each sliding window
// func maxSlidingWindow(nums []int, k int) []int {
// 	if len(nums) == 0 || k == 0 {
// 		return []int{}
// 	}
//
// 	result := []int{}
// 	stale := map[int]int{} // Tracks elements that need lazy removal
// 	h := &maxHeap{data: []int{}, index: map[int]int{}}
// 	heap.Init(h)
//
// 	// Initialize the first window
// 	for i := 0; i < k; i++ {
// 		heap.Push(h, nums[i])
// 	}
// 	h.RemoveStale(stale) // Ensure the heap is clean
// 	result = append(result, h.data[0])
//
// 	// Process the sliding window
// 	for i := k; i < len(nums); i++ {
// 		outgoing := nums[i-k] // Element leaving the window
// 		incoming := nums[i]  // Element entering the window
//
// 		// Mark outgoing element as stale
// 		stale[outgoing]++
//
// 		// Add the new element
// 		heap.Push(h, incoming)
//
// 		// Remove any stale elements at the top
// 		h.RemoveStale(stale)
//
// 		// Append the current max
// 		result = append(result, h.data[0])
// 	}
//
// 	return result
// }
