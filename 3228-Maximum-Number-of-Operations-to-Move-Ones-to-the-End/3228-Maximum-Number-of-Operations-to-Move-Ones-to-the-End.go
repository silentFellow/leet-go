package leetcode

func maxOperations(s string) int {
	ones, res := 0, 0

	for i, v := range s {
		if v == '1' {
			ones++
		} else if i > 0 && s[i-1] == '1' {
			res += ones
		}
	}

	return res
}
