package leetcode

import (
	"strings"
	"unicode"
)

func clearDigits(s string) string {
	notAllowed := make([]bool, len(s))
	charIdx := []int{}

	for i, val := range s {
		if unicode.IsLetter(val) {
			charIdx = append([]int{i}, charIdx...)
		}

		if unicode.IsDigit(val) {
			notAllowed[i] = true

			if len(charIdx) > 0 {
				notAllowed[charIdx[0]] = true
				charIdx = charIdx[1:]
			}
		}
	}

	var ans strings.Builder
	for i, val := range s {
		if !notAllowed[i] {
			ans.WriteRune(val)
		}
	}

	return ans.String()
}
