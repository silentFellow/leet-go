package leetcode

type pos struct {
	r, c int
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	rowLen, colLen := len(obstacleGrid), len(obstacleGrid[0])
	hmap := make(map[pos]int)

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if val, ok := hmap[pos{r, c}]; ok {
			return val
		} else if r == rowLen-1 && c == colLen-1 {
			return 1
		}

		right, bottom := 0, 0
		if c+1 < colLen && obstacleGrid[r][c+1] != 1 {
			right = dfs(r, c+1)
		}
		if r+1 < rowLen && obstacleGrid[r+1][c] != 1 {
			bottom = dfs(r+1, c)
		}

		val := right + bottom
		hmap[pos{r, c}] = val
		return val
	}

	return dfs(0, 0)
}
