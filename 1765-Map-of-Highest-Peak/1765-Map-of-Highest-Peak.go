package leetcode

func highestPeak(isWater [][]int) [][]int {
	rowLen, colLen := len(isWater), len(isWater[0])

	res := make([][]int, rowLen)
	for i := range rowLen {
		col := make([]int, colLen)
		for j := range colLen {
			col[j] = -1
		}
		res[i] = col
	}

	queue := [][]int{}
	for i := range rowLen {
		for j := range colLen {
			if isWater[i][j] == 1 {
				queue = append(queue, []int{i, j})
				res[i][j] = 0
			}
		}
	}

	directions := [][]int{
		{0, -1}, {0, 1}, {-1, 0}, {1, 0},
	}

	for len(queue) != 0 {
		first := queue[0]
		nr, nc := first[0], first[1]
		queue = queue[1:]

		for _, dir := range directions {
			r, c := nr+dir[0], nc+dir[1]
			if r >= 0 && r < rowLen && c >= 0 && c < colLen && res[r][c] == -1 {
				res[r][c] = res[nr][nc] + 1
				queue = append(queue, []int{r, c})
			}
		}
	}

	return res
}
