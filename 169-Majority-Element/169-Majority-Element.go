package leetcode

// Boyer-Moore Voting Algorithm
func majorityElement(nums []int) int {
	candidate, count := nums[0], 0

	for _, val := range nums {
		if count == 0 {
			candidate = val
		}

		if candidate == val {
			count++
		} else {
			count--
		}
	}

	return candidate
}
