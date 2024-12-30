package leetcode

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	f, s := 0, 1
	for i := 2; i <= n; i++ {
		f, s = s, f+s
	}

	return s
}
