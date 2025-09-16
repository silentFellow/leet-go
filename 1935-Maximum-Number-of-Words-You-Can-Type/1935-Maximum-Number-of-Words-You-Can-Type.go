package leetcode

import "strings"

func canBeTypedWords(text string, brokenLetters string) int {
	words := strings.Split(text, " ")

	ans := 0
	for _, word := range words {
		if !strings.ContainsAny(word, brokenLetters) {
			ans++
		}
	}

	return ans
}
