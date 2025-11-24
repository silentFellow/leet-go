package leetcode

import (
	"math"
	"sort"
)

func maxSumDivThree(nums []int) int {
	sum := 0
	r1, r2 := make([]int, 0), make([]int, 0)

	for _, v := range nums {
		sum += v
		switch v % 3 {
		case 1:
			r1 = append(r1, v)
		case 2:
			r2 = append(r2, v)
		}
	}

	sort.Ints(r1)
	sort.Ints(r2)

	if sum%3 == 1 {
		cand1 := math.MaxInt
		if len(r1) > 0 {
			cand1 = r1[0]
		}
		cand2 := math.MaxInt
		if len(r2) > 1 {
			cand2 = r2[0] + r2[1]
		}

		// if sum % 3 == 1, then need to remove 1 extra candy
		// either remove one (one remainded candy)
		// or remove two (two remainded candy)
		sum -= min(cand1, cand2)
	}

	if sum%3 == 2 {
		cand1 := math.MaxInt
		if len(r2) > 0 {
			cand1 = r2[0]
		}
		cand2 := math.MaxInt
		if len(r1) > 1 {
			cand2 = r1[0] + r1[1]
		}

		// if sum % 3 == 2, then need to remove 1 extra candy
		// either remove two (two remainded candy)
		// or remove one (one remainded candy)
		sum -= min(cand1, cand2)
	}

	return sum
}
