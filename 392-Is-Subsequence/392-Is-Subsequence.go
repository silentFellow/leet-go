package leetcode

func isSubsequence(s string, t string) bool {
	l, r := 0, 0

	for l < len(s) && r < len(t) {
		if s[l] == t[r] {
			l++
			r++
			continue
		}

		r++
	}

	if l == len(s) {
		return true
	}

	return false
}
