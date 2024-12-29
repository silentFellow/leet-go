package leetcode

import "slices"

func shipWithinDays(weights []int, days int) int {
	max := slices.Max(weights)
	sum := 0
	for _, weight := range weights {
		sum += weight
	}

	// best case sum => all pkg in one go
	// worst case all are max values
	l, r := max, sum

	ans := sum
	for l <= r {
		m := (l + r) / 2

		count, individual := 0, 0
		for _, weight := range weights {
			individual += weight
			// if exactly equals to m
			// inc count ans and reset individual
			if individual == m {
				count++
				individual = 0
			} else if individual > m { // if greater inc count ans assign current weight to individual
				count++
				individual = weight
			}
		}
		// cleanup
		// if last pkg still goint
		if individual != 0 {
			count++
		}

		// normal binary search
		if count <= days {
			ans = m
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return ans
}
