package leetcode

import "strings"

func divideString(s string, k int, fill byte) []string {
	ans := make([]string, 0)

	var cur strings.Builder
	for i, v := range s {
		cur.WriteRune(v)
		if (i+1)%k == 0 {
			ans = append(ans, cur.String())
			cur.Reset()
		}
	}

	if cur.Len() != 0 {
		for range k - cur.Len() {
			cur.WriteByte(fill)
		}
		ans = append(ans, cur.String())
	}

	return ans
}
