package leetcode

func shortestCommonSupersequence(str1 string, str2 string) string {
	m, n := len(str1), len(str2)

	prev := make([]string, n+1)
	for i := range n {
		prev[i] = str2[i:]
	}
	prev[n] = ""

	for i := m - 1; i >= 0; i-- {
		cur := make([]string, n+1)
		cur[n] = str1[i:]

		for j := n - 1; j >= 0; j-- {
			if str1[i] == str2[j] {
				cur[j] = string(str1[i]) + prev[j+1]
			} else {
				r1 := string(str1[i]) + prev[j]
				r2 := string(str2[j]) + cur[j+1]

				if len(r1) < len(r2) {
					cur[j] = r1
				} else {
					cur[j] = r2
				}
			}
		}

		prev = cur
	}

	return prev[0]
}
