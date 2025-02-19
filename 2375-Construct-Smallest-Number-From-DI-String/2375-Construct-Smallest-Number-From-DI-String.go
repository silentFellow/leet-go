package leetcode

import (
	"strconv"
	"strings"
)

func smallestNumber(pattern string) string {
  stack := []int{}
  var ans strings.Builder

  for i:=0; i<=len(pattern); i++ {
    stack = append(stack, i+1)

    if i == len(pattern) || pattern[i] == 'I' {
      for j:=len(stack)-1; j>=0; j-- {
        ans.WriteString(strconv.Itoa(stack[j]))
      }
      stack = []int{}
    }
  }

  return ans.String()
}
