package leetcode

func countOfSubstrings(word string, k int) int64 {
	return countK(word, k) - countK(word, k+1)
}

func isvowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c ==
		'o' || c == 'u'
}

func countK(word string, k int) int64 {
	vowels := make(map[byte]int)

	total := int64(0)
	cons := 0
	l := 0

	for r := 0; r < len(word); r++ {
		if isvowel(word[r]) {
			vowels[word[r]]++
		} else {
			cons++
		}

		for len(vowels) == 5 && cons >= k {
			total += int64(len(word) - r)
			if isvowel(word[l]) {
				vowels[word[l]]--
				if vowels[word[l]] == 0 {
					delete(vowels, word[l])
				}
			} else {
				cons--
			}
			l++
		}
	}

	return total
}
