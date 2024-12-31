package leetcode

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	mtrx := make([]string, numRows)

	curRow := 0
	goingDown := false // coz initial change will happen

	for _, char := range s {
		mtrx[curRow] += string(char)

		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}

		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	var result strings.Builder
	for _, row := range mtrx {
		result.WriteString(row)
	}

	return result.String()
}
