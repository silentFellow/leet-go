package leetcode

import (
	"errors"
	"strconv"
	"strings"
)

func isNumber(s string) bool {
	var resErr error
	if strings.Contains(s, ".") || strings.Contains(s, "e") || strings.Contains(s, "E") {
		if _, err := strconv.ParseFloat(s, 64); err != nil {
			resErr = err
		}
	} else {
		if _, err := strconv.Atoi(s); err != nil {
			resErr = err
		}
	}

	return resErr == nil || errors.Is(resErr, strconv.ErrRange)
}
