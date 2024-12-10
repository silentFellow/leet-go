package leetcode

import (
	"strings"
)

// complexity:
// Time: O(n)
// space: O(n) => can be optimized with O(1) => try to implement it

func maximumLength(s string) int {
  // calculate prefix sum
  prefixArray := make([]int, 0)
  prefix := s[0]
  prefixCount := 1

  for i:=1; i<len(s); i++ {
    if s[i] == prefix {
      prefixCount++
      continue
    }

    prefixArray = append(prefixArray, prefixCount)
    prefix = s[i]
    prefixCount = 1
  }
  prefixArray = append(prefixArray, prefixCount)

  // freqMap creation
  freqMap := make(map[string]int)
  valueAt := 0
  for _, times := range prefixArray {
    for i:=1; i<=times; i++ {
      valueInc := strings.Repeat(string(s[valueAt]), i)
      freqMap[valueInc] += (times-i+1)
    }

    valueAt += times
  }

  // finding answer
  ans := -1
  for key, val := range freqMap {
    if val >= 3 {
      ans = max(ans, len(key))
    }
  }

  return ans
}
