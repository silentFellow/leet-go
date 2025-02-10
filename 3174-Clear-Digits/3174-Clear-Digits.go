package leetcode

import (
	"strings"
	"unicode"
)

func clearDigits(s string) string {
	notAllowed := make([]bool, len(s))
	characters := []int{}

	for i, val := range s {
		if unicode.IsLetter(val) {
			characters = append([]int{i}, characters...)
		}

		if unicode.IsDigit(val) {
			notAllowed[i] = true

			if len(characters) > 0 {
				notAllowed[characters[0]] = true
				characters = characters[1:]
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
