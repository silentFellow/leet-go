package leetcode

func minDeletionSize(strs []string) int {
	n := len(strs)
	m := len(strs[0])

	sorted := make([]bool, n-1)
	ans := 0

	for c := range m {
		bad := false

		for i := 0; i < n-1; i++ {
			if !sorted[i] && strs[i][c] > strs[i+1][c] {
				bad = true
				break
			}
		}

		if bad {
			ans++
			continue
		}

		for i := 0; i < n-1; i++ {
			if strs[i][c] < strs[i+1][c] {
				sorted[i] = true
			}
		}
	}

	return ans
}
