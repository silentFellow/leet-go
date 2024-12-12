package leetcode

import "sort"

func maximumBeauty(nums []int, k int) int {
	sort.Ints(nums)

	maxSize := 0
	for i, val := range nums {
		threshold := val + (2 * k)
		currentMax := 0

		l, r := 0, len(nums)-1
		for l <= r {
			m := (l + r) / 2

			if nums[m] <= threshold {
				currentMax = m
				l = m + 1
			} else {
				r = m - 1
			}
		}

		maxSize = max(maxSize, currentMax-i+1)
	}

	return maxSize
}
