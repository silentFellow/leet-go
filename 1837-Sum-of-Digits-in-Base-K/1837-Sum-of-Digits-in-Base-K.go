package leetcode

import (
	"strconv"
)

func sumBase(n int, k int) int {
  val, err := strconv.Atoi(strconv.FormatInt(int64(n), k))
  if err != nil {
    return -1
  }

  sum := 0
  for val != 0 {
    sum += int(val%10)
    val /= 10
  }

  return sum
}
