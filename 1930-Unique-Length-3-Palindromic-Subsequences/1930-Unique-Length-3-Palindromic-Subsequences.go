package leetcode

func countPalindromicSubsequence(s string) int {
	first, last := [26]int{}, [26]int{}

	for i := range 26 {
		first[i] = -1
		last[i] = -1
	}

	for i, v := range s {
		pos := v - 'a'
		if first[pos] == -1 {
			first[pos] = i
		}
		last[pos] = i
	}

	ans := 0
	for i := range 26 {
		if first[i] == -1 {
			continue
		}

		set := make(map[rune]struct{})
		for j := first[i] + 1; j < last[i]; j++ {
			set[rune(s[j])] = struct{}{}
		}

		ans += len(set)
	}

	return ans
}
