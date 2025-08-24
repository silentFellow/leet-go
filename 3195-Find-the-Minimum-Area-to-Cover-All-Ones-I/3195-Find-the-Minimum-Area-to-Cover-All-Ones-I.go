package leetcode

import "math"

func minimumArea(grid [][]int) int {
	rowStart, colStart := math.MaxInt, math.MaxInt
	var rowEnd, colEnd int

	for i, row := range grid {
		for j, col := range row {
			if col == 1 {
				rowStart = min(rowStart, i)
				colStart = min(colStart, j)

				rowEnd = max(rowStart, i)
				colEnd = max(colEnd, j)
			}
		}
	}

	return (rowEnd - rowStart + 1) * (colEnd - colStart + 1)
}
