package leetcode

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	s1, s2 := strings.Split(version1, "."), strings.Split(version2, ".")
	n1, n2 := len(s1), len(s2)
	n := min(n1, n2)

	for i := range n {
		v1, err := strconv.Atoi(s1[i])
		if err != nil {
			return -1
		}

		v2, err := strconv.Atoi(s2[i])
		if err != nil {
			return -1
		}

		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
	}

	if n1 > n {
		for i := n; i < n1; i++ {
			val, err := strconv.Atoi(s1[i])
			if err != nil || val != 0 {
				return 1
			}
		}
	} else if n2 > n {
		for i := n; i < n2; i++ {
			val, err := strconv.Atoi(s2[i])
			if err != nil || val != 0 {
				return -1
			}
		}
	}

	return 0
}
