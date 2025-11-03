package leetcode

func minCost(colors string, neededTime []int) int {
	ans := 0

	combined := 0
	maxv := 0
	for i := 1; i < len(colors); i++ {
		if colors[i] == colors[i-1] {
			if combined == 0 {
				combined += neededTime[i-1]
				maxv = neededTime[i-1]
			}
			combined += neededTime[i]
			maxv = max(maxv, neededTime[i])
		} else {
			ans += (combined - maxv)
			combined = 0
			maxv = 0
		}
	}

	if combined != 0 {
		ans += (combined - maxv)
		combined = 0
		maxv = 0
	}

	return ans
}
