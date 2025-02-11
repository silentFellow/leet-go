package leetcode

import "strings"

func removeOccurrences(s string, part string) string {
	for strings.Index(s, part) != -1 {
		idx := strings.Index(s, part)
		s = s[:idx] + s[idx+len(part):]
	}

	return s
}
