package leetcode

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	for n%2 == 0 {
		n /= 2
	}

	for n%3 == 0 {
		n /= 3
	}

	for n%5 == 0 {
		n /= 5
	}

	return n == 1
}

// short hand:
// for _, factors := range []int{2, 3, 5} {
// 	for n%factors == 0 {
// 		n /= factors
// 	}
// }
