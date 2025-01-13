package leetcode

func canBeValid(s string, locked string) bool {
	if len(s)%2 != 0 {
		return false
	}

	open, close := 0, 0
	for i := range s {
		if s[i] == '(' || locked[i] == '0' {
			open++
		} else {
			close++
		}

		if close > open {
			return false
		}
	}

	open, close = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' || locked[i] == '0' {
			close++
		} else {
			open++
		}

		if open > close {
			return false
		}
	}

	return true
}
