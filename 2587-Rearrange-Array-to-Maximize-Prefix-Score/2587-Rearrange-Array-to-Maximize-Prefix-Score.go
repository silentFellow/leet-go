package leetcode

import "slices"

func maxScore(nums []int) int {
	slices.SortFunc(nums, func(i, j int) int {
		return j - i
	})

	prefix, ans := 0, 0
	for _, v := range nums {
		prefix += v
		if prefix > 0 {
			ans++
		} else {
			break
		}
	}

	return ans
}
