package leetcode

// NOTE: optimal solution:
// return int64(1 + n * (n-1) * 2)

func coloredCells(n int) int64 {
	ans := 1
	for i := range n {
		ans += i * 4
	}

	return int64(ans)
}
