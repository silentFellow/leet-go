package leetcode

func sumZero(n int) []int {
	ans := make([]int, 0, n)

	if n%2 == 1 {
		ans = append(ans, 0)
	}

	for i := 1; i <= n/2; i++ {
		ans = append(ans, i)
		ans = append(ans, -i)
	}

	return ans
}
