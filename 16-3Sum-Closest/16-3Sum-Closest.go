package leetcode

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	ans := nums[0] + nums[1] + nums[2]

	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1

		for l < r {
			val := nums[i] + nums[l] + nums[r]
			if val == target {
				return val
			}

			absDiffVal := abs(val - target)
			absDiffAns := abs(ans - target)

			if absDiffVal < absDiffAns {
				ans = val
			}

			if val < target {
				l++
			} else {
				r--
			}
		}
	}

	return ans
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
