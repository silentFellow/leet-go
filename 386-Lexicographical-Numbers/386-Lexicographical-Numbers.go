package leetcode

func lexicalOrder(n int) []int {
	ans := make([]int, 0, n)

	cur := 1
	for range n {
		ans = append(ans, cur)

		if cur * 10 <= n {
			cur *= 10
		} else {
			for cur % 10 == 9 || cur+1 > n {
				cur /= 10
			}

			cur++
		}
	}

	return ans
}
