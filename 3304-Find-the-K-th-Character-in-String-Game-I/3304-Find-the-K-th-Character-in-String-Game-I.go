package leetcode

import "strings"

func kthCharacter(k int) byte {
	getNext := func(v rune) rune {
		if v == 'z' {
			return 'a'
		}

		return rune(v + 1)
	}

	var str strings.Builder
	str.WriteString("a")

	for str.Len() < k {
		for _, v := range str.String() {
			str.WriteRune(getNext(v))
		}
	}

	return str.String()[k-1]
}
