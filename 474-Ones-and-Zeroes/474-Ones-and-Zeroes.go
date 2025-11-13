package leetcode

import "strings"

func findMaxForm(strs []string, m int, n int) int {
	// dp[i][j]: max subset size with at most i zeros and j ones
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		zeros := strings.Count(str, "0")
		ones := strings.Count(str, "1")
		// iterate backwards to avoid overwriting needed values
		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				if dp[i-zeros][j-ones]+1 > dp[i][j] {
					dp[i][j] = dp[i-zeros][j-ones] + 1
				}
			}
		}
	}
	return dp[m][n]
}
