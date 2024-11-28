package leetcode

func xorOperation(n int, start int) int {
	ans := 0

	for i := range n {
		ans ^= (start + 2*i)
	}

	return ans
}
