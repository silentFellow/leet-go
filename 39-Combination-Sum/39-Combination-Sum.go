package leetcode

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)

	var dfs func(cur []int, running int, idx int)
	dfs = func(cur []int, running int, idx int) {
		if running == target {
			temp := make([]int, len(cur))
			copy(temp, cur)
			ans = append(ans, temp)
			return
		}

		if running > target {
			return
		}

		for i := idx; i < len(candidates); i++ {
			dfs(append(cur, candidates[i]), running+candidates[i], i)
		}
	}

	dfs([]int{}, 0, 0)
	return ans
}
