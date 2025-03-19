package leetcode

import (
	"math/bits"
	"slices"
)

func sortByBits(arr []int) []int {
  slices.SortFunc(arr, func(i, j int) int {
    icnt, jcnt := bits.OnesCount(uint(i)), bits.OnesCount(uint(j))
    if icnt == jcnt {
      return i - j
    }
    return icnt - jcnt
  })

  return arr
}
