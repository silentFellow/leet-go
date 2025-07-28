package leetcode

func countMaxOrSubsets(nums []int) int {
	maxVal := nums[0]
	for _, v := range nums {
		maxVal |= v
	}

	ans := 0
	var dfs func(idx, cur int)
	dfs = func(idx, cur int) {
		if idx == len(nums) {
			if cur == maxVal {
				ans++
			}
			return
		}

		dfs(idx+1, cur|nums[idx])
		dfs(idx+1, cur)
	}

	dfs(0, 0)
	return ans
}
