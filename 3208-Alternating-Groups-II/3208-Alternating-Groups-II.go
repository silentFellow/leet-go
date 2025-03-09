package leetcode

func numberOfAlternatingGroups(colors []int, k int) int {
	ans := 0

	l, n := 0, len(colors)
	for r := 1; r < (n + k - 1); r++ {
		if colors[r%n] == colors[(r-1)%n] {
			l = r
		}

		if r-l+1 > k {
			l++
		}

		if r-l+1 == k {
			ans++
		}
	}

	return ans
}
