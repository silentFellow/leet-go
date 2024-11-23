package leetcode

import (
	"math"
)

func scoreOfString(s string) int {
  val := 0

  for i := range (len(s)-1) {
    abs := math.Abs(float64(s[i]) - float64(s[i+1]))
    val += int(abs)
  }

  return val
}
