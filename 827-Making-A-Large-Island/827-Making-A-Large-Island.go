package leetcode

func largestIsland(grid [][]int) int {
	rowSize, colSize := len(grid), len(grid[0])
	dir := [][]int{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}

	outOfBounds := func(i, j int) bool {
		return i < 0 || i >= rowSize || j < 0 || j >= colSize
	}

	var dfs func(i, j, label int) int
	dfs = func(i, j, label int) int {
		if outOfBounds(i, j) || grid[i][j] == 0 || grid[i][j] == label {
			return 0
		}

		area := 1
		grid[i][j] = label
		for _, nei := range dir {
			area += dfs(i+nei[0], j+nei[1], label)
		}
		return area
	}

	label := 2
	areaMap := make(map[int]int)
	ans := len(areaMap)

	for i, row := range grid {
		for j, col := range row {
			if col == 1 {
				area := dfs(i, j, label)
				ans = max(ans, area)
				areaMap[label] = area
				label++
			}
		}
	}

	var connect func(i, j int) int
	connect = func(i, j int) int {
		visited := make(map[int]struct{})
		area := 1

		for _, nei := range dir {
			ni, nj := i+nei[0], j+nei[1]
			if !outOfBounds(ni, nj) && grid[ni][nj] > 1 {
				label := grid[ni][nj] // Get the island label
				if _, ok := visited[label]; !ok {
					area += areaMap[label]
					visited[label] = struct{}{}
				}
			}
		}

		return area
	}

	for i, row := range grid {
		for j, col := range row {
			if col == 0 {
				ans = max(ans, connect(i, j))
			}
		}
	}

	return ans
}
