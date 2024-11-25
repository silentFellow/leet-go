package leetcode

// TODO: verify once after sometime

import (
	"sort"
	"strings"
)

func kthLargestNumber(nums []string, k int) string {
	sort.Slice(nums, func(i,j int) bool {
		if len(nums[i]) == len(nums[j]) {
			return strings.Compare(nums[i], nums[j]) > 0
		} else {
			return len(nums[i]) > len(nums[j])
		}
	})

	return nums[k-1]
}
