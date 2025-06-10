package leetcode

// TODO: optimize to avoid TLE

func findKthNumber(n int, k int) int {
	cur, pos := 1, 1
	for range n {
		if pos == k {
			return cur
		}

		if cur*10 <= n {
			cur *= 10
		} else {
			for cur%10 == 9 || cur+1 > n {
				cur /= 10
			}
			cur++
		}

		pos++
	}

	return -1
}
