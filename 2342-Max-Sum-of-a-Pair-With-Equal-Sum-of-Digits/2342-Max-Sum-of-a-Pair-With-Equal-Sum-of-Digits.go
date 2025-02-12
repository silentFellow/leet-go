package leetcode

import "container/heap"

func maximumSum(nums []int) int {
	hmap := make(map[int]*HeapStruct)

	ans := -1
	for _, val := range nums {
		sum := sumOfDigit(val)
		if _, ok := hmap[sum]; ok {
			maxv := (*hmap[sum])[0]
			ans = max(ans, val+maxv)
		}

		if _, ok := hmap[sum]; !ok {
			hmap[sum] = &HeapStruct{}
			heap.Init(hmap[sum])
		}
		heap.Push(hmap[sum], val)
	}

	return ans
}

// heap structure
type HeapStruct []int

// sort interfece
func (h HeapStruct) Len() int { return len(h) }

func (h HeapStruct) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h HeapStruct) Less(i, j int) bool { return h[i] > h[j] }

// heap interfece
func (h *HeapStruct) Push(v any) {
	*h = append(*h, v.(int))
}

func (h *HeapStruct) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// calculate sum of digits
func sumOfDigit(v int) int {
	ans := 0

	for v != 0 {
		ans += v % 10
		v /= 10
	}

	return ans
}
