package leetcode

func totalMoney(n int) int {
	start := 0
	ans := 0

	for i := range n {
		if i%7 == 0 {
			start++
		}

		ans += (start + (i % 7))
	}

	return ans
}
