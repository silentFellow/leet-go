package leetcode

// TODO: Follow up: Squaring each element and sorting the new array is very trivial, could you find an O(n) solution using a different approach?

import (
	"sort"
)

func sortedSquares(nums []int) []int {
	for i := range len(nums) {
		nums[i] = nums[i] * nums[i]
	}

  sort.Ints(nums)

	return nums
}
