package leetcode

import "regexp"

func isMatch(s string, p string) bool {
    regex := regexp.MustCompile("^" + p + "$")
    return regex.MatchString(s)
}
