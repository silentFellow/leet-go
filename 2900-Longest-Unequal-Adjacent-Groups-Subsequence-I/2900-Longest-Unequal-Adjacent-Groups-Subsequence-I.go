package leetcode

func getLongestSubsequence(words []string, groups []int) []string {
	ans := []string{}

	prev := groups[0]

	findMax := func(l, r int) string {
		ans := words[l]
		for i := l; i <= r; i++ {
			if len(words[i]) > len(ans) {
				ans = words[i]
			}
		}
		return ans
	}

	s := 0
	for i := 1; i < len(groups); i++ {
		if groups[i] == prev {
			continue
		}

		ans = append(ans, findMax(s, i-1))

		s = i
		prev = groups[i]
	}
	ans = append(ans, findMax(s, len(groups)-1))

	return ans
}
