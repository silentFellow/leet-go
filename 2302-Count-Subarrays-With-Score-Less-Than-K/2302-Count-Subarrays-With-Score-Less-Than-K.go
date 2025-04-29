package leetcode

func countSubarrays(nums []int, k int64) int64 {
	ans := 0
	n := len(nums)

	left := 0
	sum := 0
	for right := range n {
		sum += nums[right]
		for (sum * (right - left + 1)) >= int(k) {
			sum -= nums[left]
			left++
		}

		ans += (right - left + 1)
	}

	return int64(ans)
}
