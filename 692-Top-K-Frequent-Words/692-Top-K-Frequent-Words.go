package leetcode

import (
	"sort"
	"strings"
)

type wordFreq struct {
	word string
	freq int
}

func topKFrequent(words []string, k int) []string {
	hmap := make(map[string]int)
	for _, val := range words {
		hmap[val] = hmap[val] + 1
	}

	freq := make([]wordFreq, len(hmap))
	index := 0
	for key, val := range hmap {
		freq[index] = wordFreq{
			word: key,
			freq: val,
		}
		index++
	}

	sort.Slice(freq, func(i, j int) bool {
		if freq[i].freq == freq[j].freq {
			return strings.Compare(freq[i].word, freq[j].word) < 0
		}
		return freq[i].freq > freq[j].freq
	})

	ans := make([]string, k)
	for i := range k {
		ans[i] = freq[i].word
	}

	return ans
}
