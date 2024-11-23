package leetcode

import "strconv"

func reverse(s string) string {
  c := []rune(s)
  for i,j := 0, len(c)-1; i<j; i,j = i+1, j-1 {
    c[i], c[j] = c[j], c[i]
  }
  return string(c)
}

func isStrictlyPalindromic(n int) bool {
  for i:=2; i<(n-1); i++ {
    valAtBase := strconv.FormatInt(int64(n), i)

    if valAtBase != reverse(valAtBase) {
      return false
    }
  }

  return true
}
