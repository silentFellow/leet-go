package leetcode

import "math"

func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	ans := math.MinInt

	prefix := make([]int, n+1)
	for i := range n {
		prefix[i+1] = prefix[i] + int(nums[i])
	}

	for offset := range k {
		minPrefix := prefix[offset]
		for i := offset + k; i <= n; i += k {
			ans = max(ans, prefix[i] - minPrefix)
			minPrefix = min(minPrefix, prefix[i])
		}
	}

	return int64(ans)
}
