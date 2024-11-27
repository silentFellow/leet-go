package leetcode

import "strings"

func toLowerCase(s string) string {
	// for interview subtract each character by 32
  // rune := []rune(s)
  // for i := range rune {
  //   if rune[i] >= 'A' && rune[i] <= 'Z' {
  //     rune[i] += 32
  //   }
  // }
  // return string(rune)

	return strings.ToLower(s)
}
