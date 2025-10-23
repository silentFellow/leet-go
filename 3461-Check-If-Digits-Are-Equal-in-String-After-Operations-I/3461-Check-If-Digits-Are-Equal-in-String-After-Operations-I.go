package leetcode

import "strconv"

func hasSameDigits(s string) bool {
	split := make([]int, len(s))
	for i, char := range s {
		if val, err := strconv.Atoi(string(char)); err == nil {
			split[i] = val
		} else {
			return false
		}
	}

	for len(split) != 2 {
		neoSplit := make([]int, 0)
		for i := range len(split) - 1 {
			neoSplit = append(neoSplit, (split[i]+split[i+1]) % 10)
		}
		split = neoSplit
	}

	return split[0] == split[1]
}
