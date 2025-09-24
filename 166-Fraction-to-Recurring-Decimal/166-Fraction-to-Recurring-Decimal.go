package leetcode

import (
	"strconv"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	var ans strings.Builder

	// Handle sign
	if (numerator < 0) != (denominator < 0) {
		ans.WriteString("-")
	}

	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}

	num := abs(numerator)
	den := abs(denominator)

	// Integer part
	ans.WriteString(strconv.Itoa(num / den))
	remainder := num % den
	if remainder == 0 {
		return ans.String()
	}

	ans.WriteString(".")

	remainderMap := make(map[int]int)

	for remainder != 0 {
		if pos, ok := remainderMap[remainder]; ok {
			res := ans.String()
			return res[:pos] + "(" + res[pos:] + ")"
		}
		remainderMap[remainder] = ans.Len()
		remainder *= 10
		ans.WriteString(strconv.Itoa(remainder / den))
		remainder %= den
	}

	return ans.String()
}
