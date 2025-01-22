package leetcode

import "math"

type pos struct {
	r, c int
}

func minPathSum(grid [][]int) int {
	hmap := make(map[pos]int)

	rowLen, colLen := len(grid), len(grid[0])
	var dfs func(r, c int) int
	dfs = func(r, c int) int {
    if val, ok := hmap[pos{r, c}]; ok {
      return val
    } else if r == rowLen-1 && c == colLen-1 {
      return grid[r][c]
    } else if r >= rowLen || c >= colLen {
      return math.MaxInt64
    }

    right := dfs(r, c+1)
    down := dfs(r+1, c)

    val := grid[r][c] + min(right, down)
    hmap[pos{r, c}] = val
    return val
	}

	return dfs(0, 0)
}
