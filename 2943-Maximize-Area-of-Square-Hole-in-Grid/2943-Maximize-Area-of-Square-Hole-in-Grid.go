package leetcode

import (
	"math"
	"slices"
)

func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	maxSpan := func(bars []int) int {
		slices.Sort(bars)
		res, streak := 1, 1

		for i := 1; i < len(bars); i++ {
			if bars[i]-bars[i-1] == 1 {
				streak++
			} else {
				streak = 1
			}

			res = max(res, streak)
		}

		return res + 1
	}

	return int(math.Pow(float64(min(maxSpan(hBars), maxSpan(vBars))), 2))
}
