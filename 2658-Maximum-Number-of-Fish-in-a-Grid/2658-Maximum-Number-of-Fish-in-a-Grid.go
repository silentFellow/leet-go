package leetcode

type pair struct {
	row int
	col int
}

func findMaxFish(grid [][]int) int {
	rowSize, colSize := len(grid), len(grid[0])

	visited := make(map[pair]bool)

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if grid[i][j] == 0 || visited[pair{row: i, col: j}] {
			return 0
		}

		visited[pair{row: i, col: j}] = true

		left, right, up, down := 0, 0, 0, 0
		if j != 0 {
			left = dfs(i, j-1)
		}
		if j < colSize-1 {
			right = dfs(i, j+1)
		}
		if i != 0 {
			up = dfs(i-1, j)
		}
		if i < rowSize-1 {
			down = dfs(i+1, j)
		}

		total := left + right + up + down + grid[i][j]
		return total
	}

	ans := 0
	for i, row := range grid {
		for j, col := range row {
			if col != 0 && !visited[pair{row: i, col: j}] {
				ans = max(ans, dfs(i, j))
			}
		}
	}

	return ans
}
