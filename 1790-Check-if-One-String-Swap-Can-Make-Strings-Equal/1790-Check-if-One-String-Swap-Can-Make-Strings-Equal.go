package leetcode

import "strings"

func areAlmostEqual(s1 string, s2 string) bool {
	var firstMismatch, secondMismatch strings.Builder
	count := 0

	for i := range len(s1) {
		if s1[i] != s2[i] {
			if count == 2 {
				return false
			}

			firstMismatch.WriteByte(s1[i])
			secondMismatch.WriteByte(s2[i])
      count++
		}
	}

	return firstMismatch.String() == reverse(secondMismatch.String())
}

func reverse(s string) string {
	runes := []rune(s)

	l, r := 0, len(runes)-1
	for l < r {
		runes[l], runes[r] = runes[r], runes[l]
    l, r = l+1, r-1
	}

	return string(runes)
}
