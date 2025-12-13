package leetcode

import (
	"slices"
	"sort"
	"unicode"
)

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	n := len(code)

	isValidCode := func(s string) bool {
		for _, c := range s {
			if !(unicode.IsLetter(c) || unicode.IsDigit(c) || c == '_') {
				return false
			}
		}

		return true
	}

	partialAns := make([][2]string, 0, n)
	for i := range n {
		if isActive[i] && code[i] != "" {
			if slices.Contains(
				[]string{"electronics", "grocery", "pharmacy", "restaurant"},
				businessLine[i],
			) {
				if isValidCode(code[i]) {
					partialAns = append(partialAns, [2]string{code[i], businessLine[i]})
				}
			}
		}
	}

	sort.Slice(partialAns, func(i, j int) bool {
		bizOrder := map[string]int{
			"electronics": 0,
			"grocery":     1,
			"pharmacy":    2,
			"restaurant":  3,
		}
		pi, pj := bizOrder[partialAns[i][1]], bizOrder[partialAns[j][1]]
		if pi != pj {
			return pi < pj
		}
		return partialAns[i][0] < partialAns[j][0]
	})

	ans := make([]string, 0, len(partialAns))
	for _, v := range partialAns {
		ans = append(ans, v[0])
	}

	return ans
}
