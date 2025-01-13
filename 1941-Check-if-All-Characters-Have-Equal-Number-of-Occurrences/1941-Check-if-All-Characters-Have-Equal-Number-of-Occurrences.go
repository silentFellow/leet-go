package leetcode

func areOccurrencesEqual(s string) bool {
	hmap := make(map[rune]int)

	for _, char := range s {
		hmap[char]++
	}

	match := hmap[rune(s[0])]
	for _, val := range hmap {
		if val != match {
			return false
		}
	}

	return true
}
