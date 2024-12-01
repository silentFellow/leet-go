package leetcode

import (
	"strconv"
	"strings"
)

func hammingWeight(n int) int {
	nstr := strconv.FormatInt(int64(n), 2)
	return strings.Count(nstr, "1")
}
