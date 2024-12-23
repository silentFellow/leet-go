package leetcode

// TODO: verify once again

import (
	"math"
	"slices"
)

func minimumSize(nums []int, maxOperations int) int {
	canDivide := func(n int) bool {
		operations := 0

		for _, val := range nums {
			required := int(math.Ceil(float64(val)/float64(n))) - 1
			operations += required

			if operations > maxOperations {
				return false
			}
		}

		return true
	}

	l, r := 0, slices.Max(nums)
	ans := -1
	for l <= r {
		m := (l + r) / 2

		if canDivide(m) {
			ans = m
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return ans
}
