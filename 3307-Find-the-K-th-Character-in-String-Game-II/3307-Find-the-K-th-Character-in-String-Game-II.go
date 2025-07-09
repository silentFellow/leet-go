package leetcode

// TODO: Optimize algorithm to prevent Time Limit Exceeded (TLE) errors.

import "strings"

func kthCharacter(k int64, operations []int) byte {
	getNext := func(v rune) rune {
		if v == 'z' {
			return 'a'
		}
		return (v + 1)
	}

	var str strings.Builder
	str.WriteString("a")

	idx := 0
	for str.Len() < int(k) && idx < len(operations) {
		switch operations[idx] {
		case 0:
			str.WriteString(str.String())
		case 1:
			for _, v := range str.String() {
				str.WriteRune(getNext(v))
			}
		}

		idx++
	}

	return str.String()[k-1]
}
