package leetcode

func getNoZeroIntegers(n int) []int {
	containsZero := func(n int) bool {
		for n != 0 {
			if n%10 == 0 {
				return true
			}
			n /= 10
		}
		return false
	}

	for i := 1; i < n; i++ {
		target := n - i
		if !containsZero(i) && !containsZero(target) {
			return []int{i, target}
		}
	}

	return []int{}
}
