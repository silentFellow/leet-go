package leetcode

import "strings"

func makeFancyString(s string) string {
	splits := []string{}

	var cur strings.Builder
	cur.WriteByte(s[0])
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			splits = append(splits, cur.String())
			cur.Reset()
		}

		cur.WriteByte(s[i])
	}
	splits = append(splits, cur.String())

	var ans strings.Builder
	for _, v := range splits {
		if len(v) >= 3 {
			ans.WriteString(v[:2])
		} else {
			ans.WriteString(v)
		}
	}

	return ans.String()
}
