package leetcode

import "math"

func areaOfMaxDiagonal(dimensions [][]int) int {
	ans := 0
	var maxDiagonal float64

	for _, dimension := range dimensions {
		f, s := dimension[0], dimension[1]

		curDiagonal := math.Sqrt(float64(f*f + s*s))
		if maxDiagonal == curDiagonal {
			ans = max(ans, f*s)
		} else if maxDiagonal < curDiagonal {
			maxDiagonal = curDiagonal
			ans = f * s
		}
	}

	return ans
}
