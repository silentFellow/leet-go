package leetcode

func maximumDifference(nums []int) int {
	curMin := nums[0]
	ans := -1

	for i := 1; i < len(nums); i++ {
		// to avoid duplicate answer resulting in answer to be 0, rather than -1
		if nums[i] != curMin {
			ans = max(ans, nums[i]-curMin)
		}
		curMin = min(curMin, nums[i])
	}

	return ans
}
