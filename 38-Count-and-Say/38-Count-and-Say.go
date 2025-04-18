package leetcode

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	var ans strings.Builder
	ans.WriteString("1")

	for i := 1; i < n; i++ {
		var temp strings.Builder
		s := ans.String()

		count := 1
		for j := 1; j < len(s); j++ {
			if s[j] == s[j-1] {
				count++
			} else {
				temp.WriteString(numToStr(count))
				temp.WriteByte(s[j-1])
				count = 1
			}
		}

		temp.WriteString(numToStr(count))
		temp.WriteByte(s[len(s)-1])
		ans.Reset()
		ans.WriteString(temp.String())
	}

	return ans.String()
}

func numToStr(v int) string {
	return strconv.Itoa(v)
}
