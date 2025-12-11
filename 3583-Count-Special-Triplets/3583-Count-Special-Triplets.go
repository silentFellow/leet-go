package leetcode

func specialTriplets(nums []int) int {
	const mod = 1_000_000_007
	n := len(nums)
	ans := 0

	// Suffix frequency map
	suffix := make(map[int]int)
	for _, num := range nums {
		suffix[num]++
	}

	prefix := make(map[int]int)
	for i := range n {
		suffix[nums[i]]--

		target := nums[i] * 2
		left := prefix[target]
		right := suffix[target]
		ans = (ans + left*right) % mod

		prefix[nums[i]]++
	}

	return ans
}
