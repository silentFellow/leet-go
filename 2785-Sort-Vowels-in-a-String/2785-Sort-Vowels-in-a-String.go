package leetcode

import (
	"slices"
	"strings"
)

func sortVowels(s string) string {
	vowels := map[rune]struct{}{
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}

	vowelsInstr := make([]rune, 0, len(s))
	for _, char := range s {
		if _, ok := vowels[char]; ok {
			vowelsInstr = append(vowelsInstr, char)
		}
	}

	slices.Sort(vowelsInstr)

	var ans strings.Builder
	idx := 0

	for _, char := range s {
		if _, ok := vowels[char]; ok {
			ans.WriteRune(vowelsInstr[idx])
			idx++
		} else {
			ans.WriteRune(char)
		}
	}

	return ans.String()
}
