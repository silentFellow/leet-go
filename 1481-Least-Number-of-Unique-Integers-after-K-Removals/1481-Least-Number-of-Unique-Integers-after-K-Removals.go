package leetcode

import "container/heap"

type freqPairs struct {
	val, freq int
}

type heapPairs []freqPairs

// sort interface
func (s heapPairs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s heapPairs) Less(i, j int) bool { return s[i].freq < s[j].freq }

func (s heapPairs) Len() int { return len(s) }

// heap interface
func (s *heapPairs) Push(v any) { *s = append(*s, v.(freqPairs)) }

func (s *heapPairs) Pop() any {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[:n-1]
	return x
}

func findLeastNumOfUniqueInts(arr []int, k int) int {
	hmap := make(map[int]int)
	for _, v := range arr {
		hmap[v]++
	}

	heapArr := &heapPairs{}
	heap.Init(heapArr)
	for k, v := range hmap {
		pair := freqPairs{
			val:  k,
			freq: v,
		}
		heap.Push(heapArr, pair)
	}

	for range k {
		top := heap.Pop(heapArr).(freqPairs)
		if top.freq-1 != 0 {
			pair := freqPairs{
				val:  top.val,
				freq: top.freq-1,
			}
			heap.Push(heapArr, pair)
		}
	}

  return heapArr.Len()
}
