package leetcode

func minNumberOperations(target []int) int {
	ans := 0
	prev := 0

	for _, v := range target {
		if v > prev {
			ans += v - prev
		}
		prev = v
	}

	return ans
}
