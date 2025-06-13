package leetcode

import "strings"

func addBinary(a string, b string) string {
	var revAns strings.Builder

	rem := '0'
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 {
		var first, second rune
		if i < 0 {
			first = '0'
		} else {
			first = rune(a[i])
		}

		if j < 0 {
			second = '0'
		} else {
			second = rune(b[j])
		}

		if first == '1' && second == '1' && rem == '1' {
			revAns.WriteByte('1')
			rem = '1'
		} else if (first == '1' && second == '1') || (first == '1' && rem == '1') || (second == '1' && rem == '1') {
			revAns.WriteByte('0')
			rem = '1'
		} else if first == '1' || second == '1' || rem == '1' {
			revAns.WriteByte('1')
			rem = '0'
		} else {
			revAns.WriteByte('0')
			rem = '0'
		}

		i--
		j--
	}

	if rem != '0' {
		revAns.WriteByte('1')
	}

	reverse := func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}

		return string(runes)
	}

	return reverse(revAns.String())
}
