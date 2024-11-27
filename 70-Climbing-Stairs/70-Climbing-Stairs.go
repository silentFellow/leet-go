package leetcode

func climbStairs(n int) int {
	f, s := 1, 2

	if n == 1 {
		return f
	} else if n == 2 {
		return s
	}

	for range n - 2 {
		f, s = s, f+s
	}

	return s
}
