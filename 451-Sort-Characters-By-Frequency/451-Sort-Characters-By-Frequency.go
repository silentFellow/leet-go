package leetcode

import (
	"sort"
	"strings"
)

type pair struct {
  char byte
  count int
}

func frequencySort(s string) string {
	hmap := make(map[byte]int)
	for _, val := range []byte(s) {
		hmap[val] = hmap[val] + 1
	}

  freqPairs := make([]pair, len(hmap))
  for char, freq := range hmap {
    freqPairs = append(freqPairs, pair{char, freq})
  }

	sort.Slice(freqPairs, func(i, j int) bool {
    return freqPairs[i].count > freqPairs[j].count
	})

	var sb strings.Builder
  for _, val := range freqPairs {
    sb.WriteString(strings.Repeat(string(val.char), val.count))
  }

	return sb.String()
}
