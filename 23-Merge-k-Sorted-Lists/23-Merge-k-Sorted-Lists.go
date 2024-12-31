package leetcode

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

// heap
type ListNodeHeap []*ListNode

// sort interface
func (l ListNodeHeap) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l ListNodeHeap) Less(i, j int) bool {
	return l[i].Val < l[j].Val
}

func (l ListNodeHeap) Len() int {
	return len(l)
}

// heap interface
func (l *ListNodeHeap) Push(v any) {
	*l = append(*l, v.(*ListNode))
}

func (l *ListNodeHeap) Pop() any {
	old := *l
	n := len(old)
	x := old[n-1]
	*l = old[:n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	lh := &ListNodeHeap{}
	heap.Init(lh)

	for _, node := range lists {
		if node != nil {
			heap.Push(lh, node)
		}
	}

	dummy := &ListNode{}
	dummyTail := dummy

	for lh.Len() > 0 {
		top := heap.Pop(lh).(*ListNode)

		dummyTail.Next = top
		dummyTail = top

		if top.Next != nil {
			heap.Push(lh, top.Next)
		}
	}

	return dummy.Next
}
