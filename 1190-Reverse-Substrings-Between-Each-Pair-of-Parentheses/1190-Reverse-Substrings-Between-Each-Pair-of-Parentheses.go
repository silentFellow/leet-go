package leetcode

import (
	"strings"
)

func reverseParentheses(s string) string {
  stack := []rune{}

  for _, val := range s {
    if val == ')' {
      var temp strings.Builder

      for len(stack) > 0 && stack[len(stack)-1] != '('  {
        temp.WriteRune(stack[len(stack)-1])
        stack = stack[:len(stack)-1] // pop
      }
      stack = stack[:len(stack)-1] // pop remove the ')'

      stack = append(stack, []rune(temp.String())...)
      continue
    }

    stack = append(stack, val)
  }

  return string(stack)
}
