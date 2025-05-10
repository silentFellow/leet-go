package leetcode

import "slices"

func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)

	l, r := 0, len(people)-1
	ans := 0

	for l <= r {
		if people[l]+people[r] <= limit {
			l++
		}
		r--
		ans++
	}

	return ans
}
