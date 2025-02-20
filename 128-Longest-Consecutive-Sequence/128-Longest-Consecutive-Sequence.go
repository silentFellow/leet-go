package leetcode

import "sort"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	ans, cur := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			continue
		}
		if nums[i-1]+1 == nums[i] {
			cur++
			continue
		}

		ans = max(ans, cur)
		cur = 1
	}
	ans = max(ans, cur)

	return ans
}
