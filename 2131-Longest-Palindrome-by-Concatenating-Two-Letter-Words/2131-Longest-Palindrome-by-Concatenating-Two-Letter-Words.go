package leetcode

func longestPalindrome(words []string) int {
	ans := 0

	reverse := func(word string) string {
		runes := []rune(word)

		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}

		return string(runes)
	}

	hmap := make(map[string]int)
	for _, word := range words {
		rev := reverse(word)

		if hmap[rev] > 0 {
			ans += 4
			hmap[rev]--
			if hmap[rev] == 0 {
				delete(hmap, rev)
			}
		} else {
			hmap[word]++
		}
	}

	// After pairing, check for any remaining symmetric word like "aa"
	for key := range hmap {
		if key == reverse(key) {
			ans += 2
			break
		}
	}

	return ans
}
