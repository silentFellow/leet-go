package leetcode

func longestSubarray(nums []int) int {
	left, zeros, ans := 0, 0, 0

	for right := range len(nums) {
		if nums[right] == 0 {
			zeros++
		}

		for zeros > 1 {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}

		ans = max(ans, right-left)
	}

	return ans
}
