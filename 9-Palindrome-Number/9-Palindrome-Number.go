package leetcode

import "strconv"

func reverse(s string) string {
  val := []rune(s)
  for i,j := 0, len(val)-1; i<j; i, j = i+1, j-1 {
    val[i], val[j] = val[j], val[i]
  }
  return string(val)
}

func isPalindrome(x int) bool {
  xStr := strconv.Itoa(x)
  return xStr == reverse(xStr)
}
