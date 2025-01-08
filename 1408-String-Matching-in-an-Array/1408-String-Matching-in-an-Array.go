package leetcode

import "strings"

func stringMatching(words []string) []string {
	prefix := make([]string, len(words))
	prefix[0] = ""
	for i := 1; i < len(words); i++ {
		prev := prefix[i-1]
		prefix[i] = prev + "," + words[i-1]
	}

	postfix := make([]string, len(words))
	postfix[len(words)-1] = ""
	for i := len(words) - 2; i >= 0; i-- {
		next := postfix[i+1]
		postfix[i] = next + "," + words[i+1]
	}

	ans := []string{}
	for i, val := range words {
		if strings.Contains(prefix[i], val) || strings.Contains(postfix[i], val) {
			ans = append(ans, val)
		}
	}

	return ans
}
