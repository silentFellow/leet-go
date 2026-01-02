package leetcode

import "slices"

func minimumBoxes(apple []int, capacity []int) int {
	totalApples := 0
	for _, a := range apple {
		totalApples += a
	}

	slices.Sort(capacity)
	slices.Reverse(capacity)

	ans := 0
	for _, c := range capacity {
		if totalApples <= 0 {
			break
		}

		totalApples -= c
		ans++
	}

	return ans
}
