package leetcode

import "math"

func thirdMax(nums []int) int {
	m, sm, tm := math.MinInt, math.MinInt, math.MinInt

	for _, val := range nums {
		if val > m {
			tm, sm, m = sm, m, val
		} else if val > sm && val != m {
			tm, sm = sm, val
		} else if val > tm && val != m && val != sm {
			tm = val
		}
	}

	if tm == math.MinInt {
		return m
	} else {
		return tm
	}
}
