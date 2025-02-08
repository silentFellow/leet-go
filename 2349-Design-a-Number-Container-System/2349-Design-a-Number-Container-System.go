package leetcode

import "container/heap"

type NumberContainers struct {
	indexMap map[int]int
	hmap     map[int]*MinHeap
}

func Constructor() NumberContainers {
	return NumberContainers{
		indexMap: map[int]int{},
		hmap:     map[int]*MinHeap{},
	}
}

func (this *NumberContainers) Change(index int, number int) {
	// Remove the old number from the map
	if old, exists := this.indexMap[index]; exists {
		if minHeap, ok := this.hmap[old]; ok {
			minHeap.Remove(index)
			if minHeap.Len() == 0 {
				delete(this.hmap, old)
			}
		}
	}

	// Update the indexMap with the new number
	this.indexMap[index] = number

	// Add the new index to the heap
	if _, exists := this.hmap[number]; !exists {
		this.hmap[number] = &MinHeap{}
		heap.Init(this.hmap[number])
	}
	heap.Push(this.hmap[number], index)
}

func (this *NumberContainers) Find(number int) int {
	if minHeap, exists := this.hmap[number]; exists && minHeap.Len() > 0 {
		return (*minHeap)[0]
	}
	return -1
}

// MinHeap is a min-heap of integers
type MinHeap []int

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap) Remove(index int) {
	for i, v := range *h {
		if v == index {
			heap.Remove(h, i)
			break
		}
	}
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */
