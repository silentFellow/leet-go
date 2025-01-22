package leetcode

type pos struct {
	r int
	c int
}

func uniquePaths(m int, n int) int {
	count := 0
	hmap := make(map[pos]int)

	var dfs func(m, n int) int
	dfs = func(r, c int) int {
		if val, ok := hmap[pos{r, c}]; ok {
			return val
		} else if r >= m || c >= n {
			return 0
		} else if r == m-1 && c == n-1 {
			return 1
		}

		right := dfs(r, c+1)
		down := dfs(r+1, c)

		val := right + down
		hmap[pos{r, c}] = val
		return val
	}

	count = dfs(0, 0)
	return count
}
