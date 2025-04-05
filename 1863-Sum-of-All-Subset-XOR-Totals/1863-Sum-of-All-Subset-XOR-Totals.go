package leetcode

// TODO: optimize it, current complexity 2^n

func subsetXORSum(nums []int) int {
	ans := 0

	var dfs func(idx, running int)
	dfs = func(idx, running int) {
		if idx == len(nums) {
			ans += running
			return
		}

		dfs(idx+1, running)
		dfs(idx+1, running^nums[idx])
	}

	dfs(0, 0)
	return ans
}
