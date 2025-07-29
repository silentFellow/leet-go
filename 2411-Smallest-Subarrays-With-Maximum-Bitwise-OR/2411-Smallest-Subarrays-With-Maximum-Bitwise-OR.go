package leetcode

func smallestSubarrays(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	lastSeen := make([]int, 32)

	for i := n - 1; i >= 0; i-- {
		v := nums[i]

		for b := range 32 {
			if (v>>b)&1 == 1 {
				lastSeen[b] = i
			}
		}

		maxIndex := i
		for b := range 32 {
			if lastSeen[b] > maxIndex {
				maxIndex = lastSeen[b]
			}
		}

		ans[i] = maxIndex - i + 1
	}

	return ans
}
