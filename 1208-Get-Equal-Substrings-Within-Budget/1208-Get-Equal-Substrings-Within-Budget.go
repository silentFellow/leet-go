package leetcode

func equalSubstring(s string, t string, maxCost int) int {
	ans := 0
	l, r := 0, 0

	cur := 0
	score := 0
	for r < len(s) {
		diff := abs(int(s[r]) - int(t[r]))
		score += diff
		cur++

		if score <= maxCost {
			ans = max(ans, cur)
		} else {
			score -= abs(int(s[l]) - int(t[l]))
			cur--
			l++
		}
		r++
	}

	return ans
}

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}
