package leetcode

import "strings"

func findWordsContaining(words []string, x byte) []int {
	ans := []int{}

	for i, word := range words {
		if strings.Contains(word, string(x)) {
			ans = append(ans, i)
		}
	}

	return ans
}
