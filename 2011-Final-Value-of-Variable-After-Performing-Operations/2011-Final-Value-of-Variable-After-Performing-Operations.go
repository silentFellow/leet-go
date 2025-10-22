package leetcode

func finalValueAfterOperations(operations []string) int {
	ans := 0

	for _, op := range operations {
		switch op[0] {
		case '+':
			ans += 1
		case '-':
			ans -= 1
		default:
			switch op[1] {
			case '+':
				ans += 1
			case '-':
				ans -= 1
			}
		}
	}

	return ans
}
