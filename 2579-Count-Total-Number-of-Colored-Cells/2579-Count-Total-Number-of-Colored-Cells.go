package leetcode

// NOTE: optimal solution:
// 1. return int64(1 + n * (n-1) * 2)
// 2. return int64(math.Pow(float64(n), 2) + math.Pow(float64(n-1), 2))

func coloredCells(n int) int64 {
	ans := 1
	for i := range n {
		ans += i * 4
	}

	return int64(ans)
}
