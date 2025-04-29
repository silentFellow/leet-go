package leetcode

func countSubarrays(nums []int, k int) int64 {
	maxv := nums[0]
	for _, v := range nums {
		maxv = max(maxv, v)
	}

	n := len(nums)
	left, count, ans := 0, 0, 0
	for right := range n {
		if nums[right] == maxv {
			count++
		}

		for count == k {
			if nums[left] == maxv {
				count--
			}
			left++
		}

		ans += left
	}

	return int64(ans)
}
