package leetcode

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	arr := strings.Split(s, " ")
	return len(arr[len(arr)-1])
}
