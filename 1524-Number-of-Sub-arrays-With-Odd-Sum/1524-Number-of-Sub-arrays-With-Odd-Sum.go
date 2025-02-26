package leetcode

func numOfSubarrays(arr []int) int {
	const MOD = 1_000_000_007
	eve, odd := 1, 0
	sum, ans := 0, 0

	for _, v := range arr {
		sum += v

		if sum%2 == 0 {
			ans = (ans + odd) % MOD
			eve++
		} else {
			ans = (ans + eve) % MOD
			odd++
		}
	}

	return ans
}
