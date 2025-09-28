package leetcode

import "slices"

func largestPerimeter(nums []int) int {
	slices.Sort(nums)
	n := len(nums)

	for i := n - 1; i >= 2; i-- {
		f, s, t := nums[i-2], nums[i-1], nums[i]
		if f+s > t {
			return f + s + t
		}
	}

	return 0
}
