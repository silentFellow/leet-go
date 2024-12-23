package leetcode

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)

	l, r := 0, len(nums)-1

	ans := 0
	for l < r {
		sum := nums[l] + nums[r]

		if sum == k {
			ans++
			l++
			r--
		} else if sum < k {
			l++
		} else {
			r--
		}
	}

	return ans
}
