package leetcode

// TODO: optimize the solution

func numMatchingSubseq(s string, words []string) int {
	count := 0

	for _, word := range words {
		if isSubString(s, word) {
			count++
		}
	}

	return count
}

func isSubString(s, t string) bool {
	l, r := 0, 0

	for l < len(s) && r < len(t) {
		if s[l] == t[r] {
      // l++, r++, continue will cause TLE: better optimize the solution

			r++
		}

		l++
	}

	return r == len(t)
}
