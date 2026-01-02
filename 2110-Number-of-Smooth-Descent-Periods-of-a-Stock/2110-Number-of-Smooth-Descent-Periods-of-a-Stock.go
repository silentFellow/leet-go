package leetcode

func getDescentPeriods(prices []int) int64 {
	n := len(prices)
	ans := 0

	l := 0
	for r := 1; r < len(prices); r++ {
		if prices[r]+1 == prices[r-1] {
			continue
		}

		length := r - l
		ans += (length * (length + 1)) / 2
		l = r
	}

	length := n - l
	ans += (length * (length + 1)) / 2

	return int64(ans)
}
