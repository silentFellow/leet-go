package leetcode

// imports
import "container/heap"

// type to store both index and value
type state struct {
	val   int
	index int
}

// heap type
type final []state

// sort interface
func (f final) Len() int { return len(f) }

func (f final) Swap(i, j int) { f[i], f[j] = f[j], f[i] }

func (f final) Less(i, j int) bool {
	if f[i].val == f[j].val {
		return f[i].index < f[j].index
	}
	return f[i].val < f[j].val
}

// heap interface
func (f *final) Push(val any) { *f = append(*f, val.(state)) }

func (f *final) Pop() any {
	old := *f
	n := len(old)
	x := old[n-1]
	*f = old[:n-1]
	return x
}

func getFinalState(nums []int, k int, multiplier int) []int {
	f := &final{}

	heap.Init(f)
	for i, val := range nums {
		heap.Push(f, state{val, i})
	}

	for range k {
		top := heap.Pop(f).(state)
		top.val *= multiplier
		heap.Push(f, top)
	}

	for f.Len() > 0 {
		top := heap.Pop(f).(state)
		nums[top.index] = top.val
	}

	return nums
}
