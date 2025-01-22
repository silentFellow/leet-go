package leetcode

func tribonacci(n int) int {
	if n <= 0 {
		return 0
	} else if n <= 2 {
		return 1
	} else if n <= 3 {
		return 2
	}

	f, s, t := 1, 1, 2

	for range n - 4 {
		f, s, t = s, t, f+s+t
	}

	return f + s + t
}
