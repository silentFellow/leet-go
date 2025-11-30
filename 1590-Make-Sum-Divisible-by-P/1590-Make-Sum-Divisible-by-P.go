package leetcode

func minSubarray(nums []int, p int) int {
	extraSum := 0
	for _, num := range nums {
		extraSum = (extraSum + num) % p
	}

	if extraSum == 0 {
		return 0
	}

	hmap := make(map[int]int)
	hmap[0] = -1

	ans := len(nums)
	prefixSum := 0
	for i, num := range nums {
		prefixSum = (prefixSum + num) % p
		req := (prefixSum - extraSum + p) % p

		if start, ok := hmap[req]; ok {
			ans = min(ans, i-start)
		}

		hmap[prefixSum] = i
	}

	if ans == len(nums) {
		return -1
	}
	return ans
}
