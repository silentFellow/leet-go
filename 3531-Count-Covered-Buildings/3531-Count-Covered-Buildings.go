package leetcode

import "math"

func countCoveredBuildings(n int, buildings [][]int) int {
	minX, maxX := make([]int, n+1), make([]int, n+1)
	minY, maxY := make([]int, n+1), make([]int, n+1)

	for i := 1; i <= n; i++ {
		minX[i] = math.MaxInt
		minY[i] = math.MaxInt
	}

	for _, b := range buildings {
		x, y := b[0], b[1]

		minX[y] = min(minX[y], x)
		maxX[y] = max(maxX[y], x)

		minY[x] = min(minY[x], y)
		maxY[x] = max(maxY[x], y)
	}

	ans := 0
	for _, b := range buildings {
		x, y := b[0], b[1]

		if minY[x] < y && y < maxY[x] && minX[y] < x && x < maxX[y] {
			ans++
		}
	}

	return ans
}
