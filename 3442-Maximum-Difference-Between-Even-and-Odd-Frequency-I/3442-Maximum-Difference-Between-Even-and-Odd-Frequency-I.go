package leetcode

import "math"

func maxDifference(s string) int {
	hmap := make(map[rune]int)

	for _, char := range s {
		hmap[char]++
	}

	oc, ec := 0, math.MaxInt64
	for _, freq := range hmap {
		if freq%2 == 0 {
			ec = min(ec, freq)
		} else {
			oc = max(oc, freq)
		}
	}

	return oc - ec
}
