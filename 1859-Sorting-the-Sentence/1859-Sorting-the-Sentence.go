package leetcode

import (
	"sort"
	"strings"
)

func sortSentence(s string) string {
	sArr := strings.Split(s, " ")
	sort.Slice(sArr, func(i, j int) bool {
		m, n := len(sArr[i]), len(sArr[j])
		return sArr[i][m-1] < sArr[j][n-1]
	})

	for i, val := range sArr {
		sArr[i] = val[:len(val)-1]
	}

	return strings.Join(sArr, " ")
}
