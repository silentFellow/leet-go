package leetcode

func countWords(words1 []string, words2 []string) int {
	hmap := make(map[string]int)
	for _, word := range words1 {
		hmap[word]++
	}

	for _, word := range words2 {
		if val, ok := hmap[word]; ok {
			if val > 1 {
				continue
			}
		}
		hmap[word]--
	}

	count := 0
	for _, val := range hmap {
		if val == 0 {
			count++
		}
	}

	return count
}
