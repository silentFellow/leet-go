package leetcode

import (
	"math"
	"slices"
)

func minDays(bloomDay []int, m int, k int) int {
	l, r := 1, slices.Max(bloomDay)

	ans := math.MaxInt64
	for l <= r {
		mid := (l + r) / 2

		count := 0
		continuous := 0
		for _, val := range bloomDay {
			if val-mid <= 0 {
				continuous++
			}

			if continuous == k {
				count++
				continuous = 0
			}

			if val-mid > 0 {
				continuous = 0
			}
		}

		if count >= m && ans >= mid {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	if ans == math.MaxInt64 {
		return -1
	}
	return ans
}
