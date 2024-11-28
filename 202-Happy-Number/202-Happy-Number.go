package leetcode

import (
	"math"
)

func sumOfSquares(n int) int {
	val := 0

	for n != 0 {
		val += int(math.Pow(float64(n%10), 2))
		n /= 10
	}

	return val
}

func isHappy(n int) bool {
	hmap := make(map[int]bool)

	for {
		if _, exists := hmap[n]; exists {
			break
		} else {
			hmap[n] = true
		}

		n = sumOfSquares(n)
		if n == 1 {
			return true
		}
	}

	return false
}
