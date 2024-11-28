package leetcode

import (
	"strconv"
	"unicode"
)

func secondHighest(s string) int {
	var fm, sm int = -1, -1

	for _, val := range s {
		if unicode.IsDigit(val) {
			num, err := strconv.Atoi(string(val))

			if err == nil {
				if num > fm {
					sm, fm = fm, num
				} else if num > sm && num != fm {
					sm = num
				}
			}

		}
	}

	return sm
}
