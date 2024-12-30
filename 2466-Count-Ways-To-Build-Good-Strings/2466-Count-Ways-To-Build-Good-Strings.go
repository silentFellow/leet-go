package leetcode

import "math"

func countGoodStrings(low int, high int, zero int, one int) int {
	mod := int(math.Pow(10, 9)) + 7
	dp := make(map[int]int)

	var backtrack func(length int) int
	backtrack = func(length int) int {
		if length > high {
			return 0
		}

		if val, ok := dp[length]; ok {
			return val
		}

		res := 0
		if length >= low && length <= high {
			res = 1
		}

		zeroType := length + zero
		oneType := length + one

		res += backtrack(zeroType)
		res += backtrack(oneType)
    dp[length] = res

		return res % mod
	}

	return backtrack(0)
}
