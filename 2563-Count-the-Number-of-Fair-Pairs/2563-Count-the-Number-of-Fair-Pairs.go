package leetcode

import "sort"

func countFairPairs(nums []int, lower int, upper int) int64 {
	ans := 0
	n := len(nums)

	sort.Ints(nums)
	for i := range n {
		low := lower - nums[i]
		high := upper - nums[i]

		left := sort.Search(n, func(j int) bool {
			return nums[j] >= low
		})

		right := sort.Search(n, func(j int) bool {
			return nums[j] > high
		})

		if right > i+1 {
			ans += right - max(left, i+1)
		}
	}

	return int64(ans)
}
