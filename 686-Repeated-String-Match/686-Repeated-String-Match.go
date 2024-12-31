package leetcode

import (
	"math"
	"strings"
)

func repeatedStringMatch(a string, b string) int {
	concatString := ""
	times := 0
	maxTime := int(math.Ceil(float64(len(b))/float64(len(a)))) + 2

	for {
		if times > maxTime {
			return -1
		}

		concatString += a
		times++

		if strings.Index(concatString, b) != -1 {
			return times
		}
	}
}
