package leetcode

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	memo := make(map[[2]int]bool)

	var dfs func(i, cur int) bool
	dfs = func(i, cur int) bool {
		if cur == target {
			return true
		} else if i == len(nums) || cur > target {
			return false
		}

		key := [2]int{i, cur}
		if val, ok := memo[key]; ok {
			return val
		}

		pick := dfs(i+1, cur+nums[i])
		skip := dfs(i+1, cur)

		memo[key] = pick || skip
		return memo[key]
	}

	return dfs(0, 0)
}
