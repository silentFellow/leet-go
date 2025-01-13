package leetcode

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	hmap := make(map[string]int)

	for _, word := range strings.Fields(s1) {
		hmap[word]++
	}

	for _, word := range strings.Fields(s2) {
		hmap[word]++
	}

	ans := []string{}
	for key, val := range hmap {
		if val == 1 {
			ans = append(ans, key)
		}
	}

	return ans
}
