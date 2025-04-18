package leetcode

func countGood(nums []int, k int) int64 {
	n := len(nums)
	hmap := make(map[int]int)
	left, right := 0, 0
	possible := 0
	ans := 0

	for right < n {
		// Add the current number to the map and update `possible`
		hmap[nums[right]]++
		possible += hmap[nums[right]] - 1

		// Shrink the window from the left until `possible` is less than `k`
		for possible >= k {
			ans += n - right
			hmap[nums[left]]--
			possible -= hmap[nums[left]]
			left++
		}

		right++
	}

	return int64(ans)
}
