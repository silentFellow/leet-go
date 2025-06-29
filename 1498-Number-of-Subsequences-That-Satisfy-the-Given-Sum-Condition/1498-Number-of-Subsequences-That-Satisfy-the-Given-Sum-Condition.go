package leetcode

import (
	"slices"
)

var MOD int = 1_000_000_007

func numSubseq(nums []int, target int) int {
	n := len(nums)
	slices.Sort(nums)

	// Precompute powers of 2 modulo MOD
	pow2 := make([]int, n+1)
	pow2[0] = 1
	for i := 1; i <= n; i++ {
		pow2[i] = (pow2[i-1] * 2) % MOD
	}

	ans := 0
	left, right := 0, n-1

	for left <= right {
		if nums[left]+nums[right] <= target {
			ans = (ans + pow2[right-left]) % MOD
			left++
		} else {
			right--
		}
	}

	return ans
}
