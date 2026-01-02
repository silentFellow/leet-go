package leetcode

func countNegatives(grid [][]int) int {
	rowSize, colSize := len(grid), len(grid[0])
	ans := 0

	r, c := rowSize-1, 0
	for r >= 0 && c < colSize {
		v := grid[r][c]

		if v < 0 {
			ans += (colSize - c)
			r--
		} else {
			c++
		}
	}

	return ans
}
