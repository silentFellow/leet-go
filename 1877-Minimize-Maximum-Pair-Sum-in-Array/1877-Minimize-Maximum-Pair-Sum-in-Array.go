package leetcode

import (
	"math"
	"slices"
)

func minPairSum(nums []int) int {
	n := len(nums)
	ans := math.MinInt

	slices.Sort(nums)
	for i:=0; i<n/2; i++ {
		ans = max(ans, nums[i]+nums[n-i-1])
	}

	return ans
}
