package leetcode

import "strings"

func addSpaces(s string, spaces []int) string {
  var ans strings.Builder

  i := 0
  for _, indexVal := range spaces {
    ans.WriteString(s[i:indexVal])
    ans.WriteString(" ")
    i = indexVal
  }
  ans.WriteString(s[i:])

  return ans.String()
}
