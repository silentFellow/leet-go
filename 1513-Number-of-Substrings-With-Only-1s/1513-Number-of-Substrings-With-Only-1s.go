package leetcode

const MOD int = 1_000_000_007

func numSub(s string) int {
	ans, count := 0, 0

	for _, v := range s {
		if v == '1' {
			count++
		} else {
			ans += ((count * (count + 1)) / 2) % MOD
			count = 0
		}
	}

	if count > 0 {
		ans += ((count * (count + 1)) / 2) % MOD
		count = 0
	}

	return ans
}
