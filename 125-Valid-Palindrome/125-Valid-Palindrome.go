package leetcode

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
    var sb strings.Builder

    for _, val := range s {
        if unicode.IsLetter(val) {
            sb.WriteRune(unicode.ToLower(val))
        } else if unicode.IsDigit(val) {
            sb.WriteRune(val)
        }
    }

    s = sb.String()
    for i,j := 0, len(s)-1; i<j; i,j = i+1,j-1 {
        if s[i] != s[j] {
            return false
        }
    }

    return true
}
