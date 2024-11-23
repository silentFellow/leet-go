package leetcode

import "math"

func judgeSquareSum(c int) bool {
	l := 0
	r := int(math.Sqrt(float64(c)))

	for l <= r {
		cur := (l * l) + (r * r)

		if cur == c {
			return true
		} else if cur < c {
			l += 1
		} else {
			r -= 1
		}
	}

	return false
}
