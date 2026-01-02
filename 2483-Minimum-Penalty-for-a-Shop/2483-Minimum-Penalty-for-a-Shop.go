package leetcode

func bestClosingTime(customers string) int {
	fullPenalty := 0
	for _, c := range customers {
		if c == 'Y' {
			fullPenalty++
		}
	}

	curPenalty := 0
	minPenalty := fullPenalty
	ans := 0
	for i, c := range customers {
		if c == 'Y' {
			curPenalty++
		} else {
			curPenalty--
		}

		actualPenalty := fullPenalty - curPenalty
		if actualPenalty < minPenalty {
			minPenalty = actualPenalty
			ans = i + 1
		}
	}

	return ans
}
