package leetcode

import "strconv"

func convertToBase7(num int) string {
	val := strconv.FormatInt(int64(num), 7)
	return val
}
