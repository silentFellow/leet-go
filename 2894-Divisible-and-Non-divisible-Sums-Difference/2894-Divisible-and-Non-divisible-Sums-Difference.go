package leetcode

func differenceOfSums(n int, m int) int {
	f, s := 0, 0

	for i := 1; i <= n; i++ {
		if i%m != 0 {
			f += i
		} else {
			s += i
		}
	}

	return f - s
}
