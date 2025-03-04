package leetcode

import (
	"strconv"
	"strings"
)

func checkPowersOfThree(n int) bool {
  baseThree := strconv.FormatInt(int64(n), 3)

  return !strings.Contains(baseThree, "2")
}
