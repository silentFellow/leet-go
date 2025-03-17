package leetcode

import "math"

func minCapability(nums []int, k int) int {
	// Binary search range
	left, right := math.MaxInt, math.MinInt

	// Find the range of possible capability values
	for _, num := range nums {
		left = min(left, num)
		right = max(right, num)
	}

	// Helper function to check if we can pick `k` houses with `capability`
	canPick := func(capability int) bool {
		count, i := 0, 0
		for i < len(nums) {
			if nums[i] <= capability { // Can pick this house
				count++
				i += 2 // Move to next non-adjacent house
			} else {
				i++ // Skip this house
			}
			if count >= k {
				return true
			}
		}
		return count >= k
	}

	// Perform binary search on the capability range
	for left < right {
		mid := (left + right) / 2
		if canPick(mid) {
			right = mid // Try smaller max capability
		} else {
			left = mid + 1 // Increase the capability limit
		}
	}

	return left
}
