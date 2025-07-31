package leetcode

func longestSubarray(nums []int) int {
	maxv := nums[0]
	for _, v := range nums {
		maxv = max(maxv, v)
	}

	ans := 0
	cur := 0
	for _, v := range nums {
		if v != maxv {
			ans = max(ans, cur)
			cur = 0
			continue
		}
		cur++
	}
	ans = max(ans, cur)

	return ans
}
