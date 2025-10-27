package leetcode

import "strings"

func numberOfBeams(bank []string) int {
	ans := 0
	prev := 0

	for _, v := range bank {
		curr := strings.Count(v, "1")
		if curr == 0 {
			continue
		}

		ans += (prev * curr)
		prev = curr
	}

	return ans
}
