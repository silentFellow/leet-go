package leetcode

import "strings"

func spellchecker(wordlist []string, queries []string) []string {
	exact := make(map[string]string)
	caseInsensitive := make(map[string]string)
	vowelInsensitive := make(map[string]string)

	formatToVowelInsensitive := func(word string) string {
		var formatted strings.Builder

		for _, char := range word {
			if strings.ContainsRune("aeiou", char) {
				formatted.WriteRune('*')
			} else {
				formatted.WriteRune(char)
			}
		}

		return formatted.String()
	}

	for _, word := range wordlist {
		exact[word] = word

		v := strings.ToLower(word)
		if _, ok := caseInsensitive[v]; !ok {
			caseInsensitive[v] = word
		}

		v = formatToVowelInsensitive(strings.ToLower(word))
		if _, ok := vowelInsensitive[v]; !ok {
			vowelInsensitive[v] = word
		}
	}

	ans := make([]string, len(queries))
	for i, query := range queries {
		if word, ok := exact[query]; ok {
			ans[i] = word
		} else if word, ok := caseInsensitive[strings.ToLower(query)]; ok {
			ans[i] = word
		} else if word, ok := vowelInsensitive[formatToVowelInsensitive(strings.ToLower(query))]; ok {
			ans[i] = word
		} else {
			ans[i] = ""
		}
	}

	return ans
}
