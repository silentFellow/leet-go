package leetcode

import "math"

func maxProduct(nums []int) int {
	f, s := math.MinInt, math.MinInt

	for _, val := range nums {
		if val > f {
			s, f = f, val
		} else if val > s {
			s = val
		}
	}

	return (f - 1) * (s - 1)
}
