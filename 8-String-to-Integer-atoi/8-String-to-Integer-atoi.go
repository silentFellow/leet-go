package leetcode

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

func myAtoi(s string) int {
	regexFilter := regexp.MustCompile(`^\s*([+-]?\d+)`) // remove extra characters
	trimmed := strings.TrimSpace(regexFilter.FindString(s)) // remove whitespace

	atoi, err := strconv.Atoi(trimmed)

	// err occurs when empty or when atoi crosses int64 range
	if err != nil && len(trimmed) > 0 {
		if trimmed[0] == '-' {
			return math.MinInt32
		}

		return math.MaxInt32
	}

	// check if atoi is within int32 range
  if err == nil {
    if atoi > math.MaxInt32 {
      return math.MaxInt32
    } else if atoi < math.MinInt32 {
      return math.MinInt32
    }

    return atoi
  }

  return 0
}
