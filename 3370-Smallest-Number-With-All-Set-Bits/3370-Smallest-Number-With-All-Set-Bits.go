package leetcode

import (
	"math"
	"strconv"
	"strings"
)

func smallestNumber(n int) int {
	// find next 2**n, then the value will be it's -1
	// alternatively you can use constraint and use 1000
	x := 0
	for (1 << x) <= n {
		x++
	}

	for i := n; i < int(math.Pow(2, float64(x))); i++ {
		binary := strconv.FormatInt(int64(i), 2)

		if !strings.Contains(binary, "0") {
			return i
		}
	}

	return -1
}
