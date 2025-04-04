package leetcode

func partitionLabels(s string) []int {
	last := make(map[rune]int)
	for i, char := range s {
		last[char] = i
	}

	ans := make([]int, 0)
	l, r := 0, 0
	for i, char := range s {
		r = max(r, last[char])

		if i == r {
			ans = append(ans, r-l+1)
			l = i + 1
		}
	}

	return ans
}
