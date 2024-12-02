package leetcode

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	words := strings.Fields(sentence) // words := strings.Split(sentence, " ")

	for i, val := range words {
		searchIndex := strings.Index(val, searchWord)
		if searchIndex == 0 {
			return (i + 1)
		}
	}

	return -1
}
