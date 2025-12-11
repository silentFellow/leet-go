package leetcode

func countTriples(n int) int {
	ans := 0

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			l, r := 1, n

			for l <= r {
				m := (l + r) / 2

				if (i*i)+(j*j) == m*m {
					ans++
					break
				} else if (i*i)+(j*j) < m*m {
					r = m - 1
				} else {
					l = m + 1
				}
			}
		}
	}

	return ans
}
