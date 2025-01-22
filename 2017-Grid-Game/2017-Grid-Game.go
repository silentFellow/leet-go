package leetcode

import "math"

func gridGame(grid [][]int) int64 {
	n := len(grid[0])

	prefixFirst, prefixSecond := make([]int, n), make([]int, n)
	copy(prefixFirst, grid[0])
	copy(prefixSecond, grid[1])

	for i := 1; i < n; i++ {
		prefixFirst[i] += prefixFirst[i-1]
		prefixSecond[i] += prefixSecond[i-1]
	}

	ans := math.MaxInt64
	for i := range n {
		top := prefixFirst[n-1] - prefixFirst[i]

		var bottom int
		if i > 0 {
			bottom = prefixSecond[i-1]
		} else {
			bottom = 0
		}

		secondRobot := max(top, bottom)
		ans = min(ans, secondRobot)
	}

	return int64(ans)
}
