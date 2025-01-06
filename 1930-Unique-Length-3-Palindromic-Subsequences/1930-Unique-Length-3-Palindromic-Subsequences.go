package leetcode

func countPalindromicSubsequence(s string) int {
	first, last := [26]int{}, [26]int{}

	for i := range 26 {
		first[i] = -1
	}

	for i, char := range s {
		charIdx := char - 'a'
		if first[charIdx] == -1 {
			first[charIdx] = i
		}
		last[charIdx] = i
	}

	count := 0
	for char := range 26 {
		if first[char] != -1 && last[char] > first[char] {
			set := make(map[rune]struct{})
			for i := first[char] + 1; i < last[char]; i++ {
				set[rune(s[i])] = struct{}{}
			}

			count += len(set)
		}
	}

	return count
}
