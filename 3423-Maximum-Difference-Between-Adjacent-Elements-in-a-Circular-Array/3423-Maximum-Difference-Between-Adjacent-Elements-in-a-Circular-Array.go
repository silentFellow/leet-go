package leetcode

func maxAdjacentDistance(nums []int) int {
	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}

	n := len(nums)
	ans := abs(nums[0] - nums[n-1])
	for i := range n - 1 {
		ans = max(ans, abs(nums[i]-nums[i+1]))
	}

	return ans
}
