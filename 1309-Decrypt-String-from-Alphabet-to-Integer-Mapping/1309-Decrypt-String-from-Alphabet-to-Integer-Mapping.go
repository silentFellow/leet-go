package leetcode

import "strings"

func freqAlphabets(s string) string {
	var ans strings.Builder
	idx := len(s) - 1

	for idx >= 0 {
		if s[idx] == '#' {
			num := (s[idx-2]-'0')*10 + (s[idx-1] - '0')
			ans.WriteByte('a' + num - 1)
			idx -= 3
		} else {
			num := s[idx] - '0'
			ans.WriteByte('a' + num - 1)
			idx -= 1
		}
	}

	ansBytes := []byte(ans.String())
	for i, j := 0, len(ansBytes)-1; i < j; i, j = i+1, j-1 {
		ansBytes[i], ansBytes[j] = ansBytes[j], ansBytes[i]
	}

	return string(ansBytes)
}
