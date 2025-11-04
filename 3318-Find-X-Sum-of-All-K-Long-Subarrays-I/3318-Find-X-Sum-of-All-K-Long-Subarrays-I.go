package leetcode

import "container/heap"

type pair struct {
	val  int
	freq int
}

type pairList []pair

func (this pairList) Less(i, j int) bool {
	if this[i].freq == this[j].freq {
		return this[i].val > this[j].val
	}
	return this[i].freq > this[j].freq
}

func (this pairList) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this pairList) Len() int {
	return len(this)
}

func (this *pairList) Push(v any) {
	*this = append(*this, v.(pair))
}

func (this *pairList) Pop() any {
	old := *this
	n := len(old)
	x := old[n-1]
	*this = old[:n-1]
	return x
}

func findXSum(nums []int, k int, x int) []int {
	ans := make([]int, 0)

	hmap := make(map[int]int)
	for i := range k {
		hmap[nums[i]]++
	}
	ans = append(ans, topXSum(hmap, x))

	left := 0
	for right := k; right < len(nums); right++ {
		hmap[nums[left]]--
		if hmap[nums[left]] == 0 {
			delete(hmap, nums[left])
		}
		left++

		hmap[nums[right]]++
		ans = append(ans, topXSum(hmap, x))
	}

	return ans
}

// can use a helper function
func topXSum(hmap map[int]int, x int) int {
	pairHeap := &pairList{}
	heap.Init(pairHeap)

	for k, v := range hmap {
		heap.Push(pairHeap, pair{
			val:  k,
			freq: v,
		})
	}
	cur := 0
	topSum := 0
	for pairHeap.Len() > 0 && cur != x {
		top := heap.Pop(pairHeap).(pair)
		topSum += (top.val * top.freq)
		cur++
	}

	return topSum
}
