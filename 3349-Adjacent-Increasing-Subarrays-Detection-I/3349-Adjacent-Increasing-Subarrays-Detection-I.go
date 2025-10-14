package leetcode

func hasIncreasingSubarrays(nums []int, k int) bool {
	n := len(nums)
	for i := 0; i+2*k <= n; i++ {
		first := true
		second := true

		for j := i + 1; j < i+k; j++ {
			if nums[j] <= nums[j-1] {
				first = false
				break
			}
		}

		for j := i + k + 1; j < i+2*k; j++ {
			if nums[j] <= nums[j-1] {
				second = false
				break
			}
		}

		if first && second {
			return true
		}
	}

	return false
}
