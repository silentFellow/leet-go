package leetcode

import "strings"

func countPrefixes(words []string, s string) int {
	count := 0

	for _, word := range words {
		if strings.HasPrefix(s, word) {
			count++
		}
	}

	return count
}
