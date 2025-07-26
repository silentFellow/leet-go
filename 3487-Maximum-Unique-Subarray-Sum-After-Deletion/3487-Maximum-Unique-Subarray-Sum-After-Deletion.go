package leetcode

import "math"

func maxSum(nums []int) int {
	var maxNeg int = math.MinInt

	ans := 0
	hmap := make(map[int]struct{})
	for _, v := range nums {
		if v < 0 {
			maxNeg = max(maxNeg, v)
		} else {
			if _, ok := hmap[v]; !ok {
				ans += v
				hmap[v] = struct{}{}
			}
		}
	}

	_, ok := hmap[0]
	if ans == 0 && !ok {
		return maxNeg
	}

	return ans
}
