package leetcode

import "unicode"

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}

	isVowel, isConsonant := false, false
	vowelSet := map[rune]struct{}{
		'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {},
		'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {},
	}

	for _, v := range word {
		if !(unicode.IsLetter(v) || unicode.IsDigit(v)) {
			return false
		}

		if !unicode.IsDigit(v) {
			if _, ok := vowelSet[v]; ok {
				isVowel = true
			} else {
				isConsonant = true
			}
		}
	}

	return isVowel && isConsonant
}
