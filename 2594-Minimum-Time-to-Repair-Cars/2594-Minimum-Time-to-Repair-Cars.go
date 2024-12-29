package leetcode

import (
	"math"
	"slices"
)

func repairCars(ranks []int, cars int) int64 {
	min := slices.Min(ranks)

	l, r := 1, (min * cars * cars)

	ans := r
	for l <= r {
		m := (l + r) / 2

		count := 0
		for _, rank := range ranks {
			rankCount := int(math.Sqrt(float64(m) / float64(rank)))
			count += rankCount
		}

		if count >= cars {
			ans = m
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return int64(ans)
}
