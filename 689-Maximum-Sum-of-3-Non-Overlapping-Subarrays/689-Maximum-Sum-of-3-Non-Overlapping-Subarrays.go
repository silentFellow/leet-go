package leetcode

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums)
	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}

	dp := make([][]int, 4)
	for i := 0; i < 4; i++ {
		dp[i] = make([]int, n-k+2)
	}

	maxSum := 0
	for l := 1; l <= 3; l++ {
		for i := 1; i <= n-k+1; i++ {
			sum := prefixSum[i+k-1] - prefixSum[i-1]
			dp[l][i] = max(dp[l-1][i], sum)
            if i > 0 {
                dp[l][i] = max(dp[l][i], dp[l][i-1])
            }
            if i >= k {
                dp[l][i] = max(dp[l][i], sum+dp[l-1][i-k])
            }
			maxSum = max(maxSum, dp[l][i])
		}
	}
	res := make([]int, 3)
	for i := 3; i >= 1; i-- {
		firstMaxIdx := -1
		for j := 1; j <= n-k+1; j++ {
			if dp[i][j] == maxSum {
				firstMaxIdx = j
				break
			}
		}

		res[i-1] = firstMaxIdx - 1
        sum := prefixSum[firstMaxIdx+k-1] - prefixSum[firstMaxIdx-1]
		maxSum -= sum
	}

	return res
}
