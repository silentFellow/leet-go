package leetcode

func minEatingSpeed(piles []int, h int) int {
	max := piles[0]
	for _, val := range piles {
		if val > max {
			max = val
		}
	}

	ans := max
	l, r := 1, max
	for l <= r {
		m := (l + r) / 2

		count := 0
		for _, val := range piles {
			count += (val / m)
			if val%m != 0 {
				count++
			}
		}

		if count <= h && ans >= m {
			ans = m
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return ans
}
