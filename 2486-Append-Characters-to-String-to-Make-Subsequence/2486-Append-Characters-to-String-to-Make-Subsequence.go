package leetcode

func appendCharacters(s string, t string) int {
	l, r := 0, 0

	for l < len(s) && r < len(t) {
		if s[l] == t[r] {
			l++
			r++
			continue
		}

		l++
	}

	return len(t) - r
}
