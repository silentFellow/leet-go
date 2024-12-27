package leetcode

func maxScoreSightseeingPair(values []int) int {
	prevMax := values[0]

	ans := 0
	for i := 1; i < len(values); i++ {
		ans = max(ans, prevMax+values[i]-i)
		prevMax = max(prevMax-i+1, values[i]) // max(prevMax, values[i]+i)
	}

	return ans
}
