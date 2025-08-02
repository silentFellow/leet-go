package leetcode

import "sort"

func minCost(basket1 []int, basket2 []int) int64 {
	minVal := basket1[0]

	h1, h2 := make(map[int]int), make(map[int]int)
	for _, v := range basket1 {
		h1[v]++
		minVal = min(minVal, v)
	}
	for _, v := range basket2 {
		h2[v]++
		minVal = min(minVal, v)
	}

	swapNeededArr1, swapNeededArr2 := make([]int, 0), make([]int, 0)
	for k, v := range h1 {
		sv := h2[k]
		total := v + sv
		if total%2 == 1 {
			return -1
		}

		if v > sv {
			times := (v - sv) / 2
			for range times {
				swapNeededArr1 = append(swapNeededArr1, k)
			}
		}
	}
	for k, v := range h2 {
		sv := h1[k]
		total := v + sv
		if total%2 == 1 {
			return -1
		}

		if v > sv {
			times := (v - sv) / 2
			for range times {
				swapNeededArr2 = append(swapNeededArr2, k)
			}
		}
	}

	if len(swapNeededArr1) != len(swapNeededArr2) {
		return -1
	}

	sort.Ints(swapNeededArr1)
	sort.Sort(sort.Reverse(sort.IntSlice(swapNeededArr2)))

	ans := 0
	for i := 0; i < len(swapNeededArr1); i++ {
		normalSwap := min(swapNeededArr1[i], swapNeededArr2[i])
		doubleSwap := 2 * minVal
		ans += min(normalSwap, doubleSwap)
	}

	return int64(ans)
}
